// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g degree F or %g degree C\n", f, c)
	// Output:
	// boling point = 212 F or 100 C
}
