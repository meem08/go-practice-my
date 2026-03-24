package main

import (
	"html/template"
	"fmt"
	"net/http"
)


type Data struct{
    Name string
}
 
var templ = template.Must(template.New("index.html").ParseFiles("./template/index.html"))

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/generate", GenerateHandler)

	fmt.Println("Sever listening on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		
	}
}


func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(200)
	templ.Execute(w, nil)
}


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte("Hello blessing"))

}
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
     
    if r.Method != http.MethodPost{
         http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	r.ParseForm()
      name := r.FormValue("myName")

	 templ.Execute(w,  Data{
		Name:name,
	 })

} 