package main

import "fmt"

func main(){
	var arr [3]int
	for i,_:=range arr {
		arr[i]=i*2
	}
	var arr2 *[3]int
	arr2=&arr
	for i:=range arr2{
		fmt.Print(arr2[i],"")
	}
	fmt.Println()
	arr3:=arr
	arr3[0]=100
	fmt.Println(arr[0])
	var arr4 [3]int
	fun(arr4)
	fmt.Println(arr4[0])
	var fib [50]int
	fib[0],fib[1]=1,1
	for i:=2;i<len(fib);i++{
		fib[i]=fib[i-1]+fib[i-2]
		fmt.Print(fib[i]," ")
	}

	var arr5=[5]int{1,2,3}
	fmt.Println(arr5[3])
	var arr6=[3]string{1:"hhe",2:"yiyi"}
	fmt.Println(arr6[1])
	print3([3]int{1,2,3})
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice

	for i:=0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}

}
func fun(arr[3]int)(){
	for i:=range arr{
		arr[i]=10
	}
}
func print3(arr [3]int)(){
	for i:= range arr{
		fmt.Print(arr[i]," ")
	}
	fmt.Println()
}
