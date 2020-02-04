package main

import (
	"bufio"
	"flag" // command line option parser
	"fmt"
	"io"
	"os"
	"strconv"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)
func cat(r *bufio.Reader){
	i:=1
	for{
		buf,err:=r.ReadBytes('\n')
//		fmt.Println(string(buf))
		if err==io.EOF{
			break

		}
		if *NewLine{
			prefix:=strconv.Itoa(i)+": "
			buf=append([]byte(prefix),buf...)
			i++
		}
		fmt.Fprintf(os.Stdout,"%s",buf)
	}
	return
}
func main() {
	//flag.PrintDefaults()
	//flag.Parse() // Scans the arg list and sets up flags
	//var s string = ""
	//for i := 0; i < flag.NArg(); i++ {
	//	if i > 0 {
	//		s += " "
	//		if *NewLine { // -n is parsed, flag becomes true
	//			s += Newline
	//		}
	//	}
	//	s += flag.Arg(i)
	//}
	//os.Stdout.WriteString(s)

	flag.Parse()
	if flag.NArg()==0{
		cat(bufio.NewReader(os.Stdin))
	}
	for i:=0;i<flag.NArg();i++{
		f,err:=os.Open(flag.Arg(i))
		if err!=nil{
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}