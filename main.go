package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	FillingBuskets()
}

func EasyCycle() {
	var num int

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	for i := 0; i <= num; i++ {
		fmt.Printf("%d ", i)
	}
}

func SumOfTwoNumberByOne() {
	var firstNumber, secondNumber int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&secondNumber)

	var sum int = firstNumber

	for i := firstNumber; i <= firstNumber+secondNumber; i++ {
		sum += 1
	}

	fmt.Print(sum)
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

	result := price / discount

	if result > 2000 {
		fmt.Print("Слишком большая скидка")
		return
	}

	fmt.Print(result)
}

func TenHundreadThousand() {
	a := 0
	b := 0
	c := 0

	for {
		if a != 10 {
			a += 1
		}
		if b != 100 {
			b += 1
		}
		if c != 1000 {
			c += 1
		}

		if a == 10 && b == 100 && c == 1000 {
			break
		}
	}

	fmt.Print(a, b, c)
}

func FillingBuskets() {
	var firstBusket, secondBusket, thirdBusket int

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

	var applesInFirstBucket = 0
	var applesInSecondBucket = 0
	var applesInThirdBucket = 0

	for {

		if applesInFirstBucket != firstBusket {
			applesInFirstBucket += 1
			continue
		}
		if applesInSecondBucket != secondBusket {
			applesInSecondBucket += 1
			continue
		}
		if applesInThirdBucket != thirdBusket {
			applesInThirdBucket += 1
			continue
		}

		if firstBusket == applesInFirstBucket && secondBusket == applesInSecondBucket && thirdBusket == applesInThirdBucket {
			break
		}
	}

	fmt.Print(applesInFirstBucket, applesInSecondBucket, applesInThirdBucket)

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
