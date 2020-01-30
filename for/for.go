package main
import(
	"fmt"
	"strings"
)
func main(){
	for i:=1;i<3;i++{
		fmt.Printf("the first for %d\n",i)
	}
	for i,j:=1,10;i<j;i,j=i+1,j-1{
		fmt.Printf("the second for %d\n",i)
	}
	for i:=1;i<3;i++{
		for j:=1;j<2;j++{
			fmt.Printf("the tird for %d\n",j)
		}
	}
	str1:="go is a beautifu language"
	for i:=0;i<len(str1);i++{
		fmt.Printf("The char porstition %d is %c\n",i+1,str1[i])
	}
	counter:=1
	start:
		counter++
		if counter<15{
			goto start
		}else{
			fmt.Printf("%d\n",counter)
		}
	for i:=1;i<5;i++{
		fmt.Println(strings.Repeat("G",i))
	}
	for i:=1;i<101;i++{
		switch {
		case i%3==0&&i%5==0:fmt.Print("hehe ")
		case i%3==0 :fmt.Print("yoyo ")
		case i%5==0:fmt.Print("keke ")
		default:
			fmt.Print(i," ")

		}
	}
	fmt.Println()
	t:=0
	for t<=2{
		t++
		fmt.Print(t)
	}
	for pos,c :=range str1{
		fmt.Printf("character on position %d is %c \n",pos,c)
	}
	str2:="charse:我是好人"
	for pos,c:= range str2{
		fmt.Printf("character on postition %d is %c \n",pos,c)
	}
	fmt.Println("idex int(rune) rune char bytes")
	for index,rune:=range str2{
		fmt.Printf("%-2d    %d      %U  '%c'  %X \n",index,rune,rune,rune,[]byte(string(rune)) )
	}
}