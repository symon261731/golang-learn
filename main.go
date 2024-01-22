package main

import "fmt"

func main() {
	var firstNum = 4
	var secondNum = 10
	fmt.Println(firstNum, secondNum)

	GetValuesByLink(&firstNum, &secondNum)

	fmt.Print(firstNum, secondNum)

}

func GetValuesByLink(first *int, second *int) {
	var valueFirst = *first
	var valueSecond = *second

	*first = valueSecond
	*second = valueFirst

}

func CalcSumOfOddNumbersInRow(firstNum int, secondNum int) {
	sum := 0

	if firstNum < secondNum {
		for i := firstNum; i <= secondNum; i++ {
			if i%2 == 0 {
				sum += i
			}
		}
	} else {
		for i := secondNum; i <= firstNum; i++ {
			if i%2 == 0 {
				sum += i
			}
		}
	}

	fmt.Print(sum)
}

// лёгкие:
// Задание 1. Функция, принимающая аргументы
// Напишите функцию, которая принимает в качестве аргументов два числа типа int, вычисляет сумму чётных чисел заданного диапазона и выводит результат в консоль.

// Рекомендация
// Если первое число, переданное в качестве аргумента, будет больше второго, просто поменяйте их местами.

// Задание 2. Функция, принимающая значение по ссылке
// Напишите функцию, которая принимает в качестве аргументов указатели на два типа int и меняет их значения местами.

// Рекомендация
// В методе main создайте и присвойте значения двум переменным типа int, выведите значения этих переменных в консоль до вызова функции и после.
