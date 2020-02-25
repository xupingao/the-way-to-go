package main

import (
	"fmt"
	"io"
	"log"
)

func main(){
	a:=3
	fmt.Println(c(&a))
	//a()
	//b("HELLO")
}
func trace(str string)(s string){
	fmt.Printf("entering:%s\n",str)
	return str
}
func untrace(str string){
	fmt.Printf("leaving:%s\n",str)
}

func a(){
	defer untrace(trace("a"))
	fmt.Printf("running a\n")
}
func b(s string)(n int ,err error){
	defer func(){
		log.Printf("func(%q) = %d,%v",s,n,err)
	}()
	return 7,io.EOF
}

func c(a *int)int{
	defer func(){
		*a++
	}()
	return *a
}