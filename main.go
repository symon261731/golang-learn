package main

import "fmt"

func main() {
	MultiplyMatrix()
}

// нужно решать через 3 цикла
func MultiplyMatrix() {
	var firstMatrix = [3][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 5},
		{1, 2, 3, 4, 5},
	}
	var secondMatrix = [5][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{6, 7, 8, 9},
		{3, 4, 5, 6},
		{2, 3, 4, 5},
	}
	var result = [3][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	for i := 0; i < len(firstMatrix); i += 1 {
		for j := 0; j < len(secondMatrix); j += 1 {
			result[i][i] += firstMatrix[i][j] * secondMatrix[j][i]
		}
	}

	fmt.Println(result)
}
