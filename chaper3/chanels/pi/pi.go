package main

import (
	"fmt"
	"math"
	"time"
)

func util(ch chan float64)(){
	t:=1
	for i:=0;i<1000;i++{
		go term(ch,float64(i))
	//	ch<-(4/float64(2*i+1))*float64(t)
		t*=-1
	}
//	close(ch)
}
func main(){
//	runtime.GOMAXPROCS(10)
	start:=time.Now()
	ch:=make(chan float64)
	go util(ch)
	pi:=0.0
	for i:=0;i<1000;i++ {
		pi+=<-ch

	}
	fmt.Println(pi)
	end:=time.Now()
	fmt.Println(end.Sub(start).String())
}
func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}