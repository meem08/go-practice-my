package main

import (
	"fmt"
	"go-reloaded/function"
)
func main(){
	// check := function.ParseBin("10 (bin) files")
	// fmt.Println(check)
	// check2:= function.ParseHex("1E (hex) items")
	// fmt.Println(check2)
	// check3 := function.Up("Ready, set, go (up) !")
	// fmt.Println(check3)
	 check4 := function.Vowel("There is no greater agony than bearing A untold story inside you")
	 fmt.Println(check4)
	
	check5:= function.Numcap("it was the age of foolishness (cap, 4)")
	fmt.Println(check5)
}