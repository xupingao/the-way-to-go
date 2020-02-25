package main
import(
	"fmt"
	"io"
	"net/http"
	"time"
)
const form=`
<html>
 <body>
  <form action="#" method="post" name="bar">
    <input type="text" name="in"/>
    <input type="submit" value="submit"/>
  </form>
 </body>
</html>`
func simpleServer(w http.ResponseWriter,request *http.Request){
	io.WriteString(w,"<h1>hello,world</h1>")
}
func FormServer(w http.ResponseWriter,request *http.Request){
	w.Header().Set("Content-Type","text/html")
	switch request.Method{
	case "GET":io.WriteString(w,form)
	case "POST":
		{
			ch:=make(chan int)
			go process(ch)
			<-ch
			fmt.Fprintln(w, request.FormValue("in"))
		}
	}

}
func main(){
	http.Handle("/test1",http.HandlerFunc(logHandler(simpleServer)))
	http.HandleFunc("/test2",FormServer)
	err:=http.ListenAndServe(":8080",nil)
	handleError(err)

}
func handleError(err error){
	if err!=nil{
		fmt.Println(err.Error())
	}
}
func logHandler(handerfunc func(w http.ResponseWriter,r *http.Request)())func(w http.ResponseWriter,r *http.Request)(){
	return func(w http.ResponseWriter,r *http.Request)(){
		defer func()(){if err:=recover();err!=nil{fmt.Println(err)}}()
		handerfunc(w,r)
	}
}
func process(ch chan int)(){
	time.Sleep(5*1e9)
	ch<-1
}