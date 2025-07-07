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
			if n % 2 != 0 {
				n ++
			}
			sum := n

			for i := 1; i < 5; i++ {
				n += 2
				sum += n
			}

			fmt.Println(sum)

		}
	}
}