package main

import "fmt"

func main() {
	CalcDiscountInRestaurant()
}

func ExamResult() {
	neededResult := 275
	var firstExam int
	var secondExam int
	var thirdExam int
	fmt.Print("Результаты первого экзамена: ")
	fmt.Scan(&firstExam)

	fmt.Print("Результаты второго экзамена экзамена: ")
	fmt.Scan(&secondExam)

	fmt.Print("Результаты третьего экзамена: ")
	fmt.Scan(&thirdExam)

	if firstExam+secondExam+thirdExam < neededResult {
		fmt.Println("Вы не поступили")
	} else {
		fmt.Println("Вы поступили")
	}

}

func CheckVarMoreFive() {
	var firstValue int
	var secondValue int
	var thirdValue int

	fmt.Print("Пожалуйста введите числа: ")
	fmt.Scan(&firstValue, &secondValue, &thirdValue)

	if firstValue > 5 || secondValue > 5 || thirdValue > 5 {
		fmt.Println("Среди введённых чисел есть число больше 5.")
	} else {
		fmt.Println("Среди введенных чисел нет числа больше 5")
	}
}

func CalcATMMoney() {
	maxValue := 100000
	var inputValue int
	fmt.Print("Введите суммы которую хотите снять: ")
	fmt.Scan(&inputValue)

	if inputValue > maxValue || inputValue <= 0 {
		fmt.Println("неккоректная сумма")
	} else {
		if (inputValue % 100) == 0 {
			fmt.Println("Операция выполнена успешно")
		} else {
			fmt.Println("Сумма должна быть кратная 100")
		}
	}
}

func ValueOfIntMoreThenFive() {
	var firstValue int
	var secondValue int
	var thirdValue int

	var arr [3]int
	numberOfTrueValues := 0

	fmt.Print("Пожалуйста введите числа: ")
	fmt.Scan(&firstValue, &secondValue, &thirdValue)

	arr[0] = firstValue
	arr[1] = secondValue
	arr[2] = thirdValue
	for i := 0; i < len(arr); i += 1 {
		if arr[i] >= 5 {
			numberOfTrueValues += 1
		}
	}
	fmt.Printf("Среди введеных чисел %d больше или равны 5", numberOfTrueValues)
}

func CalcDiscountInRestaurant() {
	var dayOfWeek int
	var numberOfGuests int
	var sumOfCheck int

	procentDiscount := 1.00
	serviceSurcharge := 1.00

	fmt.Print("Введите день недели: ")
	fmt.Scan(&dayOfWeek)

	fmt.Print("Введите количество гостей: ")
	fmt.Scan(&numberOfGuests)

	fmt.Print("Введите сумму чека: ")
	fmt.Scan(&sumOfCheck)

	if dayOfWeek < 1 || dayOfWeek > 7 {
		fmt.Println("неккоректный день недели")
	} else {
		if dayOfWeek == 1 {
			procentDiscount = 0.9
		}
		if dayOfWeek == 5 && sumOfCheck > 10000 {
			procentDiscount = 0.95
		}

		if numberOfGuests > 5 {
			serviceSurcharge = 1.10
		}

		formula := float64(sumOfCheck) * serviceSurcharge * procentDiscount

		fmt.Println(int64(formula))

	}
}

func CipherCaesar() {
	cipherCaesar := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(cipherCaesar); i += 1 {
		letter := cipherCaesar[i]
		if letter > 'a' && letter < 'z' {
			letter = letter - 3
			if letter < 'a' {
				letter += 26
			}
		} else if letter > 'A' && letter <= 'Z' {
			letter = letter - 3
			if letter < 'A' {
				letter = letter + 26
			}
		}
		fmt.Printf("%c", letter)
	}
}
