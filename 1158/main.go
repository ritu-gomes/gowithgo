package main

import (
	"fmt"
)

func main() {
	var n, x, y int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		
		if x % 2 == 0 {
			x += 1
		}

		sum := x

		for i := 1; i < y; i++ {
			x += 2
			sum += x
		}
		 
		fmt.Println(sum)
	}
}