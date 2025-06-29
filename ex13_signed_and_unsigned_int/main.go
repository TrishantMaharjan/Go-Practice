package main

import (
	"fmt"
)

func main() {
	// Display highest value of uint8 and int8
	var x uint8 = 255 // Can't take greater than 256 as it can take only 2^8, ie. 256 (0 to 255)
	var y int8 = 127  // Can't take greater than -128 to +127 as it can take only 2^8, ie. 256 (-128 to 127)
	fmt.Printf("%v is of type %T \n", x, x)
	fmt.Printf("%v is of type %T \n", y, y)
}
