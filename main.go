package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	LuckyTicket()
}

func CalcCoordinateArea() {
	var x, y int

	fmt.Print("Введите значение точки X: ")
	fmt.Scan(&x)

	fmt.Print("Введите значение точки Y: ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("точка находиться в 1-ой области")
	} else if x < 0 && y > 0 {
		fmt.Println("точка находиться во 2-ой области")
	} else if x < 0 && y < 0 {
		fmt.Println("точка находиться в 3-ей области")
	} else if x > 0 && y < 0 {
		fmt.Println("точка находиться в 4-ой области")
	}

	if x > 0 && y == 0 {
		fmt.Println("точка находиться между 1-ой и 4-ой областью")
	} else if x < 0 && y == 0 {
		fmt.Println("точка находиться между 2-ой и 3-ей областью")
	}

	if x == 0 && y > 0 {
		fmt.Println("точка находиться между 1-ой и 2-ой областью")
	} else if x == 0 && y < 0 {
		fmt.Println("точка находиться между 3-ей и 4-ой областями")
	}

	if x == 0 && y == 0 {
		fmt.Println("точка находиться между всеми")
	}
}

func checkPositiveInput() {
	var firstValue int
	var secondValue int
	var thirdValue int
	result := 0

	fmt.Print("введите значения: ")
	fmt.Scan(&firstValue, &secondValue, &thirdValue)

	arr := [3]int{firstValue, secondValue, thirdValue}

	for i := 0; i < len(arr); i += 1 {
		if arr[i] > 0 {
			result += 1
		}
	}
	fmt.Printf("среди введенных значений %d являются положительными", result)
}

func DetectMatchingNumbers() {
	var a, b, c float64

	fmt.Print("Введите 3 числа: ")
	fmt.Scan(&a, &b, &c)

	if a == b || a == c || b == c {
		fmt.Println("Есть совпадение")
	} else {
		fmt.Println("Все числа уникальны")
	}
}

func LuckyTicket() {
	var ticket int
	fmt.Print("Введите номер билета: ")
	fmt.Scan(&ticket)

	if ticket < 1000 || ticket > 9999 {
		fmt.Println("Неккоректный билет")
		return
	}

	firstLetter := ticket / 1000 % 10
	secondLetter := ticket / 100 % 10
	thirdLetter := ticket / 10 % 10
	fourthLetter := ticket % 10

	if firstLetter == fourthLetter && secondLetter == thirdLetter {
		fmt.Println("Зеркальный номер")
		return
	}
	if firstLetter+secondLetter == thirdLetter+fourthLetter {
		fmt.Println("Счастливый номер")
		return
	}
	fmt.Println("Обычный номер")

}

func CalcQuadraticEquation() {
	var a, b, c float64

	fmt.Print("Введите переменные a,b,c: ")
	fmt.Scan(&a, &b, &c)

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant > 0 {
		resolvePositive := (-b + math.Sqrt(discriminant)) / 2 * a
		resolceNegative := (-b - math.Sqrt(discriminant)) / 2 * a
		fmt.Printf("Решение уравнения: %4.2f  %4.2f", resolvePositive, resolceNegative)
	} else if discriminant == 0 {
		resolve := -(b / 2 * a)
		fmt.Printf("Решение уравнения: %4.2f", resolve)
	} else {
		fmt.Println("нет действительных корней")
	}

}

func RandomGame() {

	for try := 1; try <= 4; try += 1 {
		randomNumber := rand.Intn(10)

		fmt.Printf("ваше число %d верно? ", randomNumber)

		var answer string
		fmt.Scan(&answer)

		if answer == "да" || answer == "нет" {
			if answer == "да" {
				fmt.Printf("число %d угадано", randomNumber)
				break
			} else {
				continue
			}

		} else {
			fmt.Println("более коректный ответ пожалуйста да нет")
			return
		}

	}

	fmt.Println("попытки окончены")
}

