package storage

import "fmt"

type StudentStorage struct {
	Name  string
	Age   int
	Grade int
}

func GetAllStudents(studentStorage *[]StudentStorage) {
	fmt.Println("")
	for _, v := range *studentStorage {
		fmt.Println(v.Name)
	}
}
