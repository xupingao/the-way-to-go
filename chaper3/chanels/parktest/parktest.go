package main

import (
	"fmt"
	"time"
)

func main(){
	ch:=make(chan int,3)
	go func()(){
		for i:=1;i<=2;i++{
			ch<-i
		}
	}()
	go func()(){
		for i:=1;i<=2;i++{
			ch<-i
		}
	}()
	go func()(){
		time.Sleep(1E9)
		for i:=1;i<=4;i++{
			fmt.Println(<-ch)
		}
	}()

	time.Sleep(3*1e9)
}
