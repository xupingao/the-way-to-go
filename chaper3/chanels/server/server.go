package main

import (
	"fmt"
	"time"
)

type reply struct{
	res int
}
type request struct{
	a,b int
	replych chan *reply
}
func run(fun func(a,b int)int,re *request){
	sem<-1
	re.replych<-&reply{fun(re.a,re.b)}
	<-sem
}
var sem=make(chan int,10)
func server(fun func(a,b int)int,request <-chan *request,quitCh <-chan bool){

	la:
	for{
		select{
		case re:=<-request:		go run(fun,re)
		case <-quitCh:
			{
				fmt.Println("Server stoped!")
				break la
			}
		}
		}

}
func startServer(fun func(a,b int)int)(chan *request,chan bool){
	quitCh:=make(chan bool)
	requestCh:=make(chan *request)
	go server(fun,requestCh,quitCh)
	return requestCh,quitCh
}
func main(){

	requestCh,quitCh:=startServer(func(a,b int)int{return a+b})
	receiver:=make([]func()int,10)
	for i:=0;i<10;i++{
		receiver[i]=makeRequest(i,i*3,requestCh)
	}
	for i:=0;i<10;i++{
		fmt.Println(receiver[i]())
	}
	quitCh<-true
	time.Sleep(3*1e9)
}
func makeRequest(a,b int,requestCh chan *request)func()int{
	replych:=make(chan *reply)
	re:=request{a,b,replych}
	requestCh<-&re
	return func()int{
		return (<-replych).res
	}
}