package inprogressexercises

import "fmt"

func CombinationOfBrackets(numberOfPares int) {
	result := []string{}

	if numberOfPares == 0 {
		fmt.Print(result)
		return
	}

	if numberOfPares == 1 {
		result = []string{"()"}
		fmt.Print(result)
		return
	}

	if numberOfPares > 1 {
		fmt.Print("Ну тут что то будет")
		return
	}

}
