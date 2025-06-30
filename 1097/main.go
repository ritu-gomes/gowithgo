package main

import (
	"fmt"
)

func main() {
	secontLoop := 7

	for i := 1; i < 10; i += 2 {
		for j := secontLoop; j > secontLoop - 3; j-- {
			fmt.Printf("I=%d J=%d\n", i, j)
		}
		secontLoop += 2
	}
}