package main

import (
	"fmt"
	"math"
)

func greatest(x int, y int) int {
	var great = (x + y + int(math.Abs(float64(x) - float64(y)))) / 2
	return great
}

func main () {
	var firstNum, secondNum, thirdNum int
	fmt.Scan(&firstNum, &secondNum, &thirdNum)
	max1 := greatest(firstNum, secondNum)
	max := greatest(max1, thirdNum)

	fmt.Println(max,"eh o maior")
}