package main
import(
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
	fmt.Println("Starting the server")
	listener,err:=net.Listen("tcp","localhost:50000")

	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println(err.Error())
			return
		}
		go doServerStuf(conn)
	}
}

func doServerStuf(conn net.Conn){
	for{
		buf:=make([]byte,512)
		len,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(buf[:len]))
		var buffer bytes.Buffer
		buffer.Write(buf[:len])
		dec:=gob.NewDecoder(&buffer)
		var holder interface{}
		dec.Decode(&holder)
		fmt.Println(holder.(Sth).C)
	}
}