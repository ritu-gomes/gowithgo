package main

import (
	"fmt"
)

func main() {
	var n, x, y int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		if y == 0 {
			fmt.Println("divisao impossivel")
		}else {
			result := float64(x) / float64(y)
			fmt.Printf("%.1f\n", result)
		}
	}
}