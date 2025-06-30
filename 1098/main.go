package main

import (
	"fmt"
)

func main() {
	var i float64
	var value2 float64 = 1

	for i = 0; i <= 2; i += 0.2 {
		for j := value2; j < value2 + 3; j++ {
			fmt.Printf("I=%.2g J=%.2g\n", i, j)
		}
		value2 += 0.2
	}
}