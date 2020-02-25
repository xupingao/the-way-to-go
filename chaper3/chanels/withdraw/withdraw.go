package main

import (
	"fmt"
	"runtime"
)

func  main(){
	//withdrawOK2()
	ch1:=FanOut()
	ch2:=FanOut()
	ch:=FanIn(ch1,ch2)
	for i:=range(ch){
		fmt.Println(i)
	}
	runtime.GOMAXPROCS()
}

func withdrawOK(){
	ch1,ch2:=make(chan int),make(chan int)
	go func()(){
		defer close(ch1)
		for i:=0;;i++{
			ch1<-i
		}
	}()
	go 	func()(){
		defer close(ch2)
		for i:=0;i<3;i++{
			ch2<--1
		}
	}()
	for ;;{
		select{
		case i,ok:=<-ch2:{
			if !ok{
			fmt.Println("withdraw by ok")
			return
			}
			fmt.Println(i)
		}
		case i:=<-ch1:{
			fmt.Println(i)
		}
		}
	}
}

func withdrawOK2(){
	ch1,ch2:=make(chan int),make(chan int)
	go func()(){
		defer close(ch1)
		for i:=0;;i++{
			ch1<-i
		}
	}()
	go 	func()(){
		defer close(ch2)
		for i:=0;i<3;i++{
			ch2<--1
		}
	}()
	for ;;{
		select{
		case i,ok:=<-ch2:{
			if !ok{
				fmt.Println("ch2end ")
				ch2=nil
			}
			fmt.Println(i)
		}
		case i:=<-ch1:{
			fmt.Println(i)
		}
		}
	}
}

func FanOut()<-chan int{
	ch:=make(chan int)
	go func()(){
		defer close(ch)
		for i:=0;i<3;i++{
			ch<-i
		}
	}()
	return ch
}

func FanIn(ch1,ch2 <-chan int)(<-chan int){
	ch:=make(chan int)
	go func()(){
		defer close(ch)
		for {
			select {
			case i, ok := <-ch1:
				{
					if !ok {
						ch1 = nil
					}
					ch <- i
				}
			case i, ok := <-ch2:
				{
					if !ok {
						ch2 = nil
					}
					ch <- i
				}
			}
			if ch1==nil&&ch2==nil{
				return
			}
		}
	}()
	return ch
}