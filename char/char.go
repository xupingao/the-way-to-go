package main

import (
	"fmt"
	"unicode"
)
func main() {
	var ch0 byte=65
	var ch1 byte='\x41'
	fmt.Printf("%d-%d\n",ch0,ch1)
	fmt.Printf("%c-%c\n",ch0,ch1)

	var ch int='\u0041'
	var ch2 int='\u03B2'
	var ch3 int='\U00101234'
	fmt.Printf("%d-%d-%d\n",ch,ch2,ch3)
	fmt.Printf("%c-%c-%c\n",ch,ch2,ch3)
	fmt.Printf("%X-%X-%X\n",ch,ch2,ch3)
	fmt.Printf("%U-%U-%U\n",ch,ch2,ch3)
	fmt.Println(unicode.IsLetter(rune(ch)))
}
