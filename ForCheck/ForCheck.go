package ForCheck

import (
	"fmt"
)

func SpreadOfTwoArrays(first []int, second []int) []int {
	var result []int

	for _, v := range first {
		result = append(result, v)
	}
	for _, v := range second {
		result = append(result, v)
	}

	return result
}

//var arrayForSort = []int{1, 2, 4, 6, 8, 9, 5, 3}
//var arrayForSort = []int{1, 2, 3, 4, 5, 6, 7}

func BubbleSort(arrayForSort []int) []int {
	var initialArray = arrayForSort

	for i := 0; i < len(initialArray)-1; i += 1 {
		isAlreadySorted := false

		for j := 0; j < len(initialArray)-1; j += 1 {
			if initialArray[j] > initialArray[j+1] {
				var valWithI = initialArray[j]
				var valNext = initialArray[j+1]
				initialArray[j+1] = valWithI
				initialArray[j] = valNext
				isAlreadySorted = true
			}
		}
		//fmt.Println(testArray)
		if !isAlreadySorted {
			break
		}

	}

	return initialArray
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

//	var matrix = [3][3]int{
//		{1, -2, 3},
//		{4, 0, 6},
//		{-7, 8, 9},
//	}

func CalcCheatMatrixDeterminant(matrix [3][3]int) int {

	var resultFirst = matrix[0][0]*matrix[1][1]*matrix[2][2] + matrix[2][0]*matrix[0][1]*matrix[1][2] + matrix[1][0]*matrix[2][1]*matrix[0][2]
	var resultSecond = -matrix[2][2]*matrix[1][1]*matrix[0][0] - matrix[0][0]*matrix[2][1]*matrix[1][2] - matrix[1][0]*matrix[0][1]*matrix[2][2]

	return resultFirst + resultSecond
}

// 5
func CalcSumOfOddAndEvenNumberInArray() (int, int) {
	var array = make([]int, 10)

	for i := 0; i < 10; i++ {

		fmt.Printf("Введите число для индекса %d \n", i)
		fmt.Scan(&array[i])
	}
	var sumOfOdd = 0
	var sumOfEven = 0

	for _, v := range array {
		if v%2 == 0 {
			sumOfOdd += 1
		} else {
			sumOfEven += 1
		}
	}

	return sumOfOdd, sumOfEven
}

// 6
func ReverseArray(arr []int) []int {
	return processReverseArray(arr)
}

func processReverseArray(arr []int) []int {
	var result = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		result[i] = arr[len(arr)-i-1]
	}
	return result
}
