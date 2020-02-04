package main

import "os"

func main(){
	//fmt.Println("Starting the program")
	//	//panic("A severe error occurred: stopping the program!")
	//	//fmt.Println("Ending the program")
	var user=os.Getenv("USER")
	if user==""{
		panic("Unkown user:novalue for $USER")
	}
}
