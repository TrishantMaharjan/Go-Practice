package main

import (
	"exercise-14/furtherexplored" // "<module name in go.mod>/package name ie. in this case furtherexplored"
	"fmt"
)

// x is in "package block" scope
var x = 42

func main() {
	// x can be accessed here
	fmt.Println("x in main:", x)

	// x can be accessed in the function below
	printMe()

	// y doesn't exist here, so this won't work
	// fmt.Println(y)

	// y is in the "block" scope
	y := 43
	fmt.Println("y in main:", y)

	p1 := person{
		"Alice",
		9841000001,
	}
	p1.sayHello()

	// variable "shadowing" example
	x := 32
	fmt.Println("x in main after shadowing:", x)

	// package block x is still the same
	printMe()

	furtherexplored.Fascinating()
	// fmt.Println("This is will not work as the variable starts with small letter", furtherexplored.planetSpeed)
	fmt.Println("This is exported as the variable starts with capital letter", furtherexplored.PlanetSpeed)

	if z := 82; z > 50 {
		fmt.Println("z is greater than 50:", z)
	}
	// z is not accessible here, so this won't work
	// fmt.Println("z in main:", z)
	// This is "comma ok" idiom in https://go.dev/doc/effective_go#maps
}

func printMe() {
	// x can be accessed here
	fmt.Println("x in printMe:", x)

	// y doesn't exist here, so this won't work
	// fmt.Println(y)

	// y is not defined in this scope, so this will cause an error
	// y := 44
	// fmt.Println("y in printMe:", y)
}

type person struct {
	first string
	phone int
}

// The method sayHello, which is attached to VALUES of TYPE person is in "package block" scope
func (p person) sayHello() {
	fmt.Println("In sayHello: Hello, my name is", p.first)
	fmt.Println("In sayHello: Hello, my number is is", p.phone)
}
