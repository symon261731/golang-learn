package main

import (
	"fmt"
	"strings"
)

func main() {
	// var arr = []int{5, 5, 10, 10}
	// CalcMoney(arr)
	NumberOfWordsInUpperCase()
}

func NumberOfWordsInUpperCase() {
	exampleString := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	accumulator := 0

	for _, some := range strings.Split(exampleString, " ") {
		var test = fmt.Sprintf("%c", some[0])

		if strings.ToLower(test) != test {
			accumulator += 1
		}
	}

	fmt.Println(accumulator)
}

// func NumbersInTheString() {
// 	example := "a10 10 20b 20 30c30 30 dd"

// 	for _, some := range strings.Split(example, " ") {
// 		if
// 	}
// }

func CalcMoney(input []int) {
	sum := 0

	for _, value := range input {

		switch value {
		case 5:
			sum += value
		case 10:
			sum += 10
			sum -= 5
			if sum < 0 {
				fmt.Print(false)
				return
			}
			sum += 5
		case 20:
			sum += 20
			sum -= 15
			if sum < 0 {
				fmt.Print(false)
				return
			}
			sum += 5
		}

	}

	fmt.Print(true)
}

func IdentifyTimeOfSeazon() {
	var month string

	months := map[string]int{
		"январь":   1,
		"февраль":  2,
		"март":     3,
		"апрель":   4,
		"май":      5,
		"июнь":     6,
		"июль":     7,
		"август":   8,
		"сентябрь": 9,
		"октябрь":  10,
		"ноябрь":   11,
		"декабрь":  12,
	}

	fmt.Print("Введите месяц: ")
	fmt.Scan(&month)

	resultOfInput := months[strings.ToLower(month)]

	if resultOfInput == 0 {
		fmt.Println("Неверный формат данных")
		return
	}

	switch resultOfInput {

	case 1:
		fmt.Println("Зима")
	case 2:
		fmt.Println("Зима")
	case 3:
		fmt.Println("Весна")
	case 4:
		fmt.Println("Весна")
	case 5:
		fmt.Println("Весна")
	case 6:
		fmt.Println("Лето")
	case 7:
		fmt.Println("Лето")
	case 8:
		fmt.Println("Лето")
	case 9:
		fmt.Println("Осень")
	case 10:
		fmt.Println("Осень")
	case 11:
		fmt.Println("Осень")
	case 12:
		fmt.Println("Зима")
	}

}

func ShowWeek() {

	var input string

	weekDays := map[string]string{
		"пн": "понедельник вторник среда четверг пятница",
		"вт": "вторник среда четверг пятница",
		"ср": "среда четверг пятница",
		"чт": "четверг пятница",
		"пт": "пятница",
	}

	fmt.Print("Введите пожалуйста день недели в сокращенной форме: ")
	fmt.Scan(&input)

	var resultOfInput = strings.ToLower(input)

	switch resultOfInput {
	case "пн":
		fmt.Println(weekDays["пн"])
	case "вт":
		fmt.Println(weekDays["вт"])
	case "ср":
		fmt.Println(weekDays["ср"])
	case "чт":
		fmt.Println(weekDays["чт"])
	case "пт":
		fmt.Println(weekDays["пт"])
	default:
		fmt.Println("Неверный формат данных")
		return
	}
}

func CalcGroupsOfStudents() {
	var n, k, studentIndex int

	fmt.Println("Введите пожалуйста сколько студентов, групп, порядковый номер студента: ")
	fmt.Scan(&n, &k, &studentIndex)

	groupNum := studentIndex % k
	studentsInOneGroup := n / k

	for i := 1; i <= k; i++ {
		if groupNum == 0 {
			groupNum = k
		}
		if studentIndex <= studentsInOneGroup*i {
			fmt.Printf("студент находиться в %d группе", i)
		}
	}

	return
}
