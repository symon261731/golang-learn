package main

import "fmt"

func main() {
	var (
		n   int
		c   int
		d   int
		res int
	)

	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i := 1; i <= n; i += 1 {
		if i%c == 0 && i%d != 0 {
			res = i
			break
		}

	}

	if res != 0 {
		fmt.Println(res)
	}

}
