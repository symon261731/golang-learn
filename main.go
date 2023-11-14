package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	FillingBuskets()
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
