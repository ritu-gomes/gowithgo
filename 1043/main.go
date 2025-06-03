package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a + b > c && b + c > a && a + c > b {
		perimeter := a + b + c
		fmt.Printf("Perimetro = %.1f\n", perimeter)
	}else {
		trapezium := ((a + b) * c) / 2.0
		fmt.Printf("Area = %.1f\n", trapezium)
	}
}