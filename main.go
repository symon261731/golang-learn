package main

import "main.go/ForCheck"

func main() {
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}

	chars := []rune{'H', 'E', 'l', 'П', 'М'}

	ForCheck.FindNumberOfSymbols(sentences, chars)

}
