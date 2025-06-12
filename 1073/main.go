package main

import (
	"fmt"
)

func main() {
	var limit int
	fmt.Scan(&limit)

	for i := 2; i <= limit; i += 2 {
		squared := i * i
		fmt.Printf("%d^2 = %d\n", i, squared)
	}
}