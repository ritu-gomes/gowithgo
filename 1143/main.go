package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= 3; j++ {
			result := int(math.Pow(float64(i), float64(j)))
			if j < 3 {
				fmt.Printf("%d ", result)
			}else{
				fmt.Printf("%d\n", result)
			}
		}
	}
}