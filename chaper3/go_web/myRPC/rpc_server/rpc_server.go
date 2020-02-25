package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main(){
	listener,err:=net.Listen("tcp","localhost:1234")
	if err!=nil{
		  log.Fatal(err)
	}
	http.Handle("/myrpc",http.HandlerFunc(handle))

	go http.Serve(listener,nil)
	time.Sleep(1000e9)
}
func handle(w http.ResponseWriter,r*http.Request){
	io.WriteString(w,"helloworld!")
}