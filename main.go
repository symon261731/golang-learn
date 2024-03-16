package main

import "main.go/ForCheck"

func main() {
	ForCheck.CallbackFunction(func(x int, y int) int {
		return x + y
	})

	ForCheck.CallbackFunction(func(x int, y int) int {
		return x - y
	})

	ForCheck.CallbackFunction(func(x int, y int) int {
		return x * y
	})

}
