package main

import "fmt"

func main() {

	var n int
	var maxVal int = 0
	var acc int = 0

	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > maxVal {
			maxVal = n
			acc = 1
		} else if n == maxVal {
			acc += 1
		}
	}

	fmt.Println(acc)
}