//func ExamResult() {
//	neededResult := 275
//	var firstExam int
//	var secondExam int
//	var thirdExam int
//	fmt.Print("Результаты первого экзамена: ")
//	fmt.Scan(&firstExam)
//
//	fmt.Print("Результаты второго экзамена экзамена: ")
//	fmt.Scan(&secondExam)
//
//	fmt.Print("Результаты третьего экзамена: ")
//	fmt.Scan(&thirdExam)
//
//	if firstExam+secondExam+thirdExam < neededResult {
//		fmt.Println("Вы не поступили")
//	} else {
//		fmt.Println("Вы поступили")
//	}
//
//}
//
//func CheckVarMoreFive() {
//	var firstValue int
//	var secondValue int
//	var thirdValue int
//
//	fmt.Print("Пожалуйста введите числа: ")
//	fmt.Scan(&firstValue, &secondValue, &thirdValue)
//
//	if firstValue > 5 || secondValue > 5 || thirdValue > 5 {
//		fmt.Println("Среди введённых чисел есть число больше 5.")
//	} else {
//		fmt.Println("Среди введенных чисел нет числа больше 5")
//	}
//}
//
//func CalcATMMoney() {
//	maxValue := 100000
//	var inputValue int
//	fmt.Print("Введите суммы которую хотите снять: ")
//	fmt.Scan(&inputValue)
//
//	if inputValue > maxValue || inputValue <= 0 {
//		fmt.Println("неккоректная сумма")
//	} else {
//		if (inputValue % 100) == 0 {
//			fmt.Println("Операция выполнена успешно")
//		} else {
//			fmt.Println("Сумма должна быть кратная 100")
//		}
//	}
//}
//
//func ValueOfIntMoreThenFive() {
//	var firstValue int
//	var secondValue int
//	var thirdValue int
//
//	var arr [3]int
//	numberOfTrueValues := 0
//
//	fmt.Print("Пожалуйста введите числа: ")
//	fmt.Scan(&firstValue, &secondValue, &thirdValue)
//
//	arr[0] = firstValue
//	arr[1] = secondValue
//	arr[2] = thirdValue
//	for i := 0; i < len(arr); i += 1 {
//		if arr[i] >= 5 {
//			numberOfTrueValues += 1
//		}
//	}
//	fmt.Printf("Среди введеных чисел %d больше или равны 5", numberOfTrueValues)
//}
//
//func CalcDiscountInRestaurant() {
//	var dayOfWeek int
//	var numberOfGuests int
//	var sumOfCheck int
//
//	procentDiscount := 1.00
//	serviceSurcharge := 1.00
//
//	fmt.Print("Введите день недели: ")
//	fmt.Scan(&dayOfWeek)
//
//	fmt.Print("Введите количество гостей: ")
//	fmt.Scan(&numberOfGuests)
//
//	fmt.Print("Введите сумму чека: ")
//	fmt.Scan(&sumOfCheck)
//
//	if dayOfWeek < 1 || dayOfWeek > 7 {
//		fmt.Println("неккоректный день недели")
//	} else {
//		if dayOfWeek == 1 {
//			procentDiscount = 0.9
//		}
//		if dayOfWeek == 5 && sumOfCheck > 10000 {
//			procentDiscount = 0.95
//		}
//
//		if numberOfGuests > 5 {
//			serviceSurcharge = 1.10
//		}
//
//		formula := float64(sumOfCheck) * serviceSurcharge * procentDiscount
//
//		fmt.Println(int64(formula))
//
//	}
//}
//
//func CipherCaesar() {
//	cipherCaesar := "L fdph, L vdz, L frqtxhuhg."
//
//	for i := 0; i < len(cipherCaesar); i += 1 {
//		letter := cipherCaesar[i]
//		if letter > 'a' && letter < 'z' {
//			letter = letter - 3
//			if letter < 'a' {
//				letter += 26
//			}
//		} else if letter > 'A' && letter <= 'Z' {
//			letter = letter - 3
//			if letter < 'A' {
//				letter = letter + 26
//			}
//		}
//		fmt.Printf("%c", letter)
//	}
//}
