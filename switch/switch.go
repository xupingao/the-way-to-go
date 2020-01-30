package main
import(
	"fmt"
)
func main(){
	val:=1
	switch val{
	case 1:  {
		fmt.Println("1")

	}
	fallthrough
	case 2/1: 	fmt.Println("2")
	fallthrough
	default:
		fmt.Println("3")
	}
	b:=1
	switch {   //switch true
	case 1+2==3:fmt.Println("1+2=3")
	fallthrough
	case b==1 : fmt.Println("b==1")
	case !(b!=1): fmt.Println("!b1=1")
	}
	month:=1
	switch{
	case month>=1&&month<=3:fmt.Println()
	}
}
