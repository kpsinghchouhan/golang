package treesort_test

import (
	"fmt"
	"golang/TheGoProgrammingLanguage/04_CompositeTypes/treesort"
)

func main() {
	values := []int{1, 5, 2, 3, 4}
	treesort.Sort(values)
	fmt.Println(values)
}
