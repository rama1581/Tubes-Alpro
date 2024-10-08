package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{5, 2, 7, 1, 8, 4}
	fmt.Println("Before sorting: ", list)
	sort.Ints(list)
	fmt.Println("After sorting: ", list)

	median := list[len(list)/2]
	fmt.Println("Median: ", median)
}
