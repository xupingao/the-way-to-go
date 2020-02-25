package main

import (
	"fmt"
	"time"
)

func main(){
	//out:=make(chan int)
	//close(out)
	//if val,ok:=<-out;ok{
	//	fmt.Println(val)
	//}else{
	//	fmt.Println("closed")
	//}
	ch:=make(chan string)
	ch1:=make(chan bool)
	go func()(){
		ch<-"hello"
		ch<-"man"
		ch1<-true
	}()
	time.Sleep(5*1e9)
	for{
		//if val,ok:=<-ch;ok{
		//	fmt.Println(val)
		//}else{
		//	break
		//}
		select{
		case res:=<-ch:fmt.Println(res)
		case <-ch1:goto out
		}
	}
	out:
	fmt.Println("end main()")
}
