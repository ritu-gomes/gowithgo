package main

import (
	"fmt"
)

func main() {
	const Pi = 3.14159
	var radius float64
	fmt.Scan(&radius)

	volume := (4.0/3.0) * Pi * radius * radius * radius

	fmt.Printf("VOLUME = %.3f\n", volume)
}