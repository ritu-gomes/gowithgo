package main

import (
	"fmt"
)

func main() {
	var n int
	indicator := 0
	numbers := []int{}
	numberCount := 0
	sum := 0
	result := 0
	for {
		fmt.Scan(&n)
		if n > indicator {
			numbers = append(numbers, n)
			numberCount ++
			indicator = n
		}
		if numberCount == 2{
			break
		}
	}
	for sum < numbers[1] {
		sum += numbers[0]
		numbers[0] += 1
		result ++
	}
	fmt.Println(result)
}