package main

import "fmt"

func main(){
	var str1 string="\u00ff\u754c"
	for i,c:=range str1{
		fmt.Printf("%d: %c\n",i,c)
	}
	c:=[]int32(str1)
	r:=[]rune(str1)
	for i,v:=range c{
		fmt.Printf("%d: %d\n",i,v)
	}
	for i,v:=range r{
		fmt.Printf("%d: %c\n",i,v)
	}
	var str2="teststr"
	var slice1 []byte=[]byte(str2)
	slice1[0]='p'
	str2=string(slice1)
	fmt.Println(str2)
	str3:="str3"
	str4:="str2"
	fmt.Println(compare([]byte(str3),[]byte(str4)))
	array:=[4]byte{'a','b','c','d'}
	slice2:=array[0:2]
	slice2[0]='q'
	slice2=append(slice2,'h')
	fmt.Println(string(slice2))
	fmt.Println(string([]byte(array[:])))
	var slice3 []int=[]int{1,2,3}
	fmt.Println(len(slice3))
	slice3=append(slice3,make([]int,3)...)
	fmt.Println(len(slice3))
	slice4:=[]int{1,2,5}
	slice4=append(slice4[:2],append([]int{3,4},slice4[2:]...)...)
	for _,i:=range slice4{
		fmt.Printf("%d  ",i)
	}
	fmt.Println()
	arr:=[3]int{}
	sli:=arr[:]
	sli=append(sli,1)
	fmt.Println(len(arr)," ",len(sli))
	arr2:=[4]int{}
	copy(arr2[:],sli)
	for _,i:=range arr2{
		fmt.Printf("%d  ",i)
	}
	str:="1234"
	a,b:=split(str,2)
	fmt.Println(a," ",b)
	array2:=[5]int{5,4,3,2,1}
	sort(array2[:])
	for _,i:=range array2{
		fmt.Printf("%d  ",i)
	}
}

func compare(a,b []byte)(int){
	if len(a)>len(b){
		return 1
	}else if len(a)<len(b){
		return -1
	}else{
		for i:=0;i<len(a);i++{
			if a[i]>b[i]{
				return 1
			}
			if a[i]<b[i]{
				return -1
			}
		}
		return 0
	}
}

func split(str string,i int )(a,b string){
	a=string(str[:i])
	b=string(str[i:])
	return
}
func sort(arr []int)(){
	n:=len(arr)
	for i:=0;i<n-1;i++{
		flag:=false
		for j:=n-1;j>i;j--{
			if arr[j]<arr[j-1]{
				t:=arr[j]
				arr[j]=arr[j-1]
				arr[j-1]=t
				flag=true
			}
		}
		if !flag{
			break
		}
	}
}