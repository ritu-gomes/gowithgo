package main

import (
	"fmt"
)

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)
		if x < y {
			fmt.Println("Crescente")
		}else if x > y {
			fmt.Println("Decrescente")
		}else{
			break
		}
	}
}