package main

import (
	"flag"
	"fmt"
)

func main() {
	fullString := flag.String("str", "", "fullString")
	subStr := flag.String("substr", "", "sub string")

	flag.Parse()

	runeFullString := []rune(*fullString)
	runeSubStr := []rune(*subStr)

	fmt.Println(runeSubStr)

	numberOfСoincidences := 0

	for i := 0; i < len(runeFullString); i += 1 {
		for j := 0; j < len(runeSubStr); j += 1 {
			if i+j >= len(runeFullString) {
				break
			}
			if runeSubStr[j] == runeFullString[j+i] {
				numberOfСoincidences += 1
			} else {
				numberOfСoincidences = 0
			}
		}
		if numberOfСoincidences == len(runeSubStr) {
			fmt.Println(true)
			break
		}

	}

	if numberOfСoincidences != len(runeSubStr) {
		fmt.Println(false)
	}

}
