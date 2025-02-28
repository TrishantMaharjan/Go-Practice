package main

import (
	"fmt"
)

type ByteSize = uint

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("1 KB = \t %d \n", KB)
	fmt.Printf("1 MB = \t %d \n", MB)
	fmt.Printf("1 GB = \t %d \n", GB)
	fmt.Printf("1 TB = \t %d \n", TB)
	fmt.Printf("1 PB = \t %d \n", PB)
	fmt.Printf("1 EB = \t %d \n", EB)
}
