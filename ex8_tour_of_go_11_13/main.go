package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// Basic types
	fmt.Printf("Type: %T Value %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v \n", z, z)

	// Zero value
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q \n", i, f, b, s) // Variables declared without an explicit initial value are given their zero value

	// Type conversion
	var x, y int = 3, 4
	var g float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(g) // The expression T(v) converts the value v into the type T
	fmt.Println(x, y, z)
}
