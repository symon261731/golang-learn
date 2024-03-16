package doneexercise

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//var arrayForSort = []int{1, 2, 4, 6, 8, 9, 5, 3}
//var arrayForSort = []int{1, 2, 3, 4, 5, 6, 7}

func BubbleSort(arrayForSort []int) []int {
	var initialArray = arrayForSort

	for i := 0; i < len(initialArray)-1; i += 1 {
		isAlreadySorted := false

		for j := 0; j < len(initialArray)-1; j += 1 {
			if initialArray[j] > initialArray[j+1] {
				initialArray[j+1], initialArray[j] = initialArray[j], initialArray[j+1]
				isAlreadySorted = true
			}
		}
		//fmt.Println(testArray)
		if !isAlreadySorted {
			break
		}

	}

	return initialArray
}

// 1
func IsOddNumber(num int) bool {
	return num%2 == 0
}

// 2
func CalcCoordinates() (int, int) {
	var x, y int
	fmt.Println("Введите пожалуйста значение X")
	fmt.Scan(&x)
	fmt.Println("Введите пожалуйста значение Y")
	fmt.Scan(&y)

	return prepareCoordinates(x, y)
}

// Вроде будет недоступна вне пакета
func prepareCoordinates(x int, y int) (int, int) {
	var resultX = 2*x + 10
	var resultY = -3*y - 5

	return resultX, resultY
}

// 3
func CalcAddAndMultiplyResult(num int) (int, int) {
	return multiplyNumber(num), addNumber(num)
}

func multiplyNumber(num int) int {
	var multiplyNum int
	fmt.Println("На сколько умножить")
	fmt.Scan(&multiplyNum)
	return multiplyNum * num
}

func addNumber(number int) int {
	var addNumber int
	fmt.Println("Сколько прибавить")
	fmt.Scan(&addNumber)
	return addNumber + number
}

// 4
var globalVariable = 5

func CalcResultOfSumFunctions() int {
	var callFirstFunction = addFirstFunction(globalVariable)
	var callSecondFunction = addSecondFunction(callFirstFunction)
	var callThirdFunction = addThirdFunction(callSecondFunction)

	return callThirdFunction
}

// 5
func CalcSumOfOddAndEvenNumberInArray() (int, int) {
	var array = make([]int, 10)

	for i := 0; i < 10; i++ {

		fmt.Printf("Введите число для индекса %d \n", i)
		fmt.Scan(&array[i])
	}
	var sumOfOdd = 0
	var sumOfEven = 0

	for _, v := range array {
		if v%2 == 0 {
			sumOfOdd += 1
		} else {
			sumOfEven += 1
		}
	}

	return sumOfOdd, sumOfEven
}

// 6
func ReverseArray(arr []int) []int {
	return processReverseArray(arr)
}

func processReverseArray(arr []int) []int {
	var result = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		result[i] = arr[len(arr)-i-1]
	}
	return result
}

func addFirstFunction(num int) int {
	return num + 5
}

