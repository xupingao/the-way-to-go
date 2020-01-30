package function

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
)

func main(){
	fmt.Println(Fun1(fun2()))
	fun3()
	a:=1
	callByRef(&a)
	fmt.Println(a)
	fmt.Println(funWithReturnName(1,2))
	var reply int
	multiply(2,3,&reply)
	fmt.Println(reply)
	greeting("hello:","yo","he")
	s:=[]int{1,2,3,4,0}
	fmt.Println(getMin(s...))

	fmt.Println(intMulfloat(vals{1,1.1}))
	deferTest()
}
func Fun1(a,b,c int)(int){
	return a+b+c
}
func fun2 ()(a,b,c int){
	return 1,2,3
}
type fun_type func()()

var fun3 fun_type=fun4
func fun4(){
	fmt.Println("this is a func")
}

func callByRef(a *int){
	*a++
}
func funWithReturnName(a,b int)(sum int) {
	sum=a+b
	return
}
func sqrt(a float64)(val float64,err error){
	if a<0{
		val=math.NaN()
		err=errors.New("Error!")
	}else{
		val=math.Sqrt(a)
	}
	return
}
func multiply(a,b int ,c *int){
	*c=a*b
}

func greeting(prex string,strs ...string){
	fmt.Print(prex," ")
	for _,str:=range strs{
		fmt.Print(str ,",")
	}
}
func getMin(s ...int)(min int){
	min=math.MaxInt32
	for _,val:=range s{
		if val<min{
			min=val
		}
	}
	return
}
type vals struct{
	i int
	f float64
}
func intMulfloat(in vals) (result float64){
	result=float64(in.i)*in.f
	return
}
func deferTest(){
	defer printSth()
	fmt.Println("this is a string")
	return
}
func printSth(){
	fmt.Println("this is a string frm printsth")
}