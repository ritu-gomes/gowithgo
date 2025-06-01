package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	delta := (b * b) - (4 * a * c)

	if delta >= 0 && a!= 0 {
		x1 := (- b + math.Sqrt(delta)) / (2 * a)
		x2 := (- b - math.Sqrt(delta)) / (2 * a)
		
		fmt.Printf("R1 = %.5f\n", x1)
		fmt.Printf("R2 = %.5f\n", x2)
	}else{
		fmt.Println("Impossivel calcular")
	}
} 