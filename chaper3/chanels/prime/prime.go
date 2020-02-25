package main

import "fmt"

func generate()(<-chan int){
	ch:=make(chan int)
	go func()(){
		for i:=2;;i++{
			ch<-i
		}
	}()
	return ch
}
func filter(in <-chan int,prime int)(out chan int){
	out=make(chan int)
	go func()() {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return
}
func sieve()(out<-chan int){
	out=make(chan int)
	ch:=generate()
	for{
		prime:=<-ch
		fmt.Print(prime," ")
		ch1:=filter(ch,prime)
		ch=ch1
	}
	return
}
func main(){
	ch:=sieve()
	for{
		fmt.Println(<-ch)
	}
}
