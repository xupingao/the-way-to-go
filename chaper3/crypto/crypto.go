package main
import(
	"crypto/sha256"
	"fmt"
	"io"
)
func main(){
	s256:=sha256.New()
	io.WriteString(s256,"helloworld")
	fmt.Fprintln(s256,"hello")
	fmt.Println(s256.Sum([]byte("helloworld")))
}