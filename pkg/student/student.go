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
		studentInfo := ReadStudentInfo(scanner)
		AddStudent(StudentStorageVariable, studentInfo)
	}

	storage.GetAllStudents(StudentStorageVariable)
}

func ReadStudentInfo(scanner *bufio.Scanner) storage.StudentStorage {
	text := scanner.Text()

	studentInfo := strings.Split(text, " ")
	age, err := strconv.Atoi(studentInfo[1])

	if err != nil {
		fmt.Println("Неверный формат данных")
	}

	grade, err := strconv.Atoi(studentInfo[2])

	if err != nil {
		fmt.Println("Неверный формат данных")
	}

	newStudent := storage.StudentStorage{Age: age, Grade: grade, Name: studentInfo[0]}

	return newStudent
}

func AddStudent(storage *[]storage.StudentStorage, info storage.StudentStorage) {
	*storage = append(*storage, info)
}
