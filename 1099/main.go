package main

import (
	"fmt"
)

func main() {
	var n, x, y int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		sum := 0
		fmt.Scan(&x, &y)
		if x > y {
			a := x
			x = y
			y = a
		}
		if x % 2 == 0 {
			x += 1
		}else{
			x += 2
		}
		for i := x; i < y; i += 2 {
			sum += i
		}
		fmt.Println(sum)
	}
}