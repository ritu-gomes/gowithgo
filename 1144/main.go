package main

import (
	"fmt"
	"math"
) 

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for k := 1; k <= 3; k++ {
			result := int(math.Pow(float64(i), float64(k)))
			if k < 3 {
				fmt.Printf("%d ", result)
			}else{
				fmt.Printf("%d\n", result)
			}
			
		}
		for k := 1; k <= 3; k++ {
			result := int(math.Pow(float64(i), float64(k)))
			if k > 1 {
				result += 1
			}
			if k < 3 {
				fmt.Printf("%d ", result)
			}else{
				fmt.Printf("%d\n", result)
			}
			
		}
	}
}