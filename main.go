package main

import "fmt"

func main() {

	var matrix = [3][3]int{
		{5, 6, 7},
		{-8, 1, 2},
		{3, 4, 5},
	}
	fmt.Println(CalcCheatMatrixDeterminant(matrix))
}

// [
//
//	[a1,b1,c1,]
//	[a2,b2,c2,]
//	[a3,b3,c3]
//
// ]
// a1b2c3 + a3b1c2 + a2b3c1
// -
// a3b2c1 - a1b3c2 - a2b1c3

func CalcCheatMatrixDeterminant(matrix [3][3]int) int {

	mainDiagonal := matrix[0][0] * matrix[1][1] * matrix[2][2]
	alternateDiagonal := matrix[0][2] * matrix[1][1] * matrix[2][0]

	leftSide := mainDiagonal + (matrix[0][1] * matrix[1][2] * matrix[2][0]) + (matrix[1][0] * matrix[2][1] * matrix[0][2])
	rightSide := alternateDiagonal + (matrix[2][1] * matrix[1][2] * matrix[0][0]) + (matrix[1][0] * matrix[0][1] * matrix[2][2])

	return leftSide - rightSide
}
