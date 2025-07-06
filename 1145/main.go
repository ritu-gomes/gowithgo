package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	start := 1

	for start < y {
		for i := start; i < x + start; i++ {
			if i < x + start - 1 {
				fmt.Printf("%d ", i)
			}else{
				fmt.Printf("%d\n", i)
			}	
		}		
		start += x
	}
}