package main

import (
	"fmt"
	"strings"
	//"strings"
)

// 1. Print Each Character in a String
// Question
// Write a function that prints each character in a string on a new line.
// Function Signature
// func PrintChars(s string)
func PrintChars(s string) {
	for i := range s {
		fmt.Println(string(s[i]))
	}
}

// 2.Count Characters in a String
// Question
// Write a function that returns the number of characters in a string.
// Function Signature
// func CountChars(s string) int
func CountChars(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		// count = count+1
		count++
	}
	return count
}

// 3. Convert String to Uppercase
// Question: Write a function that converts a string to uppercase.
// Function Signature:func ToUpperCase(s string) string
func ToUpperCase(s string) string {
	// return strings.ToUpper(s)
	var ch string
	for _, val := range s {
		if val >= 'a' && val <= 'z' {
			ch += string(val - 32)
		} else {
			ch += string(val)
		}
	}
	return ch
}

// 4. Convert String to Lowercase
// Question: Write a function that converts a string to lowercase.
// Function Signature:func ToLowerCase(s string) string
func ToLowerCase(s string) string {
	// return strings.ToLower(s)
	var ch string
	for _, val := range s {
		if val >= 'A' && val <= 'Z' {
			ch += string(val+32)
		} else {
			ch += string(val)
		}
	}
	return ch
}

// 5. Reverse a String
// Question:Write a function that reverses a string.
// Function Signature:func Reverse(s string) string
func Reverse(s string) string {
	text := ""
	for i := len(s) - 1; i >= 0; i-- {
		text += string(s[i])
	}
	return text
}

// 6. Check if String is Palindrome
// Question:Write a function that checks whether a string is a palindrome.
// Function Signature:func IsPalindrome(s string) bool
func IsPalindrome(s string) bool{
	// racecar
	left := 0
	right := len(s)-1
	for left < right{
		if s[left]!= s[right]{
			return false
		}
		left++
		right--
	}
	return true
}
// 7. Swap Case of Characters
// Question:Write a function that changes lowercase letters to uppercase and uppercase letters to lowercase.
// Function Signature
func SwapCase(s string) string{
	// GoLanG
	store := ""
	for _, val:= range s{
		if val >= 'A' && val <= 'Z'{
			store += strings.ToLower(string(val))
			
		} else if val >= 'a' && val <= 'z'{
			store +=strings.ToUpper(string(val))
		}
	}
	return store
}
// func Checkstr(s string)bool{
// 	runes := []rune(s)
// 	switch runes {
// 	case runes >= 'A' && runes <= 'Z':
// 		return true
// 	}
// }
func main() {
	// PrintChars("Hello")
	// fmt.Println(CountChars("golang"))
	text:= ToUpperCase("helloDD")
	fmt.Println(text)
	// see:= Reverse("hello")
	// fmt.Println(see)
	fmt.Println(ToLowerCase("HeLLo TheRE"))
	tr:=IsPalindrome("raceca")
	fmt.Println(tr)
	ss := SwapCase("FatHER")
	fmt.Println(ss)
}
