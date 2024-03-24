package ForCheck

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//var matrix = [3][3]int{
//	{1, 2, 3},
//	{4, 5, 6},
//	{7, 8, 9},
//}

// TODO переделать
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

func FindIndexMethod(arr []int, neededNumber int) int {
	for index, val := range arr {
		if val == neededNumber {
			return index
		}
	}

	return -1
}

// Задание 3. Чётные и нечётные

func CalcOddAndEven(arr []int) ([]int, []int) {
	var oddArr, evenArr []int

	for _, val := range arr {
		if val%2 == 0 {
			oddArr = append(oddArr, val)
		} else {
			evenArr = append(evenArr, val)
		}
	}

	return oddArr, evenArr
}

// Задание 4.1 Поиск символов в нескольких строках p.s неправильно понял задачу

func MistakeFindNumberOfSymbols(sentences []string, chars []rune) map[string]int {
	var sumOfСhars = prepareSymbolsRuneMap(chars)
	var onceArray = spreadStringsElementsInOneArray(sentences)

	for _, val := range onceArray {
		runeSlice := []rune(val)
		for _, runeVal := range runeSlice {
			_, exists := sumOfСhars[runeVal]
			if exists {
				sumOfСhars[runeVal] += 1
			}
		}
	}

	return mistakeFindNumberOfSymbolsAdapter(sumOfСhars)
}

func prepareSymbolsRuneMap(chars []rune) map[rune]int {
	var resultMap = map[rune]int{}

	for _, val := range chars {
		resultMap[val] = -1
	}

	return resultMap
}

func spreadStringsElementsInOneArray(arr []string) []string {
	var result []string

	for _, val := range arr {
		stringsElements := strings.Split(val, " ")
		result = append(result, stringsElements...)
	}

	return result
}

func mistakeFindNumberOfSymbolsAdapter(runeMap map[rune]int) map[string]int {
	var result = map[string]int{}
	for key, val := range runeMap {
		if val == -1 {
			result[string(key)] = 0
		} else {
			result[string(key)] = val
		}
	}
	return result
}

// Задание 4.2 Поиск символов в нескольких строках

func FindNumberOfSymbolsAdapter(runeMap map[rune]int) [][2]string {
	var result [][2]string
	for key, val := range runeMap {

		resultVal := 0
		if val != -1 {
			resultVal = val
		}
		result = append(result, [2]string{string(key), strconv.Itoa(resultVal)})
	}
	return result
}

func FindNumberOfSymbols(sentences []string, chars []rune) [][2]string {
	var result [][2]string

	for _, sentence := range sentences {
		var mapOfRune = prepareSymbolsRuneMap(chars)
		lastWord := strings.Split(sentence, " ")[len(strings.Split(sentence, " "))-1]
		lastWordInRuneType := []rune(lastWord)
		for i := 0; i < len(lastWordInRuneType); i += 1 {
			_, exist := mapOfRune[lastWordInRuneType[i]]
			if exist && mapOfRune[lastWordInRuneType[i]] == -1 {
				mapOfRune[lastWordInRuneType[i]] = len([]rune(sentence)) - len(lastWordInRuneType) + i
			}
		}

		fmt.Println(FindNumberOfSymbolsAdapter(mapOfRune))

	}
	return result
}

//
