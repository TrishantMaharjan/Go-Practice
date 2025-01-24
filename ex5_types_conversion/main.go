package main

import (
	"fmt"
)

func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// This does not work. In go you can not take a VALUE that is in float 32 and store it in a variable that is declared to hold a VALUE of float 64
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/

	// This works
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)
}
