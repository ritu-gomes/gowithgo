package main

import (
	"fmt"
)

func main() {
	var num int
	highest := 0
	position := 0

	for i := 1; i <= 100; i++ {
		fmt.Scan(&num)
		if highest < num {
			highest = num
			position = i
		}
	}

	fmt.Println(highest)
	fmt.Println(position)
}