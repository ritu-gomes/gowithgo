package main

import (
	"fmt"
)

func blackOrWhite(values []int) (int, int){
	blackCount := 0
	position := 0
	for i := 0; i < 5; i++ {
		if values[i] <= 127 {
		blackCount += 1
		position = i
		}
	}
	return blackCount, position
}

func main() {
	var n int
	for {
		fmt.Scan(&n)

		if n == 0 {
			break
		}else{
			for i := 0; i < n; i++ {
				var A, B, C, D, E int
				fmt.Scan(&A, &B, &C, &D, &E)
				options := []int{A, B, C, D, E}
				optionMap := map[int]string{0: "A", 1: "B", 2: "C", 3: "D", 4: "E"}
				black, position := blackOrWhite(options)
				
				if black == 1 {
					fmt.Println(optionMap[position])
				}else{
					fmt.Println("*")
				}
			}
		}
	}
}