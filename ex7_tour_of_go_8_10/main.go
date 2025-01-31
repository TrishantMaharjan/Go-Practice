package main

import (
	"fmt"
)

var c, python, java bool

const d = 42         // untyped constant
const e float32 = 24 // typed constant

// More details about constants in https://medium.com/@dpshiv12/go-constants-fa29270bd170

// Outside a function, short declaration operator ie := is not available, as every statement begins with a const, var, func, and so on

func main() {
	var i int
	fmt.Println(i, c, python, java, d, e)
	fmt.Printf("%v of type %T \n", d, d)
	fmt.Printf("%v of type %T \n", e, e)

	f := 3
	fmt.Println(f)
}
