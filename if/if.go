package main
import(
	"fmt"
	"math"
	"os"

	"strconv"
)

func main(){
	flag:=true
	if (flag==true){
		fmt.Println("The flag is true")
	}else {
		fmt.Println("The flag is false")
	}
	fmt.Println(Abs(-1))
	fmt.Println(IsGreater(3,2))
	first,second:=5,10
	if first>20{
		fmt.Printf(">20")
	}else if first>10{
		fmt.Printf("20>first>10")
	}else{
		fmt.Println("<10")
	}
	if val:=10;second>val {
		fmt.Println("the second is smaller")
	}else{
		fmt.Println("The second is larger")
	}
	_,err:=os.Open("heihei")
	if err!=nil{
		//os.Exit(1)
	}
	multiReturnTest()
	fmt.Println("test")
	if f,ok:=sqrt(25);ok{
		fmt.Println(f)
	}
}
func Abs(val int) int{
	if val>0 {
		return val
	}else{
		return - val
	}
}
func IsGreater(val1,val2 int) bool{
	if val1>val2{
		return true
	}else{
		return false
	}
}
func GreaterThen10(val int) bool{
	if k:=10;val>k{
		return true
	}else{
		return false
	}
}
func multiReturnTest(){
	str:="test"
	an,err:=strconv.Atoi(str)
	if err!=nil{
		fmt.Printf("%s is illegal\n",str)
		return   //结束当前函数
 		//os.Exit(1) //直接结束当前程序
	}
	fmt.Println(an)
}
func sqrt(f float64)(float64,bool){
	if f<0{
		return -1,false
	}else{
		return math.Sqrt(f),true
	}
}