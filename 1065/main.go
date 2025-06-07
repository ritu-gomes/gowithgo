package main

import (
	"fmt"
	"math"
)

func main() {
	var number1, number2, number3, number4, number5 float64
	fmt.Scan(&number1, &number2, &number3, &number4, &number5)
	evens := 0

	allNumbers := []float64{number1, number2, number3, number4, number5}

	for _, value := range allNumbers {
		if math.Mod(value, 2.0) == 0 {
			evens = evens + 1
		}
	}

	fmt.Println(evens, "valores pares")

}