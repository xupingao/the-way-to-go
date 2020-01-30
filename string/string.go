package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)
func main(){
	fmt.Print("\r\t\u1000\\\n")
	var str1 string
	fmt.Println(len(str1))  //默认为长度为0的空字符串
	var str2="hello"
	fmt.Printf("str2[1]:%c\n",str2[1]) //l
	//fmt.Println(&str2[2])
	//获取字符串中某个字节的地址的行为是非法的，例如：&str[i]。

	str3:="hel"+"lo"
	str3+=" world!"
	fmt.Println(str3)
	str4:=fmt.Sprintf("Get a string by %s","Sprinf()")
	fmt.Println(str4)

	str5 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str5))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str5))
	str6 := "asSASA ddd dsjkdsjsこん dk 嘿嘿"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str6))
	fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str6))
	//fmt.Printf("The number of characters in string str2 is %d", len([]rune(str6)))

	//strings 包
	str7:="This is an example of a string"
	fmt.Println("HasPrefix() test:",strings.HasPrefix(str7,"This"))
	fmt.Println("HasSuffix() test:",strings.HasSuffix(str7,"string"))
	fmt.Println("Containers() test:",strings.Contains(str7,"is"))
	//子串查找
	fmt.Println("Index() test:",strings.Index(str7,"an"))
	fmt.Println("LastIndex() test:",strings.LastIndex(str7,"an"))
	fmt.Println("IndexRune() test:",strings.IndexRune("一个字符串",rune('字')))
	fmt.Println("Replace() test:",strings.Replace("hi yoyoyo","yo","gi",2))
	fmt.Println("Count() test:",strings.Count("I love you","o"))
	//重复生成
	fmt.Println("Repeat() test:",strings.Repeat("yo",3))
	//大小写转换
	fmt.Println("ToLower() test:",strings.ToLower("Hello"))
	fmt.Println("ToUpper() test:",strings.ToUpper("Hello"))
	//修剪字符串
	fmt.Println("TrimSpace() test:",strings.TrimSpace("  Hello  "))
	fmt.Println("Trim() test:",strings.Trim("cut Hello cut","cut"))
	fmt.Println("TrimLeft() test:",strings.TrimLeft("cut Hello cut","cut"))
	fmt.Println("TrimRight() test:",strings.TrimRight("cut Hello cut","cut"))
	//分割和拼接字符串
	sl1:=strings.Fields("This is a string")
	fmt.Println("Fields() test:",sl1)
	sl2:=strings.Split("THis &is &a& string","&")
	fmt.Println("Split() test:",sl2)
	fmt.Println("join() test:",strings.Join(sl1,"￥"))
	//strconv类型转换
	var i1 int=666
	var s string ="999"
	fmt.Println(strconv.Itoa(i1))
	fmt.Println(strconv.Atoi(s))
	var pi=3.1415
	fmt.Println(strconv.FormatFloat(pi,'f',3,64))
	s2:="3.14159"
	fmt.Println(strconv.ParseFloat(s2,64))
}
