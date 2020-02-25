package main
import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)
var lenPath=len("/view/")
var titleValidator=regexp.MustCompile("^[a-zA-Z0-9]+$")
var templates=make(map[string]*template.Template)
var err error

type Page struct{
	Title string
	Body []byte
}

func init(){
	for _,tmpl:=range([]string{"edit","view"}){
		templates[tmpl]=template.Must(template.ParseFiles(tmpl+".html"))
	}
}
func makeHandler(fn func(w http.ResponseWriter,request *http.Request,title string))(http.HandlerFunc){
	return func(w http.ResponseWriter,request *http.Request){
		title:=request.URL.Path[lenPath:]
		if !titleValidator.Match([]byte(title)){
			http.NotFound(w,request)
			return
		}else{
			fn(w,request,title)
		}
	}
}
func main(){
	http.HandleFunc("/view/",makeHandler(viewHandler))
	http.HandleFunc("/edit/",makeHandler(editHandler))
	http.HandleFunc("/save/",makeHandler(saveHandler))
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
func viewHandler(w http.ResponseWriter,request *http.Request,title string){
	page,err:=Load(title)
	if err!=nil{
		http.Redirect(w,request,"/edit/"+title,http.StatusFound)
	}
	renderTemplate(w,"view",page)
}
func editHandler(w http.ResponseWriter,request *http.Request,title string){
	page,err:=Load(title)
	if err!=nil{
		page=&Page{title,nil}
	}
	renderTemplate(w,"edit",page)
}
func saveHandler(w http.ResponseWriter,request *http.Request,title string){
	body:=request.FormValue("body")
	p:=&Page{title,[]byte(body)}
	err:=p.save()
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	http.Redirect(w,request,"/view/"+title,http.StatusFound)

}
func renderTemplate(w http.ResponseWriter,tmpl string,p *Page){
	err:=templates[tmpl].Execute(w,p)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}

}
func (p *Page)save()error{
	filename:=p.Title+".txt"
	return ioutil.WriteFile(filename,p.Body,0600)
}
func Load(path string)(*Page,error){
	body,err:=ioutil.ReadFile(path+".txt")
	if err!=nil{
		return nil,err
	}
	return &Page{path,body},nil
}