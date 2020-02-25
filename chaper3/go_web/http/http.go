package main
import(
	"fmt"
	"net/http"
)
type myHandler1 struct{
	s string

}
func Handler1(w http.ResponseWriter,r *http.Request){
	fmt.Println("begin to handle request!")
	fmt.Fprintf(w,"hello,%s",r.URL.String()[1:])
}
func (h *myHandler1)ServeHTTP(w http.ResponseWriter, re *http.Request){
	fmt.Fprintf(w,"hello,%s",re.URL.String()[1:])
}
func main(){
	//http.HandleFunc("/",Handler1)
	//http.Handle("/",http.HandlerFunc(Handler1))
	http.Handle("/hello/",&myHandler1{"hi!man!"})

	go http.ListenAndServe("localhost:8080",nil)
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//s:=testStruct{}
	//var v testInterface=&s
	//
	//run(testfun)
}
