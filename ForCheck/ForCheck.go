package ForCheck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//var matrix = [3][3]int{
//	{1, 2, 3},
//	{4, 5, 6},
//	{7, 8, 9},
//}

func CalcCheatMatrixDeterminant(matrix [3][3]int) int {

	mainDiagonal := matrix[0][0] * matrix[1][1] * matrix[2][2]
	alternateDiagonal := matrix[0][2] * matrix[1][1] * matrix[2][0]

	leftSide := mainDiagonal + (matrix[0][1] * matrix[1][2] * matrix[2][0]) + (matrix[1][0] * matrix[2][1] * matrix[0][2])
	rightSide := alternateDiagonal + (matrix[2][1] * matrix[1][2] * matrix[0][0]) + (matrix[1][0] * matrix[0][1] * matrix[2][2])

	return leftSide - rightSide
}

func SpreadOfTwoArrays(first []int, second []int) []int {
	var result = append(first, second...)
	sort.Ints(result)

	return result
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

func CallbackFunction(cb func(someX int, someY int) int) {

	defer cb(10, 15)
}
