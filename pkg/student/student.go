package student

import (
	"bufio"
	"fmt"
	"main/pkg/storage"
	"os"
	"strconv"
	"strings"
)

func AppAlgorithm(StudentStorageVariable *[]storage.StudentStorage) {
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

		AddStudent(StudentStorageVariable, storage.StudentStorage{Name: studentInfo[0], Age: age, Grade: grade})
	}

	storage.GetAllStudents(StudentStorageVariable)
}

func AddStudent(storage *[]storage.StudentStorage, info storage.StudentStorage) {
	*storage = append(*storage, info)
}
