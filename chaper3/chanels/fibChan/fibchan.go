package main

import (
	"fmt"
	"time"
)

func fib(n int,ch chan int)(){
	x,y:=1,1
	for i:=0;i<n/2;i++{
		ch<-x
		x,y=y,x+y
	}
	fmt.Println("过半")
	time.Sleep(5*1e9)
	for i:=n/2;i<n;i++{
		ch<-x
		x,y=y,x+y
	}
	close(ch)
}
func main(){
	c:=make(chan int,10)
	go fib(10,c)
	for i:=range(c){
		fmt.Println(i)
	}
}
