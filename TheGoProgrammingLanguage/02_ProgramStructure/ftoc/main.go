// Ftoc prints two Fahrenheit-toCelcius conversions.
package main

import "fmt"

func main() {
	const freeingzF, boilingF = 32.0, 212.0
	fmt.Printf("%g degree F = %g degree C\n", freeingzF, fToC(freeingzF)) // "32 F = 0 C"
	fmt.Printf("%g degree F = %g degree C\n", boilingF, fToC(boilingF))   // "212 F = 100 C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
