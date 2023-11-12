package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	CalcGroupsOfStudents()
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

func CalcMoneyWithoutChange() {
	
	var totalMoney, currencyFirst, currencySecond, currencyThird int 
	var minVal, midVal, maxVal,
		count1, count2, count3 int

	fmt.Print("Введите стоимость товара: ")
	fmt.Scan(&totalMoney)
	fmt.Print("Введите значение трех валют: ")
	fmt.Scan(&currencyFirst, &currencySecond, &currencyThird)

	if currencyFirst < currencySecond && currencyFirst < currencyThird {
		minVal = currencyFirst
		if currencySecond < currencyThird {
			midVal = currencySecond
			maxVal = currencyThird
		} else {
			midVal = currencyThird
			maxVal = currencySecond
		}
	}
	if currencySecond < currencyFirst && currencySecond < currencyThird {
		minVal = currencySecond
		if currencyFirst < currencyThird {
			midVal = currencyFirst
			maxVal = currencyThird
		} else {
			midVal = currencyThird
			maxVal = currencyFirst
		}
	}
	if currencyThird < currencyFirst && currencyThird < currencySecond {
		minVal = currencyThird
		if currencyFirst < currencySecond {
			midVal = currencyFirst
			maxVal = currencySecond
		} else {
			midVal = currencySecond
			maxVal = currencyFirst
		}
	}

	flag := false
	//количество монет с наибольшей купюрой
	count3 = totalMoney / maxVal
	for {
		// если остаток от наибольшей купюры = 0 то прервать цикл
		if totalMoney%maxVal == 0 {
			break
		}
		count2 = (totalMoney - count3*maxVal) / midVal
		for {
			if (totalMoney-count3*maxVal)%midVal == 0 {
				flag = true
				break
			}
			if (totalMoney-count3*maxVal-count2*midVal)%minVal == 0 {
				count1 = (totalMoney - count3*maxVal - count2*midVal) / minVal
				flag = true
				break
			}
			if count2 != 0 {
				count2--
				continue
			}
			if count3 != 0 {
				count3--
				break
			}
			fmt.Println("\nНе получается оплатить без сдачи имеющимися монетами")
			return
		}
		if flag {
			break
		}
	}

	fmt.Println("\nДля оплаты необходимо:\n",
		count3, "монет номиналом", maxVal, "руб. и\n",
		count2, "монет номиналом", midVal, "руб. и\n",
		count1, "монет номиналом", minVal, "руб.")

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
