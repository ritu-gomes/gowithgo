package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 2; i <= 10000; i += n {
		fmt.Println(i)
	}
}