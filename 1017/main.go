package main

import (
	"fmt"
)

func main() {
	var totalTime, speed int
	fmt.Scan(&totalTime, &speed)

	totalFuel := float64(totalTime*speed) / 12 //as car does 12 km/l

	fmt.Printf("%.3f\n", totalFuel)
}