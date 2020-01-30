package main
import (
	f "../function"
	"fmt"
	"time"
)
type fun_t func()()
var fun1 fun_t
func init(){

	fun1= func()(){
		fmt.Println("hello")
	}
}
func main(){
	start:=time.Now()
	fmt.Println("this is  a review")
	f.Fun1(1,2,3)
	fun1()
	var a int=3
	fun2(&a)
	defer fmt.Println("end of func")
	fmt.Println(a)
	s1:=[]int{1,2,3}
	fmt.Println(sum(s1...))
	plus_fun:=factory(1)
	fmt.Println(plus_fun(3))

	s2:=[]int{1,2,3,4}
	s3:=[]int{1,2,3,4,5}
	s2=append(s2,s3...)
	s3[0]=100
	fmt.Println(s2[4])

	s4:=make([]int,1,3)
	fmt.Println(len(s4))

	var s string="hello"
	s5:=[]byte(s)
	s5[0]='p'
	fmt.Println(s)
	fmt.Println(s)
	end:=time.Now()
	delete:=end.Sub(start)
	fmt.Println(delete)

}

func fun2(a *int)(){
	*a--
}
func sum(n ...int)(sum int){
	if len(n)==0{
		return 0
	}else {
		for _, num := range (n) {
			sum += num
		}
	}
	return
}
func plus(a,b int)(int){
	return a+b
}
func callback(fun func(int,int)(int))(){
	fmt.Println(fun(1,2))
}

func factory(a int)(fun func(int)(int)){
	fun=func(b int)(sum int){return a+b}
	return
}