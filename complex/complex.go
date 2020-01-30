package main

import "fmt"

func main(){
	var c1 complex64=5+10i
	fmt.Printf("The value is %v",c1)
	fmt.Printf("Real：%g,imag:%g",real(c1),imag(c1))
}

