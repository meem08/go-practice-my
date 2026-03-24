package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
}

var temp = template.Must(template.New("index.html").ParseFiles("./index.html"))

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/web", WebHandler)
	fmt.Println("Server is Running at Port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Unable to connect to server port :8080", err)
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contnt-type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte("<h1>Hello! Welcome to My Website</h1>"))
}
func WebHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(200)
	use := User{
		Name: " Blessing",
	}
	temp.Execute(w, use)

}
