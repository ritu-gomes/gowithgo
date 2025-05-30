package main

import (
	"fmt"
) 

func main() {
	var distance int
	var fuel float64
	fmt.Scanln(&distance)
	fmt.Scanln(&fuel)

	averageConsumption := float64(distance) / fuel
	fmt.Printf("%.3f km/l\n", averageConsumption)
}