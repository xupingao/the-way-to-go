package main
import(
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead=25
func main(){
	flag.Parse()
	if flag.NArg()!=2{
		panic("usage:host port")
	}
	hostAndPort:=fmt.Sprintf("%s:%s",flag.Arg(0),flag.Arg(1))
	listener:=initServer(hostAndPort)
	for{
		conn,err:=listener.Accept()
		checkError(err)
		go connectionHandler(conn)
	}
}
func initServer(hostAndPort string)*net.TCPListener{
	serverAddr,err:=net.ResolveTCPAddr("tcp",hostAndPort)
	checkError(err)
	listener,err:=net.ListenTCP("tcp",serverAddr)
	checkError(err)
	println("Listening to:",listener.Addr().String())
	return listener
}
func connectionHandler(conn net.Conn){
	connFrom:=conn.RemoteAddr().String()
	println("Connection from:",connFrom)
	sayHello(conn)
	for{
		var ibuf []byte=make([]byte,maxRead+1)
		length,err:=conn.Read(ibuf)
		switch err{
		case nil:handleMsg(length,err,ibuf)
		case syscall.EAGAIN:continue
		default: goto DISCONNECT
		}
	}
	DISCONNECT:
		err:=conn.Close()
		checkError(err)
		println("Closed connection:",connFrom)
}
func sayHello(conn net.Conn){
	obuf:=[]byte("lets go!\n")
	_,err:=conn.Write(obuf)
	checkError(err)
}
func checkError(err error){
	if err!=nil{
		fmt.Println(err.Error())
	}
}
func handleMsg(length int,err error,msg []byte){
	if length>0{
		print("<",length,":")
		for i:=0;;i++{
			if msg[i]==0{
				break
			}
			fmt.Printf("%c",msg[i])
		}
		print(">")
	}
}