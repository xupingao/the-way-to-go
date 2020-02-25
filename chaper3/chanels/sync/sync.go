package main

import (
	"fmt"
	"time"
)

type person struct{
	name string
	grade int
	funcChanel chan func(*person)
}
func Process(p *person){
	for{
		fun:=<-p.funcChanel
		fun(p)
	}
}
func NewPerson(name string,grade int)(*person){
	funcCh:=make(chan func(*person))
	p:=person{name,grade,funcCh}
	go Process(&p)
	return &p
}
func (p *person)String()string{
	return fmt.Sprintf("I miss you %s",p.name)
}
func main(){
	p:=NewPerson("kobe",81)
	p.funcChanel<-func(p *person){
		fmt.Println(p)
	}
	time.Sleep(2*1e9)
}