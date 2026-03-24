package function

import "strings"

// Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one.
// (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
// func Punc(file string) string{
// 	check := []string{".", ",", "!", "?", ":", ";"}
// 	for i , val :=
// }

// Every instance of a should be turned into an if the next word begins with a vowel
//  (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").
func Vowel(file string)string{
	split := strings.Fields(file)
	for i, val :=range split{
		if isvowel(val){
			if split[i-1] == "a"{
				split[i-1] = "an"
			}else if split[i-1] == "A"{
				split[i-1]= "An"
			}
		}
	}
	newfile := strings.Join(split, " ")
	return  newfile
}
func isvowel(s string) bool{
	switch s[:1]{
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}