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

	writeDataInBuffer(&buf, info)

	if len(argumentsAfterGoCommand) == 1 {
		fmt.Println(buf.String())
		return
	}

	infoSecond, err := os.ReadFile(argumentsAfterGoCommand[1])

	if err != nil {
		fmt.Println("error with second file")
		return
	}

	writeDataInBuffer(&buf, infoSecond)

	if len(argumentsAfterGoCommand) != 3 {
		fmt.Println(buf.String())
		return
	}

	resultFile, err := os.Create(argumentsAfterGoCommand[2])
	defer resultFile.Close()

	if err != nil {
		fmt.Println("error")
		return
	}

	resultFile.WriteString(buf.String())
}

func writeDataInBuffer(buffer *bytes.Buffer, fileInfo []byte) {
	for _, b := range fileInfo {
		buffer.WriteByte(b)
	}
}
