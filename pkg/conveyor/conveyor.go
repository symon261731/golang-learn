package conveyor

import (
	"fmt"
	"strconv"
	"sync"
)

func AppAlgorithm() {
	var wg sync.WaitGroup

	flow := make(chan int)

	for {
		var text string
		fmt.Scan(&text)

		if text == "стоп" {
			close(flow)
			break
		}

		num, err := strconv.Atoi(text)

		if err != nil {
			fmt.Println("Ошибка при изменении формата данных")
			break
		}

		wg.Add(1)
		go SquareNumber(flow, &wg)

		flow <- num

		squareResult := <-flow
		fmt.Println(squareResult)

		wg.Add(1)
		go MultiplyNumberOnTwo(flow, &wg)

		flow <- squareResult

		multiplyResult := <-flow
		fmt.Println(multiplyResult)

	}
	wg.Wait()

}

func SquareNumber(flow chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := <-flow
	squareNumber := num * num
	flow <- squareNumber
}

func MultiplyNumberOnTwo(flow chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := <-flow
	result := num * 2
	flow <- result
}
