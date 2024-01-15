package inprogressexercises

import (
	"fmt"
)

func CombinationOfBrackets(numberOfPares int) {
	result := []string{}

	if numberOfPares == 0 {
		fmt.Print(result)
		return
	}

}

func MergeNumberLists(numberList ...[]int) []int {
	var result []int

	for _, value := range numberList {
		result = append(result, value...)
	}

	return result
}
