package main

import "fmt"

func main(){
	var arr1 [10]int
	var slice1 []int=arr1[1:5]//[start,end)
	for _,val:=range slice1{
		fmt.Print(val," ")
	}
	fmt.Println()
	slice2:=[]int{1,2,3,4,5}
	slice2=slice2[1:3]
	for _,val:=range slice2{
		fmt.Print(val," ")
	}
	fmt.Println()
	slice2=slice2[0:cap(slice2)]
	for _,val:=range slice2{
		fmt.Print(val," ")
	}
	fmt.Println(slice2)
	slice3:=make([]int,5,10)
	for _,val:=range slice3{
		fmt.Print(val," ")
	}
	slice4:=new([100]int)[0:50]
	for _,val:=range slice4{
		fmt.Print(val," ")
	}
	var slice5 []int=make([]int,4)
	for i:=range slice5{
		slice5[i]=i*i
	}
	seasons:=[]string{"Sprig","Summer","Autumn","Winter"}
	for _,season:=range seasons {
		fmt.Println(season)
	}
	slice6:=[]int{1,2,3}
	fmt.Println(sum(slice6...))
	slice7:=[]int{1,2,3}
	slice8:=make([]int,10)
	copy(slice8,slice7)
	slice9:=[]int{1,2,3}
	slice9=append(slice9,4,5,6)
	for _,val:=range slice9{
		fmt.Print(val," ")
	}
	slice9=append(slice9,slice7...)
	for _,val:=range slice9{
		fmt.Print(val," ")
	}
	slice10:=[]int{1,2,3,4,5,6,7,8,9,10}
	slice11:=screen(func(a int)(bool){return a%2==0},slice10)
	for _,val:=range slice11{
		fmt.Print(val," ")
	}
}
func sum(slice...int)(sum int){
	for _,val:=range slice{
		sum+=val
	}
	return
}
func screen(fun func(int)(bool),slice []int)(res []int){
	for _,val:=range slice{
		if fun(val){
			res=append(res,val)
		}
	}
	return
}