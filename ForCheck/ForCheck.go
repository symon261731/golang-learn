package ForCheck

import (
	"errors"
	"fmt"
)

// 1
func IsOddNumber(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// 2
func CalcCoordinates() (int, int) {
	var x, y int
	fmt.Println("Введите пожалуйста значение X")
	fmt.Scan(&x)
	fmt.Println("Введите пожалуйста значение Y")
	fmt.Scan(&y)

	return prepareCoordinates(x, y)
}

// Вроде будет недоступна вне пакета
func prepareCoordinates(x int, y int) (int, int) {
	var resultX = 2*x + 10
	var resultY = -3*y - 5

	return resultX, resultY
}

// 3
func CalcAddAndMultiplyResult(num int) (int, int) {
	return multiplyNumber(num), addNumber(num)
}

func multiplyNumber(num int) int {
	var multiplyNum int
	fmt.Println("На сколько умножить")
	fmt.Scan(&multiplyNum)
	return multiplyNum * num
}

func addNumber(number int) int {
	var addNumber int
	fmt.Println("Сколько прибавить")
	fmt.Scan(&addNumber)
	return addNumber + number
}

// 4
var globalVariable = 5

func CalcResultOfSumFunctions() int {
	var callFirstFunction = addFirstFunction(globalVariable)
	var callSecondFunction = addSecondFunction(callFirstFunction)
	var callThirdFunction = addThirdFunction(callSecondFunction)

	return callThirdFunction
}

func addFirstFunction(num int) int {
	return num + 5
}

func addSecondFunction(num int) int {
	return num + 6
}
func addThirdFunction(num int) int {
	return num + 8
}

//5

func CalcSumOfOddAndEvenNumberInArray() ([]int, error) {
	var array = make([]int, 10)

	for i := 0; i < 10; i++ {
		var num int

		fmt.Printf("Введите число для индекса %d \n", i)
		num, err := fmt.Scan(&num)
		if err != nil {
			return nil, errors.New("error")
		}
		array[i] = num
	}

	return array, nil
}

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
