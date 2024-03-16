package ForCheck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func MultiplyMatrix(firstMatrix [3][5]int, secondMatrix [5][4]int) [3][4]int {

	var result = [3][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	for i := 0; i < len(firstMatrix); i += 1 {
		for j := 0; j < len(secondMatrix)-1; j += 1 {
			for k := 0; k < len(secondMatrix); k += 1 {
				result[i][j] += firstMatrix[i][k] * secondMatrix[k][j]
			}
		}
	}

	return result
}

func SpreadOfTwoArrays(first []int, second []int) []int {
	var result []int = append(first, second...)
	sort.Ints(result)

	return result
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

//var matrix = [3][3]int{
//	{1, -2, 3},
//	{4, 0, 6},
//	{-7, 8, 9},
//}

func CalcCheatMatrixDeterminant(matrix [3][3]int) int {

	var resultFirst = (matrix[0][0] * matrix[1][1] * matrix[2][2]) + (matrix[2][0] * matrix[0][1] * matrix[1][2]) + (matrix[1][0] * matrix[2][1] * matrix[0][2])
	var resultSecond = -(matrix[2][2] * matrix[1][1] * matrix[0][0]) - (matrix[0][0] * matrix[2][1] * matrix[1][2]) - (matrix[1][0] * matrix[0][1] * matrix[2][2])

	return resultFirst + resultSecond
}

func CalcNumberInArray() ([]int, int) {
	rand.Seed(time.Now().UnixNano())
	var arr []int

	var findNum int
	var sum = 0
	fmt.Print("Введите число количество которых хотите найти в массиве ")
	fmt.Scan(&findNum)

	for i := 0; i < rand.Intn(50); i += 1 {
		arr = append(arr, rand.Intn(50))
	}

	for _, v := range arr {
		if v == findNum {
			sum++
		}

	}

	return arr, sum

}

func CalcFormula() any {
	var x int16
	var y uint8
	var z float32

	fmt.Println("X Y Z")
	fmt.Scan(&x, &y, &z)

	var result float32 = 2*float32(x) + float32(y*y) - 3/z

	return result
}
