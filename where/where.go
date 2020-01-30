package main

import (
	"fmt"
	"log"
	"runtime"
)

func main(){
	fun1()
}

func fun1(){
	where:=func(){
		_,file,line,_:=runtime.Caller(1)
		log.Printf("%s:%d",file,line)
	}
	where()
	fmt.Println("dosth")
	where()
}
