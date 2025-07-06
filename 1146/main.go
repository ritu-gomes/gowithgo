package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		fmt.Scan(&n)

		if n == 0 {
			break
		}else{
			for i := 1; i <= n; i++ {
			if i < n {
				fmt.Printf("%d ", i)
			}else{
				fmt.Printf("%d\n", i)
			}
		}
		}
	}
}