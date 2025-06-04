package main

import (
	"fmt"
)

func main() {
	var start, end int
	fmt.Scan(&start, &end)
	duration := 0

	if start < end {
		duration = end - start
	}else if start > end {
		duration = (24 - start) + end 
	}else {
		duration = 24
	}

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", duration)
}