package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Fprintln(os.Stdout,"helloWorld")

	bw :=bufio.NewWriter(os.Stdout)
	bw.WriteString("hi\n")
	fmt.Fprintln(bw,"hello")
	bw.Flush()

}
