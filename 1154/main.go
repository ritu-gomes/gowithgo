package main

import (
	"fmt"
)

func main() {
	var n int
	sum := 0
	count := 0
	for {
		fmt.Scan(&n)
		if n < 0 {
			break
		}else{
			sum += n
			count += 1
		}
	}
	average := float64(sum) / float64(count)
	fmt.Printf("%.2f\n", average)
}