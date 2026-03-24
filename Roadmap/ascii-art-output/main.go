package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	printerror := "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard"
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println(printerror)
		return
	}
	var text string
	var banner string = "standard"
	var store string

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("Not a banner file\nCan't read")
	}

	switch len(os.Args) {
	case 2:
		text = os.Args[1]
	case 3:
		text = os.Args[1]
		banner = os.Args[2]
	case 4:
		if strings.HasPrefix(os.Args[1], "--output=") {

			storeoutput := strings.Split(os.Args[1], "=")
			if len(storeoutput) == 2 && storeoutput[1] != "" {
				store = storeoutput[1]
			} else {
				fmt.Println("Expected:  --output=\nElse nohing else!!")
			}
		}
		text = os.Args[1]
		banner = os.Args[3]
	default:
		fmt.Println(printerror)

	}
	file, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Not a banner file", err)
	}
	data := strings.Split(string(file), "\n")

	result := Ascii(text, data)
	if store != "" {
		err := os.WriteFile(store, []byte(result), 0644)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	fmt.Println("Text:", text)
	fmt.Print(result)
}
