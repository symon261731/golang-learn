package ForCheck

import "fmt"

func isOddNumber(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

type CoordinatesStruct struct {
	x int
	y int
}

func CalcCoordinates() CoordinatesStruct {
	var x, y int
	fmt.Println("Введите пожалуйста значение X")
	fmt.Scan(&x)
	fmt.Println("Введите пожалуйста значение Y")
	fmt.Scan(&y)
	var result = prepareCoordinates(x, y)

	return result
}

// Вроде будет недоступна вне пакета
func prepareCoordinates(x int, y int) CoordinatesStruct {
	var resultX = 2*x + 10
	var resultY = -3*y - 5

	var coordinates CoordinatesStruct
	coordinates.x = resultX
	coordinates.y = resultY

	return coordinates

}