func addSecondFunction(num int) int {
	return num + 6
}
func addThirdFunction(num int) int {
	return num + 8
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
			if floorOfFirstPerson == numberOfFloor && !isTakeFirstPerson && numberOfPassengers != 2 {
				isTakeFirstPerson = true
				numberOfPassengers += 1
			}
			if floorOfSecondPerson == numberOfFloor && !isTakeSecondPerson && numberOfPassengers != 2 {
				isTakeSecondPerson = true
				numberOfPassengers += 1
			}
			if floorOfThirdPerson == numberOfFloor && !isTakeTHirdPerson && numberOfPassengers != 2 {
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

func NumberOfWordsInUpperCase() {
	exampleString := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	accumulator := 0

	for _, some := range strings.Split(exampleString, " ") {
		var firstLetter = fmt.Sprintf("%c", some[0])

		if strings.ToLower(firstLetter) != firstLetter {
			accumulator += 1
		}
	}

	fmt.Println(accumulator)
}

func NumbersInTheString() {
	example := "a10 10 20b 20 30c30 30 dd 50"

	for _, some := range strings.Split(example, " ") {
		number, err := strconv.Atoi(some)

		if err == nil {
			fmt.Printf("%d ", number)
		}
	}
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

func ModifySpaces(s, mode string) string {

	var result string

	switch mode {
	case "dash":
		result = strings.ReplaceAll(s, " ", "-")
	case "underscore":
		result = strings.ReplaceAll(s, " ", "_")
	default:
		result = strings.ReplaceAll(s, " ", "*")
	}

	return result

}

func SafeWrite(nums [5]int, i, val int) [5]int {

	var result = [5]int{}
	result[0] = nums[0]
	result[1] = nums[1]
	result[2] = nums[2]
	result[3] = nums[3]
	result[4] = nums[4]

	if i < 5 && i >= 0 {
		result[i] = val
	}

	return result

}

func Remove(nums []int, i int) []int {

	var result []int

	if i > len(nums) || i < 0 {
		return nums
	}

	for x := 0; x < len(nums); x++ {
		if x != i {
			result = append(result, nums[x])
		}
	}

	return result
}

func IntsCopy(src []int, maxLen int) []int {

	var result []int

	if maxLen <= 0 {
		return []int{}
	}

	if maxLen > len(src) {
		result = make([]int, len(src))
		copy(result, src)

		return result
	}

	result = make([]int, maxLen)

	copy(result, src)

	return result

}

func UniqueSortedUserIDsCorrectSolution(userIDs []int64) []int64 {

	if len(userIDs) < 2 {
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]

}

func UniqueSortedUserIDsMyOpinion(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.Slice(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })

	var result []int64

	for i := 0; i < len(userIDs); i++ {
		duplicate := false

		for y := 0; y < len(result); y++ {
			if userIDs[i] == result[y] {
				duplicate = true
			}
		}
		if !duplicate {
			result = append(result, userIDs[i])
		}
	}

	if len(result) == 0 {
		return []int64{}
	}

	return result
}

func UniqueUserIDs(userIDs []int64) []int64 {
	var result []int64

	hash := make(map[int64]struct{})

	for _, elem := range userIDs {
		_, ok := hash[elem]

		if !ok {
			result = append(result, elem)
			hash[elem] = struct{}{}
		}
	}

	if len(result) == 0 {
		return []int64{}
	}

	return result
}

func MostPopularWord(words []string) string {

	var hash = map[string]int{}

	for _, str := range words {

		if hash[str] == 0 {
			hash[str] = 1
		} else {
			hash[str] = hash[str] + 1
		}
	}

	var maxNumber int = 0
	var popularName string
	for word, value := range hash {
		if value > maxNumber {
			maxNumber = value
			popularName = word
		}
	}

	fmt.Println(popularName)
	fmt.Println(hash)

	return popularName
}

func ShiftASCII(s string, step int) string {

	hashOfEntryString := []byte(s)
	var hashOfNewStringInBytes []byte

	for _, val := range hashOfEntryString {
		hashOfNewStringInBytes = append(hashOfNewStringInBytes, val+byte(step))
	}

	return string(hashOfNewStringInBytes)
}

func IsASCII(s string) bool {
	MaxASCII := '\u007F'

	for _, run := range s {

		if run > MaxASCII {
			return false
		}
	}

	return true
}

func LatinLetters(s string) string {

	result := strings.Builder{}

	for _, val := range s {
		if unicode.Is(unicode.Latin, val) {
			result.WriteRune(val)
		}
	}

	return result.String()
}

func MergeNumberLists(numberList ...[]int) []int {
	var result []int

	for _, value := range numberList {
		result = append(result, value...)
	}

	return result
}

func GetValuesByLink(first *int, second *int) {
	var valueFirst = *first
	var valueSecond = *second

	*first = valueSecond
	*second = valueFirst

}

func CalcSumOfOddNumbersInRow(firstNum int, secondNum int) {
	sum := 0

	if firstNum < secondNum {
		for i := firstNum; i <= secondNum; i++ {
			if i%2 == 0 {
				sum += i
			}
		}
	} else {
		for i := secondNum; i <= firstNum; i++ {
			if i%2 == 0 {
				sum += i
			}
		}
	}

	fmt.Print(sum)
}
