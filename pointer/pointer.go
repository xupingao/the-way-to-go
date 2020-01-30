package main
import (
	"fmt"
)
func main(){
	a:=1
	var b *int=&a
	fmt.Println(b)
	fmt.Printf("%p\n",&a)
}