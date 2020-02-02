package main

import (
	"./stack"
	"fmt"
	"strconv"
	"strings"
)

type inner struct {
	a, b int
}

func (i *inner) show() {
	fmt.Println(i.a)
}

type outer struct {
	inner
	int
	a, b int
}
type T struct {
	ptr_p *person
}

type person struct {
	first  string
	second string
}
type trible struct {
	a, b, c int
}

func (t *trible) sum() {
	t.c = t.a + t.b
}

type people person

func Upper(p *person) {
	p.first = strings.ToUpper(p.first)
	p.second = strings.ToUpper(p.second)
}

func print_p(p *person) {
	fmt.Println(p.first, " ", p.second)
}
func print_p2(p person) {
	fmt.Println(p.first, " ", p.second)
}
func (p person) print() {
	fmt.Println(p.first, " ", p.second)
}
func (_ *person) print_p() {
	fmt.Println("hi！")
}

type iVector []int

func (arr iVector) sum() {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	fmt.Println(sum)
}

type integer int

func (_ integer) printself() {
	fmt.Println("yoyoyo")
}
func (i *integer) String() string {
	return strconv.Itoa(int(*i + 1))
}

type myfun func()

func (_ myfun) pr() {
	fmt.Println("hello")
}

type engine interface {
	start()
	stop()
}
type car struct {
	engine
	wheelConut int
}

func (c *car) GoToWorkIn() {
	c.start()
	c.stop()
}
func (c *car) numberOfWheels() {
	fmt.Println(c.wheelConut)
}

type mercedes struct {
	car
}
type Log struct {
	msg string
}

func (l *Log) add(s string) {
	l.msg += s
}
func (l *Log) String() string {
	return l.msg
}

type obj struct {
	int
	log *Log
}

func (o *obj) Log() *Log {
	return o.log
}

type obj2 struct {
	int
	Log
}

//func (o *obj2)String()string{
//	return o.Log.msg
//}
func (p integer) get() int {
	return int(p)

}
func f(i int) {}

type twoInt struct {
	a, b int
}

func (t *twoInt) String() string {
	return "(" + strconv.Itoa(t.a) + "," + strconv.Itoa(t.b) + ")"
}
func main() {

	//p1 := person{"hello", "ok"}
	//Upper(&p1)
	//print_p(&p1)
	//p2 := &person{"test", "Kobe"}
	//p3 := new(person)
	//Upper(p2)
	//print_p(p2)
	//Upper(p3)
	//print_p(p3)
	//
	//p4 := people{"hello", "enen"}
	//p5 := person(p4)
	//print_p(&p5)
	//persons := [2]person{p1, *p2}
	//for _, p := range (persons) {
	//
	//	p.first = ""
	//	print_p2(p)
	//}
	//print_p(&p1)
	//fmt.Println(unsafe.Sizeof(p1))
	//ts := test.NeWTest(1)
	//fmt.Println(ts.A)
	//m := make(map[int]int, 3)
	//fmt.Println(len(m))
	//var o outer = outer{inner{1, 2}, 3, 4, 5}
	//
	//fmt.Println(o.inner.a)
	//p1.print()
	//(*p2).print()
	//p1.print_p()
	//p2.print_p()
	//p6 := person(p4)
	//p6.print_p()
	//
	//arr := iVector{1: 1, 2: 2, 3: 3}
	//arr.sum()
	//var integ integer = 1
	//integ.printself()
	//var mf myfun = func() () {}
	//mf.pr()
	//t := trible{1, 2, 0}
	//t.sum()
	//fmt.Println(t.c)
	//
	//per:=test.NewPerson("hello","heiheui")
	//fmt.Println(per.GetFirstName())
	//o.show()
	o := obj{1, &Log{"hello"}}
	o.Log().add("\n kobe")
	fmt.Println(o.Log())
	o2 := &obj2{1, Log{"hai,baby"}}
	o2.add("\n my prety")
	fmt.Println(o2)
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
	i := integer(1)
	fmt.Println(i.String())
	f(i.get())

	t := &twoInt{1, 2}
	fmt.Println(t)

	stack := stack.NewStack(5)
	stack.Push(1)
	stack.Push(3)
	stack.Push(5)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}
