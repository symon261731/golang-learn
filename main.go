package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Введите имя и возраст через пробел: ")
	fmt.Scan(&name, &age)

	fmt.Println("My name is", name, "and i am", age, "years old")
}
