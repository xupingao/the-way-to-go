package main

import (
	"fmt"
	"sort"
)

type IntVector []int

func (iv IntVector) Less(i, j int) bool {
	return iv[i] < iv[j]
}
func (iv IntVector) Len() int {
	return len(iv)
}
func (iv IntVector) Swap(i, j int) {
	iv[i], iv[j] = iv[j], iv[i]
}

type stu struct {
	grade int
	name  string
}
type StuList []stu

func (s StuList) Len() int {
	return len(s)
}
func (s StuList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StuList) Less(i, j int) bool {
	return s[i].grade < s[j].grade
}
func (s StuList) print() {
	for _, v := range s {
		fmt.Printf("gread:%v,name:%v", v.grade, v.name)
	}
}

type miner interface {
	less(i, j int) bool
	len() int
}

func getMin(m miner) (a int) {
	a = 0
	for i := 1; i < m.len(); i++ {
		if m.less(i, a) {
			a = i
		}
	}
	return
}
func (s StuList) less(i, j int) bool {
	return s[i].grade < s[j].grade
}
func (s StuList) len() int {
	return len(s)
}
func main() {
	iv := IntVector{5, 7, 3, 4, 1}
	sort.Sort(iv)
	fmt.Println(iv[0])

	sv := StuList{stu{100, "xiaoming"}, stu{99, "xiaohong"}, stu{96, "xiaobao"}}
	sort.Sort(sv)
	sv.print()

	fmt.Println(sv[getMin(sv)])
}
