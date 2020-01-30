package utils

import (
	"fmt"
	"math"
)

var Pi float64

func init(){
	Pi=4*math.Atan(1)
}
func Fun()  {
	return
}
func Add(a,b int) (c,d int) {
	return a+b,a-b
}

func Uint8FromInt(n int)(uint8,error){
	if 0<=n &&n<=math.MaxUint8{
		return uint8(n),nil
	}
	return 0,fmt.Errorf("%d is out of the uint8 range",n)
}
 func IntFromFloat64(x float64)int {
 	if math.MaxInt32<=x&&x<=math.MaxInt32{
 		whole,fraction:=math.Modf(x)
 		if fraction>=0.5{
 			whole++
		}
 		return int(whole)
	}
 	panic(fmt.Stringf("%g is out of the int32 range",x))
 }