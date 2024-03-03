package main

import (
	"fmt"
	"main.go/ForCheck"
)

func main() {
	var firstMatrix = [3][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 5},
		{13, 14, 15, 16, 17},
	}
	var secondMatrix = [5][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{6, 7, 8, 9},
		{3, 4, 5, 6},
		{2, 3, 4, 5},
	}
	fmt.Println(ForCheck.MultiplyMatrix(firstMatrix, secondMatrix))
}
