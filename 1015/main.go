package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scanf("%f%f\n%f%f", &x1, &y1, &x2, &y2)

	distance := math.Sqrt((x2-x1) * (x2-x1) + (y2 - y1) * (y2 - y1))

	fmt.Printf("%.4f\n", distance)
}