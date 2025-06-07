package main

import (
	"fmt"
)

func main() {
	var number1, number2, number3, number4, number5, number6 float64
	fmt.Scan(&number1, &number2, &number3, &number4, &number5, &number6)

	allNumbers := []float64{number1, number2, number3, number4, number5, number6}
	positiveSlice := []float64{}

	for _, value := range allNumbers {
		if value > 0 {
			positiveSlice = append(positiveSlice, value)
		}
	}

	sum := 0.0

	for _, pos := range positiveSlice{
		sum = sum + pos
	}

	average := sum / float64(len(positiveSlice))

	fmt.Println(len(positiveSlice),"valores positivos")
	fmt.Printf("%.1f\n", average)
}