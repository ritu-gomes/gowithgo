package main

import (
	"fmt"
)

func main() {
	var number float64
	fmt.Scan(&number)

	if number < 0 || number > 100 {
		fmt.Println("Fora de intervalo")
	}else {
		if number >= 0 && number <= 25.0 {
			fmt.Println("Intervalo [0,25]")
		}
		if number > 25.0 && number <= 50.0 {
			fmt.Println("Intervalo (25,50]")
		}
		if number > 50.0 && number <= 75.0 {
			fmt.Println("Intervalo (50,75]")
		}
		if number > 75.0 && number <= 100.0 {
			fmt.Println("Intervalo (75,100]")
		}
	}
}