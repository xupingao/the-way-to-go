package main

import (
	"fmt"
	"time"
)

func main() {
//	tickCh:=time.NewTicker(1e9)
//	for i:=0;i<5;i++{
//		<-tickCh.C
//		fmt.Println("tick")
//	}
//	tick:=time.Tick(1e9)
//	for i:=0;i<5;i++{
//		<-tick
//		fmt.Println("tick")
//	}
//
//	timer:=time.NewTimer(3*1e9)
//	<-timer.C
//	fmt.Println("It is time!")
//
//	ring:=time.After(5*1E9)
//	<-ring
//	fmt.Println("Boom!" )
go receive(printer())
time.Sleep(10*1e9)
}
func printer()(<-chan int){
	ch:=make(chan int)
	go func()(){
		for i:=0;;i++ {
			ch <- i
		}
	}()
	return ch

}

func receive(ch <-chan int)(){
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-time.After(1e5):
			{
				fmt.Println("It is time!")
				return
			}
		}
	}
}