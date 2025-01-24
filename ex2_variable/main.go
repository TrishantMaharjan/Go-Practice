package main

import "fmt"

func main() {
	var x int
	fmt.Printf("%d is an integer\n", x)
	x = 5
	fmt.Printf("%d is an integer\n", x)

	a := 6
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "Test"
	fmt.Println(b, c, d, f)

	// Defining variable without using it gives an error
}
