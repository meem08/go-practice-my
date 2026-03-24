package main

import (
	"fmt"
	"strings"
)

func Ascii(text string, data []string) string {
	textin := strings.Split(text, "\\n")
	asciimap := make(map[rune][]string) // created an empty map tostore ascii characters
	var result strings.Builder
	char := 32 // the printable ascii character starts from 32
	start := 0

	// skip leading empty lines
	for start < len(data) && strings.TrimSpace(data[start]) == "" {
		start++
	}
	for i := 0; i+8 < len(data); i += 9 {
		asciimap[rune(char)] = data[i : i+9]
		char++
	}
	for j, val := range textin {
		if val == "" {
			if j != 0 {
				result.WriteString("\n")
			}
			continue
		}
		for row := 0; row < 8; row++ {
			for _, char := range val {
				if val, ok := asciimap[char]; ok {
					result.WriteString(val[row])
				}
			}
			result.WriteString("\n")
		}
	}
	fmt.Println("Map size:", len(asciimap))
	fmt.Println(asciimap['H'])
	return result.String()
}
