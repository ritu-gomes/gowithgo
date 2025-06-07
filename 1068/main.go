package main

import (
	"fmt"
)

func main() {
	var start int
	fmt.Scan(&start)

	if start % 2 == 0 {
		start = start + 1
	}

	for i := 0; i < 6; i++{
		fmt.Println(start)
		start = start + 2
	}
}