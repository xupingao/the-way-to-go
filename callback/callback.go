package main

import "fmt"

func main(){
	a(b)
}
func a(f func(int)){
	a:=1
	f(a)
}
func b(i int){
	fmt.Println(i)
}