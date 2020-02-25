package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)
type Counter struct{
	n int
}
type Chan chan int
var webroot=flag.String("root","/home/user","web root directory")
var helloRequests=expvar.NewInt("hello-requests")
func main(){
	http.HandleFunc("/",logger)
	http.Handle("/go/hello",http.HandlerFunc(HelloServer))
	ctr:=new(Counter)
	expvar.Publish("/counter",ctr)
	http.Handle("/counter",ctr)
	http.Handle("/go/",http.StripPrefix("/go/",http.FileServer(http.Dir(*webroot))))
	http.Handle("/flags",http.HandlerFunc(FlagServer))
	http.Handle("/args",http.HandlerFunc(ArgServer))
	http.Handle("/chan",ChanCreate())
	http.Handle("/data",http.HandlerFunc(DataServer))
	if err:=http.ListenAndServe(":12345",nil);err!=nil{
		fmt.Println(err.Error())
	}
}
func HelloServer(w http.ResponseWriter,r *http.Request){
	helloRequests.Add(1)
	io.WriteString(w,"helloworld!\n")
}
func logger(w http.ResponseWriter,r *http.Request){
	log.Print(r.URL.String())
	w.WriteHeader(404)
	w.Write([]byte("oops"))
}
func (ctr *Counter)String()string{
	return fmt.Sprintf("%d",ctr.n)
}
func (ctr *Counter)ServeHTTP(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case "GET":ctr.n++
	case "POST":
		buf:=new(bytes.Buffer)
		io.Copy(buf,r.Body)
		body:=buf.String()
		if n,err:=strconv.Atoi(body);err!=nil{
			fmt.Fprintf(w,"bad POST:%v\nbody:[%v]",err,body)
		}else{
			ctr.n=n
			fmt.Fprint(w,"Counter reset\n")
		}
	}
	fmt.Fprintf(w,"counter = %d\n",ctr.n)
}
func FlagServer(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/plain;charset=utf-8")
	fmt.Fprint(w,"Flags:\n")
	flag.VisitAll(func(f *flag.Flag){
		if f.Value.String()!= f.DefValue{
			fmt.Fprintf(w,"%s = %s[default=%s]\n",f.Value.String(),f.DefValue)
		}else{
			fmt.Fprintf(w,"%s = %s\n",f.Name,f.Value.String())
		}
	})
}
func ArgServer(w http.ResponseWriter,r *http.Request){
	for _,s:=range(os.Args){
		fmt.Fprint(w,s," ")
	}
}
func ChanCreate()Chan{
	c:=make(Chan)
	go func (c Chan){
		for x:=0;;x++{
			c<-x
		}
	}(c)
	return c
}
func (ch Chan)ServeHTTP(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,fmt.Sprintf("chanel send #%d\n",<-ch))
}
func DataServer(rw http.ResponseWriter,req *http.Request){
	rw.Header().Set("Content-Type","text/plain;charset=utf-8")
	r,w,err:=os.Pipe()
	if err!=nil{
		fmt.Fprintf(rw,"Pipe:%s\n",err)
		return
	}
	p,err:=os.StartProcess("/bin/date",[]string{"data"},&os.ProcAttr{Files:[]*os.File{nil,w,w}})
	defer r.Close()
	w.Close()
	if err!=nil{
		fmt.Fprintf(rw,"fork/exec:%s\n",err)
		return
	}
	defer p.Release()
	io.Copy(rw,r)
	wait,err:=p.Wait()
	if err!=nil{
		fmt.Fprintf(rw,"wait:%s\n",err)
		return
	}
	if !wait.Exited(){
		fmt.Fprintf(rw,"date:%v\n",wait)
		return
	}

}