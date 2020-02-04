package main

import (
	"fmt"
	"sort"
	"time"
)
type vector []int
func (v vector)Less(i,j int )bool{
	return v[i]<v[j]
}
func (v vector)Swap(i,j int){
	v[i],v[j]=v[j],v[i]
}
func (v vector)Len()int{
	return len(v)
}
func (v vector)Iter()<- chan int{
	ch:=make(chan int)
	go func(){
		for _,v:=range(v){
			ch<-v
		}
	}()
	return ch
}
func main(){
	//ch:=make(chan string,1)
	//go sender(ch)
	//go receiver(ch)
	//time.Sleep(3*1e9)

	//intChan:=make(chan int)
	//go pump(intChan)
	//go suck(intChan)
	//time.Sleep(1e8)
	//ch:=make(chan int)
	//go func()(){
	//	time.Sleep(2*1e9)
	//	fmt.Println(<-ch)
	//}()
	//fmt.Println("sending")
	//ch<-10
	//fmt.Println("sent")
	//ch:=make(chan int)
	//v:=vector{10,3,2,5,4}
	//go sortVector(v,ch)
	//<-ch
	//fmt.Println(v[0])

	v:=vector{1,2,3,4}
	ch:=v.Iter()
	go suck(ch)
	time.Sleep(1e9)

}

func sender(ch chan string){
	fmt.Println("begin send")
	ch<-"hello"
	fmt.Println("send hello")
	ch<-"man"
	fmt.Println("send man")
}
func receiver(ch chan string){
	time.Sleep(2*1e9)
	for {
		str := <-ch
		fmt.Println(str)
	}
}
func suck(ch <-chan int)(){
	for {
		fmt.Println(<-ch)
	}
}
func pump(ch chan int)(){
	for i:=0;;i++{
		ch<-i
	}
}

func sortVector(v vector,ch chan int){
	sort.Sort(v)
	ch<-1
}