package main

import (
	"fmt"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places, ie. the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99

	// Left Shift and Right Shift operators (bit shifting)
	Num1 = 20 >> 2 // 0001 0100, take out last 2 digits, gives 0101, which is 5
	Num2 = 20 << 2 // 0001 0100, add 2 digits to the last, gives 0101 0000, which is 80
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// Numeric Constants
	// An untyped constant takes the type needed by it's context
	// fmt.Println(needInt(Big)) -> Gives error cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to needInt (overflows), in int can sotre at max a 64-bit integer and sometimes less
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(Num1)
	fmt.Println(Num2)

	fmt.Printf("%d \t %b \n", 1, 1)
	fmt.Printf("%d \t %b \n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b \n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b \n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b \n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b \n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b \n", 1<<6, 1<<6)

	fmt.Printf("%d \t %b \n", 64, 64)
	fmt.Printf("%d \t %b \n", 64>>1, 64>>1)
	fmt.Printf("%d \t %b \n", 64>>2, 64>>2)
	fmt.Printf("%d \t %b \n", 64>>3, 64>>3)
	fmt.Printf("%d \t %b \n", 64>>4, 64>>4)
	fmt.Printf("%d \t %b \n", 64>>5, 64>>5)
	fmt.Printf("%d \t %b \n", 64>>6, 64>>6)
}
