package main

import (
	"main/pkg/storage"
	"main/pkg/student"
)

func main() {
	var studentStorage = storage.CreateStorage()

	student.AppAlgorithm(&studentStorage)
}
