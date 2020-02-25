package main

import (
	"fmt"
	"time"
)

func send(s string)(chan string){
	str:=fmt.Sprintf("this is a message from %s",s)
	ch:=make(chan string)
	go func()(){
		for i:=0;i<10;i++{
			ch<-str
		}
	}()
	return ch
}
func receive(ch1,ch2 chan string)(){
	for {
		select {
		case str := <-ch1:
			fmt.Println(str)
		case str := <-ch2:
			fmt.Println(str)
		}
	}
}
func main(){
	ch1:=send("kobe")
	ch2:=send("wade")
	go receive(ch1,ch2)
	time.Sleep(10*1e9)
}