package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	sum := 0

	if x > y {
		x, y = y, x
	}
	
	for i := x; i <= y; i++ {
		if i % 13 != 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}