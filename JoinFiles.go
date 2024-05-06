package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	argumentsAfterGoCommand := os.Args[1:]
	buf := bytes.Buffer{}

	info, err := os.ReadFile(argumentsAfterGoCommand[0])

	if err != nil {
		fmt.Println("error with first file")
		return
	}

	for _, b := range info {
		buf.WriteByte(b)
	}

	infoSecond, err := os.ReadFile(argumentsAfterGoCommand[1])

	if err != nil {
		fmt.Println("error with second file")
		return
	}

	for _, b := range infoSecond {
		buf.WriteByte(b)
	}

	resultFile, err := os.Create(argumentsAfterGoCommand[2])
	defer resultFile.Close()

	if err != nil {
		fmt.Println(buf.String())
		return
	}

	resultFile.WriteString(buf.String())

}
