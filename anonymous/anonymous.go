package main

import "fmt"

func main(){
	f:=func(a int)(int){fmt.Println(a)
	return a}
	f(2)
	var power2 func(int)(int)=makePowerFunc(2)
	fmt.Println(power2(2))
	var setter func(int)()
	var getter func()(int)
	setter,getter=counter()
	setter(10)
	fmt.Println(getter())
	addBmp:=makeAddsuffix(".bmp")
	fmt.Println(addBmp("test"))
	fibFun:=fib()
	for i:=0;i<20;i++{
		fmt.Print(fibFun()," ")
	}
}

func makePowerFunc(power int)(func(int)(int)){
	return func(base int)(result int){
		result=1
		for i:=0;i<power;i++{
			result*=base
		}
		return
	}
}
func counter()(func(inc int)(),func()(int)){
	x:=0
	return func(inc int)(){
		x=inc
	},func()(int){
		return x
	}
}
func makeAddsuffix(suf string)(func(string)(string)){
	return func(str string)(res string){
		res=str+suf
		return
	}
}
func fib()(func()(int)) {
	x,y:=1,1
	index:=0
	return func()(int){
		index++
		switch{

			case index<=2:return 1
			case index==3:{
				y=2
				return 2
			}
			default:{
				z:=y
				y+=x
				x=z
				return y
			}
		}

	}
}