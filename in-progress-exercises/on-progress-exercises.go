package inprogressexercises

import (
	"fmt"
	"strings"
)

func generate(open int, closen int, queue string, result *[]string) {

	// эта проверка осуществляется непосредственно в рекурсии так как на первой входе open != 0
	if open == 0 {
		// ставит просто закрывающие скобки при количестве открытых == 0
		*result = append(*result, queue+strings.Repeat(")", closen))
		return
	}

	// кол-во закрытых больше открытых
	if closen > open {
		generate(open, closen-1, queue+")", result)
	}

	generate(open-1, closen, queue+"(", result)

}

func CombinationOfBrackets(numberOfPares int) {
	result := []string{}

	if numberOfPares == 0 {
		fmt.Print(result)
		return
	}

	generate(numberOfPares, numberOfPares, "", &result)

	fmt.Print(result)

}
