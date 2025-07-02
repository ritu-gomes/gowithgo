package main

import (
	"fmt"
)

func main() {
	var score float64
	var x int
	scoreArray := []float64{}
	count := 0

	for {
		fmt.Scan(&score)
		if score < 0 || score > 10 {
			fmt.Println("nota invalida")
		}else{
			scoreArray = append(scoreArray, score)
			count ++
			if count == 2 {
				average := (scoreArray[0] + scoreArray[1]) / 2
				fmt.Printf("media = %.2f\n", average)

				for {
					fmt.Println("novo calculo (1-sim 2-nao)")
					fmt.Scan(&x)
					if x == 2 {
						return
					}else if x == 1 {
						scoreArray = nil
						count = 0
						break
					}
				}	
			}
		}
	}

}