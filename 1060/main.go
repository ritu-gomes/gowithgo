package main

import (
	"fmt"
)

func main() {
	var number1, number2, number3, number4, number5, number6 float64
	fmt.Scan(&number1, &number2, &number3, &number4, &number5, &number6)
	positiveCount := 0

	numbers := []float64{number1, number2, number3, number4, number5, number6}
	for _, value := range numbers{
		if value > 0 {
			positiveCount = positiveCount + 1
		}
	}

	fmt.Println(positiveCount, "valores positivos")
}