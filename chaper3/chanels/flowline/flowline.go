package main

import (
	"fmt"
	"sync"
)

func main(){
	//ch1:=Producer()
	//ch2:=Square(ch1)
	//for i:=range(ch2){
	//	fmt.Println(i)
	//}
	ch:=Producer()
	ch1:=Square(ch)
	ch2:=Square(ch)
	ch3:=Square(ch)
	out:=Merger(ch1,ch2,ch3)
	for i:=range(out){
		fmt.Println(i)
	}
}
func Producer()<-chan int{
	ch:=make(chan int)
	go func()(){
		defer close(ch)
		for i:=0;i<10;i++{
			ch<-i
		}
	}()

	return ch
}

func Square(in <-chan int) <-chan int{
	out:=make(chan int)
	go func()(){
		defer close(out)
		for i:=range(in){
			out<-i*i
		}
	}()

	return out
}
func Merger(chs ...<-chan int)<-chan int{
	out:=make(chan int)

	wg:=sync.WaitGroup{}
	wg.Add(len(chs))
	for _,ch:=range(chs){
		go func()(){
			for i:=range(ch){
				out<-i
			}
			wg.Done()
		}()
	}
	go func()(){
		wg.Wait()
		close(out)
	}()
	return out
}