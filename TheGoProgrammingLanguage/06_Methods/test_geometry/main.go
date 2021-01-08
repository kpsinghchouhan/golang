// Test_geometry test geometry package.
package main

import (
	"fmt"
	"golang/TheGoProgrammingLanguage/06_Methods/geometry"
)

func main() {
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance())
}
