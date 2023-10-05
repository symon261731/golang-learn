package main

import "fmt"

func main() {
	var inNum int

	fmt.Scan(&inNum)

	fmt.Println("It is", inNum/30, "hours", 2*(inNum%30), "minutes")
}
