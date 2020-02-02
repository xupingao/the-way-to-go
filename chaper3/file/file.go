package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputErr := os.Open("E:\\Go-Project\\src\\test2\\chaper3\\file\\in.text")
	if inputErr != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Println(inputString)
		if readerError == io.EOF {
			return
		}
	}
	var (
		i int
		f float32
		s string
	)
	fmt.Fscanln(inputFile, &i, &f, &s)
	fmt.Println(i, f, s)
}
