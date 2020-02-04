package main

import (
	"fmt"
	"strconv"
)
type divide struct{
	a,b int
}
type divideError struct{
	a,b int
}
func (de *divideError)Error()string{
	return strconv.Itoa(de.a)+strconv.Itoa(de.b)+"This is a error"
}

func (d *divide)div()(float32,error){
	if d.b==0{
		return 0,&divideError{d.a,d.b}
	}else{
		return float32(d.a)/float32(d.b),&divideError{}
	}
}

func myfunc1()(int,error){
	i:=9
	return i,fmt.Errorf("this is a test error %v",i)
}
func main(){
	if val,err:=myfunc1();err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(val)
	}
	d:=divide{3,0}
	res,err:=d.div()

	defer func()(){
		if e:=recover();e!=nil{
			fmt.Println("recover:",e)
		}
	}()
	if err!=nil{
		panic("ERROR occurred:"+err.Error())
	}else{
		fmt.Println(res)
	}
}
