package main

import (
	"fmt"
)

func main() {
	var n int
	inputs := []int{}
	sum := 0
	for len(inputs) < 2 {
		fmt.Scan(&n)
		if n > 0 {
			inputs = append(inputs, n)
		}else{
			continue
		}
	}
	for i := inputs[0]; i < inputs[0] + inputs[1]; i++ {
		sum += i
	}
	fmt.Println(sum)
}