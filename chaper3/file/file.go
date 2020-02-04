package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
type Page struct{
	Title string
	Body []byte
}
func (p *Page)save()(){
	err:=ioutil.WriteFile(p.Title,p.Body,0644)
	if err!=nil{
		panic(err.Error())
	}
}
func (p *Page)load()(){
	var inputErr error
	p.Body,inputErr=ioutil.ReadFile(p.Title)
	if inputErr!=nil{
		inputErr.Error()
	}
}

func main() {
	inputFile, inputErr := os.OpenFile("output.dat",os.O_WRONLY|os.O_CREATE, 0666)
	if inputErr != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()
	//inputReader := bufio.NewReader(inputFile)
	//for {
	//	inputString, readerError := inputReader.ReadString('\n')
	//	fmt.Println(inputString)
	//	if readerError == io.EOF {
	//		break
	//	}
	//}
	//var (
	//	i int
	//	f float32
	//	s string
	//)
	//fmt.Fscanln(inputFile, &i, &f, &s)
	//fmt.Println(i, f, s)

	//outputWriter:=bufio.NewWriter(inputFile)
	//outputWriter.WriteString("HELLO WORLD")
	//outputWriter.Flush()
	//p:=Page{"p1",[]byte("helloworld")}
	//p.save()
	p:=Page{"p1",nil}
	p.load()
	fmt.Println(string(p.Body))
	oldFile,_:=os.Open("p1")
	defer oldFile.Close()
	newFile,_:=os.Create("p2")
	defer newFile.Close()
	io.Copy(newFile,oldFile)
}
