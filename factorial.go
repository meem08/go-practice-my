package main

/* Factorial Calculator: Write a program that calculates the factorial of a number (e.g., 5! = 5 * 4 * 3 * 2 * 1 = 120) using a for loop. The number should be a fixed value in your code (e.g., num := 7).*/
import "fmt"
func Factorial(n int) int {
	result := 1
	for i:=1; i <=n; i++ {
		result *= i
		
	}
	return result
}
func main() {
	fmt.Println(Factorial(6))
}