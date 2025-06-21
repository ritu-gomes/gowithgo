package main

import (
	"fmt"
)

func main() {
	var n, num int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&num)

		if num == 0 {
			fmt.Println("NULL")
		}else if num > 0 {
			if num % 2 == 0 {
				fmt.Println("EVEN POSITIVE")
			}else{
				fmt.Println("ODD POSITIVE")
			}
		}else{
			if num % 2 == 0 {
				fmt.Println("EVEN NEGATIVE")
			}else{
				fmt.Println("ODD NEGATIVE")
			}
		}
	}	
} 