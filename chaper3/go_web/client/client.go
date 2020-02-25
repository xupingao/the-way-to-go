package main

import (

	"bytes"
	"encoding/gob"
	"fmt"
	"net"

)
type Sth struct{
	A,B int
	C string
}
func main(){
	conn,err:=net.Dial("tcp","localhost:50000")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	load:=Sth{1,2,"hello"}
	var buf bytes.Buffer
	g:=gob.NewEncoder(&buf)
	g.Encode(load)
	fmt.Println(buf.String())
	conn.Write(buf.Bytes())
	//inputReader:=bufio.NewReader(os.Stdin)
	//str,_:=inputReader.ReadString('\n')
	//trimmedStr:=strings.Trim(str,"\r\n")
	//for{
	//	fmt.Println("Type Q to quit")
	//	input,_:=inputReader.ReadString('\n')
	//	input=strings.Trim(input,"\r\n")
	//	if input=="Q"{
	//		return
	//	}
	//	_,err=conn.Write([]byte(trimmedStr+" Says:"+input))
	//}
}
