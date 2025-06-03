package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	original := []int{a, b, c}

	sorted := make([]int,3)
	copy(sorted, original)

	sort.Ints(sorted)

	for _, value := range sorted{
		fmt.Println(value)
	}

	fmt.Println()

	for _, value := range original{
		fmt.Println(value)
	}
}