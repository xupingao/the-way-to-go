package main

import (
	"flag"
	"fmt"
	"time"
)
var procs=flag.Int("n",2,"number of CPU cores to use")
func main(){
	//flag.Parse()
	//runtime.GOMAXPROCS(*procs)
	//fmt.Println(*procs)

	//fmt.Println("In main()")
	//go A()
	//go B()
	//time.Sleep(10*1e9)
	//fmt.Println("end of main()")
	fmt.Println("In main()")
	vector:=[]int{1,2,3,4,5,6}
	go search(vector[:len(vector)/2],3)
	go search(vector[len(vector)/2:],3)
	time.Sleep(10*1e9)
	fmt.Println("end of main()")
}

func A()(){


	fmt.Println("start function A")
	time.Sleep(5*1e9)
	fmt.Println("end function A")
}
func B()(){


	fmt.Println("start function B")
	time.Sleep(2*1e9)
	fmt.Println("end function B")
}
func search(vector []int,val int)(){
	for _,v:=range(vector){
		if v==val{
			fmt.Println(v)
			break
		}
	}
}