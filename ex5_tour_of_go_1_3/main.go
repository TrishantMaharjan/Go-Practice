package main

import (
	// list of packages
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old") // function is defined using parameters and used with arguments

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(6)) // Random number from 0 to 5, [0,n) -> inclusive of 0, exclusive of n

	fmt.Printf("Now you have %v problems \n", math.Sqrt(9))

	fmt.Println(math.Pi) // In go, a name is exported if it begins with a capital letter. Here, Pi is a constant, so () is not used. () is used to execute a function

	// In go, it's not public and private, it's exported and unexported
}
