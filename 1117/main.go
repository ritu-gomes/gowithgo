package main

import (
	"fmt"
)

func main() {
	var n float64
	scores := []float64{}
	count := 0
	for count < 2 {
		fmt.Scan(&n)
		if n >= 0 && n <= 10.0 {
			scores = append(scores, n)
			count ++
		}else {
			fmt.Println("nota invalida")
		}
	}
	x := scores[0]
	y := scores[1]
	average := (x + y) / 2.0

	fmt.Printf("media = %.2f\n", average)
}