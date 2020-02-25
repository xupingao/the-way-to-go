package main

import "fmt"

type Empty interface{}
var empty Empty
type semaphore chan Empty
func (s semaphore)P(n int){
	for i:=0;i<n;i++{
		s<-empty
	}
}
func (s semaphore)V(n int)(){
	for i:=0;i<n;i++{
		<-s
	}
}
func (s semaphore)Lock(){
	s.P(1)
}
func (s semaphore)Unlock(){
	s.V(1)
}
func (s semaphore)Wait(n int){
	s.P(n)
}
func (s semaphore)Singal(){
	s.V(1)
}
func main(){
	finshCh:=make(semaphore)
	ch:=send(finshCh)
	go receive(ch,finshCh)
	finshCh.V(1)
	finshCh.V(1)
	fmt.Println("Over")

}
func send(finCh semaphore)(ch chan int){
	ch=make(chan int)
	go func(ch chan int){
		for i:=0;i<10;i++{ch<-i}
		//close(ch)
		finCh.P(1)
	}(ch)
	return
}
func receive(ch chan int,s semaphore){
	for v:=range(ch){
		fmt.Println(v)
	}
	s.P(1)
}