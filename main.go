package main

import "fmt"

func main() {

	var i int

	fmt.Scan(&i)

	var sum int = 0

	for k := 1; k <= i; k++ {
		var scanVal int
		fmt.Scan(&scanVal)
		if scanVal >= 10 && scanVal < 100 && scanVal%8 == 0 {
			sum += scanVal
		}
	}

	fmt.Println(sum)
}
