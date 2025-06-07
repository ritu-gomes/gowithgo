package main

import (
	"fmt"
)

func main() {
	var limit int
	fmt.Scan(&limit)
	i := 1

	for i <= limit{
		fmt.Println(i)
		i = i + 2
	}
}