package main

import (
	"fmt"
	"os"
)

func main() {
	WorkWithFileSystem()
}

func WorkWithFileSystem() {
	var nameOfFiles []string

	dir, err := os.Open("../forLearn")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		nameOfFiles = append(nameOfFiles, file.Name())
	}

	fmt.Println(nameOfFiles)
	fmt.Println(len(nameOfFiles))
}
