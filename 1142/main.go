package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	start := 1

	for i := 0; i < n; i++ {
		for i := start; i < start + 3; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println("PUM")
		start += 4
	}
}