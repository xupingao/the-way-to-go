package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)
type statistics struct{
	numbers []float64
	mean float64
	median float64
}
const form=`
<html>
 <body>
  <form action="/" method="POST">
    <label for="numbers">Numers(comma or space-separated):</label><br>
    <input typr="text" name="numbers" size="30"><br />
    <input type="submit" value="calculate">
  </form>
 </body>
</html>`
const error=`<p class="error"">%s</p>`
var pageTop=""
var pageBottom=""
func main(){
	http.HandleFunc("/",homePage)
	if err:=http.ListenAndServe(":9001",nil);err!=nil{
		fmt.Println(err.Error())
	}
}

func homePage(w http.ResponseWriter,request *http.Request){
	w.Header().Set("Content-Type","text/html")
	err:=request.ParseForm()
	fmt.Fprint(w,pageTop,form)
	if err!=nil{
		fmt.Println(err.Error)
	}else{
		if numbers,message,ok:=processRequest(request);ok{
			state:=getStats(numbers)
			fmt.Fprint(w, formatStats(state))
		}else if message!=""{
			fmt.Fprintf(w,error,message)
		}
	}
	fmt.Fprint(w,pageBottom)
}

func processRequest(request *http.Request)([]float64,string,bool){
	var numbers []float64
	var text string
	if slice,found:=request.Form["numbers"];found&&len(slice)>0 {
		if strings.Contains(slice[0],"&#65292"){
			text=strings.Replace(slice[0],"&#65292"," ",-1)
		}else{
			text=strings.Replace(slice[0],","," ",-1)
		}
		for _,field:=range(strings.Fields(text)){
			if val,err:=strconv.ParseFloat(field,64);err!=nil{
				return numbers,"invalid",false
			}else{
				numbers=append(numbers,val)
			}
		}
	}
	if len(numbers)==0{
		return numbers,"",false
	}
	return numbers,"",true

}
func getStats(numbers []float64)(stats statistics){
	stats.numbers=numbers
	sort.Float64s(stats.numbers)
	sum:=0.0
	for _,x:=range(numbers){
		sum+=x
	}
	stats.mean=sum/float64(len(numbers))
	middle := len(numbers) / 2
	stats.median= numbers[middle]
	if len(numbers)%2 == 0 {
		stats.median = (stats.median + numbers[middle-1]) / 2
	}
	return
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}