package main

import "fmt"

type car struct {
	name  string
	brand string
}
type cars []car

func (cs cars) process(fun func(*car)) {
	for i, _ := range cs {
		fun(&cs[i])
	}
}
func (cs cars) screen(f func(*car) bool) cars {
	res := cars{}
	cs.process(func(c *car) {
		if f(c) {
			res = append(res, *c)
		}
	})
	return res
}
func MakeSortedAppender(brand []string) (map[string]cars, func(*car)) {
	sortedCars := map[string]cars{}
	for i, _ := range brand {
		sortedCars[brand[i]] = cars{}
	}
	fun := func(c *car) {
		sortedCars[c.brand] = append(sortedCars[c.brand], *c)
	}
	return sortedCars, fun
}
func main() {
	cs := cars{car{"gogo", "test"}, car{"hoho", "bmw"}}
	newcs := cs.screen(func(c *car) bool {
		return c.brand == "bmw"
	})
	fmt.Println(len(newcs), newcs[0].name)
	sortedCars, AppendFunc := MakeSortedAppender([]string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"})
	cs.process(AppendFunc)
	fmt.Println(sortedCars["bmw"])
	c2 := cars{car{"gogo", "test"}, car{"comeon", "bmw"}}
	c2.process(AppendFunc)

	fmt.Println(sortedCars["bmw"])
}
