package main
import(
"../rpc_objects"
"fmt"
"log"
"net/rpc"
)
const serverAddress="localhost"
func main(){
	client,err:=rpc.DialHTTP("tcp",serverAddress+":1234")
	if err!=nil{
		log.Fatal(err)
	}
	args:=&rpc_objects.Args{13,4}
	var reply int
	err=client.Call("Args.Multiply",args,&reply)

	if err!=nil{
		log.Fatal("Args error",err)
	}
	fmt.Println(reply)
}
