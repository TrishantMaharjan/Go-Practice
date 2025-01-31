package main

import (
	"fmt"
)

func add(x, y int) int { // Here, x and y share a same type so add(x int, y int) was shortned to add(x, y int)
	return x + y
}

func sayHello() {
	fmt.Println("Hello!!")
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) /* Parameters */ (x, y int) { // Named return values, this should be used only in short functions as they can harm readability in longer functions
	x = sum % 10
	y = sum - x
	return // no need to return here if returned when initializing the function
}

func main() {
	fmt.Println(add(42, 13))
	sayHello()
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(10)) // Arguments
}
