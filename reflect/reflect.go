package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

type myType struct {
	a int "this is tag of a"
	b int "this is tag of b"
}
type info struct {
	S1, s2, s3 string
}

func (i info) String() string {
	return fmt.Sprintf("[%v,%v,%v]", i.S1, i.s2, i.s3)
}
func main() {
	//tt := TagType{true, "Barak Obama", 1}
	//for i := 0; i < 3; i++ {
	//	refTag(tt, i)
	//}
	//
	//mt:=myType{1,2}
	//fmt.Println(reflect.TypeOf(mt).Field(1))
	//
	//f:=3.2
	//fmt.Println(reflect.TypeOf(f))
	//fmt.Println(reflect.ValueOf(f).Float())
	//type f32 float32
	//var val f32=3.2
	//v:=reflect.ValueOf(&val)
	//v=v.Elem()
	//fmt.Println(v.CanSet())
	//fmt.Println(v.Kind())
	//fmt.Println(v.Type())
	//v.SetFloat(2.0)

	var myinfo interface{} = &info{"yoyo", "comang", "hello"}
	rf := reflect.ValueOf(myinfo).Elem()
	fmt.Println(rf.Kind())
	fmt.Println(rf.Type())
	fmt.Println(rf)
	rf.Field(0).SetString("test")
	fmt.Println(rf.Field(0).Kind())
	for i := 0; i < rf.NumField(); i++ {
		fmt.Printf("field:%v,value:%v\n", i, rf.Field(i))
	}
	fmt.Println(rf.Method(0).Call(nil))
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
