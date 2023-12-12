package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// array := []int{5, 5, 5, 10}
	// array := []int{10, 10}
	// array := []int{20, 5, 5, 5}
	array := []int{5, 5, 10, 10, 20}
	CalcMoney(array)
}

func CalcMoney(input []int) {
	sum := 0

	for _, value := range input {
		if value == 5 {
			sum += value
		}
		if value == 10 {
			sum -= 5
			if sum < 0 {
				fmt.Print(false)
				return
			}
			sum += 5
		}
		if value == 15 {
			sum -= 10
			if sum < 0 {
				fmt.Print(false)
				return
			}
			sum += 5
		}
		if value == 20 {
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

	if resultOfInput == 12 || resultOfInput == 1 || resultOfInput == 2 {
		fmt.Println("Зима")
		return
	} else if resultOfInput >= 3 && resultOfInput <= 5 {
		fmt.Println("Весна")
		return
	} else if resultOfInput >= 6 && resultOfInput <= 8 {
		fmt.Println("Лето")
		return
	} else if resultOfInput >= 9 && resultOfInput <= 11 {
		fmt.Println("Осень")
		return
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

	if resultOfInput != "пн" && resultOfInput != "вт" && resultOfInput != "ср" && resultOfInput != "чт" && resultOfInput != "пт" {
		fmt.Println("Неверный формат данных")
		return
	}

	fmt.Println(weekDays[strings.ToLower(input)])
}

func CalcDiscaunt() {
	var price, discount int

	fmt.Print("Введите цену товара: ")
	fmt.Scan(&price)
	fmt.Print("Введите значение скидки: ")
	fmt.Scan(&discount)

	if discount > 30 {
		fmt.Print("Скидка не более 30%")
		return
	}

	result := (price / 100) * discount

	if result > 2000 {
		fmt.Print("Слишком большая скидка")
		return
	}

	fmt.Print(result)
}

func FillingBuskets() {
	var firstBusket, secondBusket, thirdBusket uint

	fmt.Print("Вместимость первого ведра: ")
	fmt.Scan(&firstBusket)
	fmt.Print("Вместимость второго ведра: ")
	fmt.Scan(&secondBusket)
	fmt.Print("Вместимость третьего ведра: ")
	fmt.Scan(&thirdBusket)

	if firstBusket == 0 || secondBusket == 0 || thirdBusket == 0 {
		fmt.Print("Вместимость каждого ведра должна быть больше нуля")
		return
	}

	var applesInFirstBucket uint = 0
	var applesInSecondBucket uint = 0
	var applesInThirdBucket uint = 0

	for {
		fmt.Println(applesInFirstBucket, applesInSecondBucket, applesInThirdBucket)
		if applesInFirstBucket != firstBusket {
			applesInFirstBucket += 1
		}
		if applesInSecondBucket != secondBusket {
			applesInSecondBucket += 1
		}
		if applesInThirdBucket != thirdBusket {
			applesInThirdBucket += 1
		}

		if firstBusket == applesInFirstBucket && secondBusket == applesInSecondBucket && thirdBusket == applesInThirdBucket {
			break
		}
	}

	fmt.Print(applesInFirstBucket, applesInSecondBucket, applesInThirdBucket)

}

func Lift() {
	floors := 24

	var floorOfFirstPerson,
		floorOfSecondPerson,
		floorOfThirdPerson int

	fmt.Print("На каком этаже находится первый человек? ")
	fmt.Scan(&floorOfFirstPerson)
	fmt.Print("На каком этаже находится второй человек? ")
	fmt.Scan(&floorOfSecondPerson)
	fmt.Print("На каком этаже находится третий человек? ")
	fmt.Scan(&floorOfThirdPerson)

	if floorOfFirstPerson <= 1 || floorOfSecondPerson <= 1 || floorOfThirdPerson <= 1 {
		fmt.Println("Человек должен находиться выше первого этаже")
		return
	}
	if floorOfFirstPerson > floors || floorOfSecondPerson > floors || floorOfThirdPerson > floors {
		fmt.Println("Человек не должен находиться выше 24 этажа")
		return
	}

	isTakeFirstPerson := false
	isTakeSecondPerson := false
	isTakeTHirdPerson := false

	direction := "up"

	numberOfFloor := 1
	numberOfPassengers := 0
	totalPassengersOnFirstFloor := 0

	for {
		if direction == "up" {
			fmt.Printf("номер этажа %d \n", numberOfFloor)
			numberOfFloor += 1

			if numberOfFloor == floors+1 {
				direction = "down"
				continue
			}

			continue
		}

		if direction == "down" {
			numberOfFloor -= 1

			//условия для подбора пассажиров
			if floorOfFirstPerson == numberOfFloor && isTakeFirstPerson == false && numberOfPassengers != 2 {
				isTakeFirstPerson = true
				numberOfPassengers += 1
			}
			if floorOfSecondPerson == numberOfFloor && isTakeSecondPerson == false && numberOfPassengers != 2 {
				isTakeSecondPerson = true
				numberOfPassengers += 1
			}
			if floorOfThirdPerson == numberOfFloor && isTakeTHirdPerson == false && numberOfPassengers != 2 {
				isTakeTHirdPerson = true
				numberOfPassengers += 1
			}

			fmt.Printf("номер этажа %d \n", numberOfFloor)

		}

		if numberOfFloor == 1 {
			totalPassengersOnFirstFloor += numberOfPassengers
			numberOfPassengers = 0
			direction = "up"
		}

		if totalPassengersOnFirstFloor == 3 {
			break
		}
	}

	fmt.Print("Пассажиры доставлены")
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

func RandomGame() {

	rand.Seed(time.Now().Unix())

	var number int
	fmt.Print("загадай число число: ")
	fmt.Scanln(&number)

	trying := 0

	random := rand.Intn(10) + 1

	for {
		trying++
		if trying == 5 {
			break
		}

		fmt.Printf("попытка №%d \n", trying)
		fmt.Printf("число %d \n", random)

		if random == number {
			fmt.Printf("угадал число %d на %d попытке\n", random, trying)
			break
		} else if random < number {
			random = random + rand.Intn(10-random+1)
			continue
		} else if random > number {
			random = rand.Intn(random)
			continue
		}

	}
	fmt.Println("попытки окончены")

}
