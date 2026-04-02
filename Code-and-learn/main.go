package main

import(
	"fmt"
)

func main() {
	var mystring string
	mystring = "Hello\nWorld"
	fmt.Println(mystring)
	mystring = `hello

world`
	fmt.Println(mystring)
	var firstname, lastname string
	firstname = "Blessing"
	lastname = "Akinsuyi"
	// 
	fullname := fmt.Sprintf("%s %s", firstname, lastname)
	fmt.Println(fullname)
	var a []int= []int{1,2, 3,4,5}
	var b []int= []int{10,20,30,40}
	a=append(a, b...)
}