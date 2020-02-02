package RW

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s1, s2 string
	//fmt.Scanln(&s1,&s2)
	//fmt.Println(s1,s2)
	fmt.Sscan("hello hi", &s1, &s2)
	fmt.Println(s1, s2)
	var (
		s string
		i int
		f float64
	)
	fmt.Sscanf("hello /123 /3.14", "%s /%d /%f", &s, &i, &f)
	fmt.Printf("%s/%d/%f \n", s, i, f)

	inputReader := bufio.NewReader(os.Stdin)
	val, _ := inputReader.ReadString('\n')
	fmt.Println(val)

}
