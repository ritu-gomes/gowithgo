package main

import (
	"fmt"
)

func main() {
	var end, start int
	fmt.Scan(&end, &start)
	sum := 0
	// odds := []int{}

	if start % 2 == 0 {
		start = start + 1
	}else{
		start = start + 2
	}

	for start < end{
		sum = sum + start
		start = start + 2
	}
	
	
	fmt.Println(sum)
}