package main

func main() {
	SomeFunc(func(x int, y int) int {
		return x + y
	})

	SomeFunc(func(i int, i2 int) int {
		return i * i2
	})

	SomeFunc(func(x int, y int) int {
		return x*x + y*y
	})

}

func SomeFunc(cb func(someX int, someY int) int) {

	defer cb(10, 15)
}
