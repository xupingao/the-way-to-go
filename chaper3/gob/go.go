package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Type1 struct{
	X,Y,Z int
}
type Type2 struct{
	X,Z int
}
func main(){
	//var network bytes.Buffer
	//enc:=gob.NewEncoder(&network)
	//dec:=gob.NewDecoder(&network)

//	t:=Type1{1,2,0}
	//enc.Encode(t)
	//var t2 Type2
	//dec.Decode(&t2)
	//fmt.Println(t2)

	//gobOut,_:=os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	//defer gobOut.Close()
	//enc:=gob.NewEncoder(gobOut)
	//enc.Encode(t)

	gobOut,_:=os.Open("vcard.gob")
	defer gobOut.Close()
	dec:=gob.NewDecoder(gobOut)
	var t2 Type2
	dec.Decode(&t2)
	fmt.Println(t2)
}
