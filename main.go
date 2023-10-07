package main

import "fmt"

func main() {
	var (
		x int
		p int
		y int
	)

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	var thisYearMOney = x

	for year := 1; ; year += 1 {
		if thisYearMOney+thisYearMOney*p/100 >= y {

			fmt.Println(year)
			break
		} else {
			thisYearMOney = thisYearMOney + thisYearMOney*p/100
		}
	}
}
