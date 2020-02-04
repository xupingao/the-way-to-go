package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Myinfo struct{
	Name string
	Address *Add
	Age int
}
type Add struct{
	Contry string
	City string
}
func main(){
	//man:= Myinfo{"kobe",&Add{"USA","LAL"},41}
	//myjs,_:=json.Marshal(man)
	//fmt.Printf("JSON format: %s\n",myjs)
	//outPut,_:=os.OpenFile("myjson.json", os.O_CREATE|os.O_WRONLY, 0666)
	//defer outPut.Close()
	//enc:=json.NewEncoder(outPut)
	//enc.Encode(man)

	input,err:=os.Open("myjson.json")
	defer  input.Close()
	if err!=nil{
		err.Error()
	}
	dec:=json.NewDecoder(input)
	var aman Myinfo
	err=dec.Decode(&aman)
	if err!=nil{
		err.Error()
	}
	fmt.Println(aman)
}

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}