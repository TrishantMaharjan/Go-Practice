package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// Print these values as both binary and hexadecimal

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("Number \t Binary \t Hexadecimal \n")
	fmt.Printf("%v \t %b \t \t %#X \n", a, a, a)
	fmt.Printf("%v \t %b \t \t %#X \n", b, b, b)
	fmt.Printf("%v \t %b \t \t %#X \n", c, c, c)
	fmt.Printf("%v \t %b \t \t %#X \n", d, d, d)
	fmt.Printf("%v \t %b \t \t %#X \n", e, e, e)
	fmt.Printf("%v \t %b \t \t %#X \n", f, f, f)
}
