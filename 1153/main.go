package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	factorial := 1
	for i := n; i >= 1; i-- {
		factorial *= i
	}
	fmt.Println(factorial)
}