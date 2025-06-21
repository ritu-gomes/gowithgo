package main

import (
	"fmt"
)

func main() {
	var n int
	var number1, number2, number3 float64
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&number1, &number2, &number3)

		average := ((number1 * 2) + (number2 * 3) + (number3 * 5)) / 10.0
		fmt.Printf("%.1f\n", average)
	}
}