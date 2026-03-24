package function

import (
	"strconv"
	"strings"
)

// Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO!")

func Up(file string) string {
	splitfile := strings.Fields(file)
	for i, val := range splitfile {
		if val == "(up)" {
			unp := splitfile[i-1]
			uup := strings.ToUpper(unp)
			splitfile[i-1] = uup
			splitfile = append(splitfile[:i], splitfile[i+1:]...)
		} else if val == "(up,"{
			idx := strings.TrimRight(splitfile[i+1], ")")
			num, _:= strconv.Atoi(idx)
			for j:= 1; j <= num; j++{
				splitfile[i-j] = strings.ToUpper(splitfile[i-j])
			}
			splitfile = append(splitfile[:i], splitfile[i+2:]... )
		}
	}
	newfile := strings.Join(splitfile, " ")
	return newfile
}

// Every instance of (low) converts the word before with the Lowercase version of it.
//
//	(Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")
func LowCase(file string) string {
	splitfile := strings.Fields(file)
	for i, val := range splitfile {
		if val == "(low)" {
			oldlow := splitfile[i-1]
			newlow := strings.ToLower(oldlow)
			splitfile[i-1] = newlow
			splitfile = append(splitfile[:i], splitfile[i+1:]...)
		} else if val == "(low," {
			idx := strings.TrimRight(splitfile[i+1], ")")
			num, _ := strconv.Atoi(idx)
			for j := 1; j <= num; j++ {
				splitfile[i-j] = strings.ToLower(splitfile[i-j])
			}
			splitfile = append(splitfile[:i], splitfile[i+2:]...)
		}
	}
	newfile := strings.Join(splitfile, " ")
	return newfile
}

// Every instance of (cap) converts the word before with the capitalized version of it.
//
//	(Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
func Capcut(file string) string {
	splitfle := strings.Fields(file)
	for i, val := range splitfle {
		if val == "(cap)" {
			oldcap := splitfle[i-1]
			newcap := title(oldcap)
			splitfle[i-1] = newcap
			splitfle = append(splitfle[:i], splitfle[i+1:]...)
		}
	}
	newfile := strings.Join(splitfle, " ")
	return newfile
}
func title(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
func Numcap(file string) string {
	splitfile := strings.Fields(file)
	for i, val := range splitfile {
		if val == "(cap," {
			idx := strings.TrimRight(splitfile[i+1], ")")
			num, _ := strconv.Atoi(idx)
			for j := 1; j <= num; j++ {
				splitfile[i-j] = title(splitfile[i-j])
			}
			splitfile = append(splitfile[:i], splitfile[i+2:]...)
		}
	}
	newfile := strings.Join(splitfile, " ")
	return newfile
}

