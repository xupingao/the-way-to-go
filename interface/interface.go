package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}
type Rectangle struct {
	length, width float32
}
type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}
func (s Square) Area() float32 {
	return s.side * s.side
}

func operator(s Shaper) {
	fmt.Println(s.Area())
}

type PeriInterface interface {
	Perimeter() float32
}

func (s *Square) Perimeter() float32 {
	return 4 * s.side
}

type mystring string

func typeSwitch(any interface{}) {
	switch any.(type) {
	case *mystring:
		fmt.Println("this is mystring")
	}
}

type Object interface {
	showSelf()
}
type node struct {
	data Object
	next *node
}
type linkList struct {
	head, tail *node
}
type List interface {
	showList()
	add(Object)
}

func showList(l List) {
	l.showList()
}

type student struct {
	grade int
	name  string
}

func (stu *student) showSelf() {
	fmt.Printf("name:%v,grade:%v \n", stu.name, stu.grade)
}
func (l linkList) showList() {
	for ptr := l.head; ptr != nil; ptr = ptr.next {
		ptr.data.showSelf()
	}
}
func (l *linkList) add(obj Object) {
	n := node{obj, nil}
	if l.head == nil {

		l.head = &n
		l.tail = &n
	} else {
		l.tail.next = &n
		l.tail = l.tail.next
	}

}

type element interface{}

type if1 interface {
	fun1()
}
type if2 interface {
	fun2()
}
type sub struct {
}

func fun(i if1) {
	i.(if2).fun2()
}
func (s sub) fun1() {}
func (s sub) fun2() {}

type functionI interface {
	flag()
}
type imple struct {
	a int
}

func (i *imple) flag() {}
func operation(i obj) {
	(i.(*imple)).a++
}

type obj interface{}
type integer struct {
	int
}

func main() {
	s := &Square{1.5}
	var shaper Shaper = s
	fmt.Println(shaper.Area())
	shapers := make([]Shaper, 10)
	c := Circle{1.5}
	shapers[2] = &c
	shapers[0] = s
	shapers[1] = &Rectangle{2.4, 5}
	fmt.Println(shapers[0].Area())
	fmt.Println(shapers[1].Area())
	fmt.Println(shapers[2].Area())
	for _, v := range shapers {
		switch v.(type) {
		case *Square:
			fmt.Println("this is a square")
		case *Rectangle:
			fmt.Println("this is a rectangle")
		case *Circle:
			fmt.Println("this is a circle")
		default:
			fmt.Println("this is not a shap")

		}
	}
	fmt.Println(shapers[0].(*Square))
	operator(s)
	var peri PeriInterface = s
	fmt.Println(peri.Perimeter())

	elements := make([]element, 3)
	elements[0] = shapers
	new_shapers := elements[0].([]Shaper)
	fmt.Println(new_shapers[0].Area())
	l := linkList{}
	var list List = &l
	list.add(&student{99, "lili"})
	list.add(&student{99, "lili"})
	list.add(&student{99, "lili"})
	fmt.Println(list)
	list.showList()

	subject := sub{}
	var i1 if1
	var i2 if2
	i1 = subject
	i2 = i1.(if2)
	i1 = i2.(if1)
	i1.fun1()
	fun(i1)

	objects := []obj{1, "hello"}
	mf := func(o obj) {
		switch (o).(type) {
		case int:
			o = (o).(int) * 2
		case string:
			o = (o).(string) + (o).(string)
		case *integer:
			o.(*integer).int = 2
		}
	}
	map_operator(objects, mf)
	fmt.Println(objects[0])
	i := imple{1}
	i_interface := []obj{&i}
	//	var i_interface obj=&i
	operation(i_interface[0])
	fmt.Println(i.a)

	var ot obj = &integer{1}
	mf(ot)
	fmt.Println(ot)
}

func map_operator(objects []obj, fun func(obj)) {
	for i, _ := range objects {
		fun(objects[i])
	}
}
