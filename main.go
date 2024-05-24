package main

import (
	"main/pkg/storage"
	"main/pkg/student"
)

func main() {
	studentStorageVariable := make([]storage.StudentStorage, 0)

	student.AppAlgorithm(&studentStorageVariable)
}
