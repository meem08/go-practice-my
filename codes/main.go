package main

import "fmt"

func main() {
	age := 0

	if age == 0 && age <= 12 {
		fmt.Println("User is a child")
	} else if age >= 13 && age <= 19 {
		fmt.Println("User is a teanage")
	} else {
		fmt.Println("User is an Adult")
	}
}
