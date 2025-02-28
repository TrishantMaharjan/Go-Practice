package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	// Type inference
	v := true // Change me!
	fmt.Printf("v is of type %T \n", v)
	/*
		i := 42 -> int
		f := 3.142 -> float64
		g := 0.867 + 0.5i
		b := true
		s := "Test"
	*/

	// Constants
	fmt.Println("Happy", Pi, "Day")
	fmt.Printf("Pi is of type %T \n", Pi)
}
