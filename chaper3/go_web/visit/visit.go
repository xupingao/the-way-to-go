package main
import(
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
func main(){
	url:="http://www.baidu.com"
	resp,err:=http.Head(url)
	checkError(err)
	fmt.Println(url+":",resp.Status)

	res,error:=http.Get(url)
	checkError(error)
	data,err:=ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got:%q\n",string(data))

	inputReader:=bufio.NewReader(os.Stdin)
	url,err=inputReader.ReadString('\n')
	url=strings.TrimSuffix(url,"/r/n")
	url=strings.TrimSpace(url)
	checkError(err)
	res,err=http.Get(url)
	checkError(err)
	data,err=ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(data))


}
func checkError(err error){
	if err!=nil{
		fmt.Println(err.Error())
	}
}