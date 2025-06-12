package main

import (
	"fmt"
)

func main() {
	var itteration, number int
	allNumbers := []int{}
	inRange := 0
	outrange := 0

	fmt.Scan(&itteration)

	for i := 0; i < itteration; i++ {
		fmt.Scan(&number)
		allNumbers = append(allNumbers, number)
	}

	for _, value := range allNumbers{
		if value >= 10 && value <= 20 {
			inRange = inRange + 1
		}else{
			outrange = outrange + 1
		}
	}

	fmt.Println(inRange, "in")
	fmt.Println(outrange, "out")
}