package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	WorkWithFiles()
}

func WorkWithFiles() {

	var input string

	for i := 1; input != "exit"; i++ {
		fmt.Print("Введите пожалуйста деятельность: ")
		fmt.Scan(&input)

		currentDate := time.Now().Day()
		currentMonth := time.Now().Month()
		currentYear := time.Now().Year()
		currentHour := time.Now().Hour()
		currentMinute := time.Now().Minute()
		currentSecond := time.Now().Second()

		file, err := os.OpenFile("./workWithFile", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

		if err != nil {
			fmt.Print("Error while create file")
			os.Exit(1)
		}
		defer file.Close()

		formatString := fmt.Sprintf(
			"№%d %d-%d-%d %d:%d:%d %s \n",
			i,
			currentYear,
			currentMonth,
			currentDate,
			currentHour,
			currentMinute,
			currentSecond,
			input,
		)

		_, err = file.WriteString(formatString)

		if err != nil {
			fmt.Print("Error")
			os.Exit(1)
		}

	}

	return

}
