package main

import "fmt"
type any interface{}
func main(){
	fmt.Println("start main()")
	fibFunc:=func(a any)(any,any){
		return a.([]int64)[0],[]int64{a.([]int64)[1],a.([]int64)[0]+a.([]int64)[1]}
	}
	generator:=lazyFunctionFactory(fibFunc,[]int64{0,1})
	for i:=0;i<10;i++{
		fmt.Println(generator().(int64))
	}

}

func integer()(<-chan int){
	out:=make(chan int)
	i:=0
	go func()(){
		for {
			out <- i
			i++
		}
	}()
	return out
}
func lazyFunctionFactory(genFunc func(s any)(any, any),start any)(func()any){
	out:=make(chan any)
	go func()(){
		var res any
		var state any=start
		for{
			res,state=genFunc(state)
			out<-res
		}
	}()
	return func()any{
		return <-out
	}
}
