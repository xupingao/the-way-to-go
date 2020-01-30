package main

import (
	"./package1"
	"container/list"
	"fmt"
	"math/big"
	"sync"
)
type test struct{
	m sync.Mutex
	v int
}
func main(){
	l1:=list.New()
	l1.PushBack(101)
	l1.PushBack(102)
	l1.PushBack(103)
	t:=new(test)
	t.v=0
	operator(t)
	operator(t)

	b1:=big.NewInt(10)
	b2:=big.NewInt(20)
	b1.Add(b1,b2)
	fmt.Println(b1.String())
	package2.Function1()
	fmt.Println(package2.Function3())
}

func operator(t *test){
	t.m.Lock()
	t.v++
	t.m.Unlock()
}