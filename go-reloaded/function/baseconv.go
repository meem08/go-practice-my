package function

import (
	"strconv"
	"strings"
)

func ParseBin(file string) string {
	splitfile := strings.Fields(file)
	for i, value := range splitfile {
		if value == "(bin)" {
			oldbin := splitfile[i-1]
			newbin, _ := strconv.ParseInt(oldbin, 2, 64)
			splitfile[i-1] = strconv.Itoa(int(newbin))
			splitfile = append(splitfile[:i], splitfile[i+1:]...)
		}
	}
	newfile := strings.Join(splitfile, " ")
	return newfile
}

func ParseHex(file string) string {
	splitfile := strings.Fields(file)
	for i, value := range splitfile {
		if value == "(hex)" {
			oldhex := splitfile[i-1]
			nehex, _ := strconv.ParseInt(oldhex, 16, 64)
			splitfile[i-1] = strconv.Itoa(int(nehex))
			splitfile = append(splitfile[:i], splitfile[i+1:]...)
		}
	}
	newfile := strings.Join(splitfile, " ")
	return newfile
}
