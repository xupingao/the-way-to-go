package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

const NCPU=4
func main(){
	runtime.GOMAXPROCS(NCPU)
	vector:=make([]int,100)
	for i,_:=range(vector){
		vector[i]=rand.Intn(200)
	}
	Mine(vector)
}
func Mine(v []int)int{
	sem:=make(chan int,NCPU)
	length:=len(v)/NCPU
	for i:=0;i<NCPU;i++{
		go FindMine(v[i*length:(i+1)*length],sem)
	}
	min:=<-sem
	for i:=1;i<NCPU;i++{
		if i:=<-sem;i<min{
			min=i
		}
	}
	fmt.Println(min)
	return min
}
func FindMine(v []int,sem chan int){
	min:=v[0]
	for i,v:=range(v){
		if i==0{
			continue
		}else{
			if v<min{
				min=v
			}
		}
	}
	sem<-min
}
