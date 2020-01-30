package main

import (
	"fmt"
	"time"
)

func main() {
	start:=time.Now()
	for i:=1;i<40;i++{
		index,val:=fib(i)
		fmt.Printf("Index: %d,Value: %d\n",index,val)
	}
	end:=time.Now()
	delta:=end.Sub(start)
	fmt.Printf("take time: %s\n",delta)
}

func fib(i int)(index int,result int){
	index=i
	if i<=1{
		result=1
	}else{
		_,a:=fib(i-1)
		_,b:=fib(i-2)
		result=a+b
	}
	return
}