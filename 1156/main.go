package main

import (
	"fmt"
)

func main()  {
	s := 1.0
	divisor := 2
	for i := 3.0; i <= 39.0; i += 2 {
		s += i / float64(divisor)
		divisor *= 2
	}
	fmt.Printf("%.2f\n", s)
}