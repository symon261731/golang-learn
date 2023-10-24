package main

import "fmt"

func main() {
	Example()
}

func Example() {
	cipherCaesar := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(cipherCaesar); i += 1 {
		letter := cipherCaesar[i]
		if letter > 'a' && letter < 'z' {
			letter = letter - 3
			if letter < 'a' {
				letter += 26
			}
		} else if letter > 'A' && letter <= 'Z' {
			letter = letter - 3
			if letter < 'A' {
				letter = letter + 26
			}
		}
		fmt.Printf("%c", letter)
	}
}
