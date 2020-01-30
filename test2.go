// test project main.go helloworld
package main

import (
	u "./utils"
	"fmt"
)

type color int

const (
	rad = iota
	bulue
	black
)
const (
	d int = 3
	e int = 4
)

var c int

func init() {
	c = 1
}

var (
	hello int
	okey  int
)
type Rope string
func main() {
	fmt.Println("Hello World!")
	u.Fun()
	a, b := u.Add(1, 2)
	fmt.Println(a, b)
	testFun()
	fmt.Println(d, e)
	fmt.Print(black)
	a,b=b,a
	fmt.Printf("The val is %f\n",getArea(1))

}
func testFun() {
	fmt.Println(c)
	return
}
func getArea(val float64) float64{
	return u.Pi*val*val
}