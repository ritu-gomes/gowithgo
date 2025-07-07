package main

import (
	"fmt"
)

func main() {
	s := 0.0
	for i := 1; i <= 100; i++ {
		s += (1.0 / float64(i))
	}

	fmt.Printf("%.2f\n", s)
}