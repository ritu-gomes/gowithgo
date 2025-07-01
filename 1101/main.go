package main

import (
	"fmt"
)

func main() {
	for {
		var x, y int
		sum := 0
		fmt.Scan(&x, &y)
		if x <= 0 || y <= 0 {
			break
		}else{
			if x > y {
				a := x
				x = y
				y = a
			}
			for i := x; i <= y; i++ {
				fmt.Printf("%d ", i)
				sum += i
			}
			fmt.Printf("Sum=%d\n",sum)
		}
	}
}