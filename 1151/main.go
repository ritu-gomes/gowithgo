package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fibonacci := []int{0,1}
	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	for i := 0; i < len(fibonacci); i++ {
		if i < len(fibonacci) - 1 {
			fmt.Printf("%d ", fibonacci[i])
		}else{
			fmt.Printf("%d\n", fibonacci[i])
		}
	}
}