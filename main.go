package main

import (
	"fmt"
	doneexercise "main.go/done-exercise"
)

func main() {
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	//sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	//
	//chars := []rune{'H', 'E', 'l', 'П', 'М'}
	//
	//ForCheck.FindNumberOfSymbols(sentences, chars)
	array := []int{1, 5, 11, 4, 54, 2, 3, 5, 6, 8}
	anon := func(array *[]int) {

		var sortedArray = doneexercise.BubbleSort(*array)
		var result []int
		for i := len(sortedArray) - 1; i >= 0; i-- {
			result = append(result, sortedArray[i])
		}

		*array = result
	}

	anon(&array)
	fmt.Println(array)

}
