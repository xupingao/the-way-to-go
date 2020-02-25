package main

import (
	"fmt"
	"os"
	"text/template"
)
type holder struct{
	A ,B int
	V []int
	O Obj
}
type Obj struct{
	A,B int
}
func main(){
	defer func()(){if err:=recover();err!=nil{
	fmt.Println(err)
	}} ()

	t:=template.Must(template.New("tmpl").Parse("{{if .A}}HELLOWORLD{{end}}{{with .O}}{{.A}}{{end}}"))
	t.Execute(os.Stdout,holder{1,2,[]int{1,2},Obj{3,4}})
	t2:=template.Must(template.New("tmpl").Parse("{{with $val:=`hello`}}{{$val}}{{.}}{{end}}"))
	t2.Execute(os.Stdout,nil)
	t3:=template.Must(template.New("tmpl").Parse("{{range .V}}{{.}}{{end}}"))
	t3.Execute(os.Stdout,holder{1,2,[]int{1,2},Obj{3,4}})

	t4:=template.Must(template.New("tmpl").Parse("{{with $str:=`hello`}}" +
		"{{printf `%s,%s!` $str `mary`}}{{end}}"))
	t4.Execute(os.Stdout,nil)
}
