package storage

import "fmt"

type StudentStorage struct {
	Name  string
	Age   int
	Grade int
}

func CreateStorage() []StudentStorage {
	studentStorageVariable := make([]StudentStorage, 0)
	return studentStorageVariable
}

func GetAllStudents(studentStorage *[]StudentStorage) {
	fmt.Println("")
	for _, v := range *studentStorage {
		fmt.Println(v.Name)
	}
}
