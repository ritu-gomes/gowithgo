package main

import (
	"fmt"
	"math"
)

func main() {
	var value int
	fmt.Scanln(&value)

	hours := int(math.Floor(float64(value) / 3600))
	value = value % 3600

	minuts := int(math.Floor(float64(value) / 60))
	value = value % 60

	seconds := int(value)

	fmt.Printf("%d:%d:%d\n", hours, minuts, seconds)
}