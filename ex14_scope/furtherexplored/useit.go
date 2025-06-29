package furtherexplored

import "fmt"

// This is also "package block" scope
// This is exported because the identifier "Fascinating" starts with an uppercase letter
func Fascinating() {
	fmt.Println("In furtherexplored: Planet speed:", planetSpeed, "miles per hour.")
	planetRotationSpeed := 1000
	fmt.Println("In furtherexplored: Planet spinning speed:", planetRotationSpeed, "miles per hour.")
}
