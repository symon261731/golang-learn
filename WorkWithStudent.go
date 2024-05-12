package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StudentStorage struct {
	name  string
	age   int
	grade int
}

func main() {
	var studentStorage []StudentStorage
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите данные студента (имя возраст оценка), затем нажмите Enter. Для завершения введите EOF (ctrl + d).")
	for scanner.Scan() {

		text := scanner.Text()

		if text == "exit" {
			break
		}

		studentInfo := strings.Split(text, " ")
		age, err := strconv.Atoi(studentInfo[1])

		if err != nil {
			fmt.Println("Неверный формат данных")
			continue
		}

		grade, err := strconv.Atoi(studentInfo[2])

		if err != nil {
			fmt.Println("Неверный формат данных")
			continue
		}

		studentStorage = append(studentStorage, StudentStorage{name: studentInfo[0], age: age, grade: grade})
	}

	for _, v := range studentStorage {
		fmt.Println(v.name)
	}
}
