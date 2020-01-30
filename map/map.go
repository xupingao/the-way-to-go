package main

import (
	"fmt"
	"sort"
)

func main(){
	var m map[int]string=map[int]string{}
	m[1]="hello"
	fmt.Println(m[1])

	var m1 =map[int]string{1:"kobe",2:"wade"}
	fmt.Println(m1[2])
	for key,val:=range(m1){
		fmt.Printf("key:%d,value:%s\n",key,val)
	}
	m2:=make(map[int]string)
	m2[1]="hello"
	fmt.Println(m2[1])
	delete(m1,2)
	_,ok:=m1[2]
	fmt.Println(ok)
	var (
		barval=map[string]int{"alpha":34,"bravo":56,
			"charlie":23}
	)
	for k,v:=range(barval){
		fmt.Printf("key:%v,balue:%v / ",k,v)
	}
	keys:=make([]string,len(barval))
	i:=0
	for key,_:=range(barval){
		keys[i]=key
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	for _,k:=range(keys){
		fmt.Printf("key:%v,value:%v / ",k,barval[k])
	}
}
