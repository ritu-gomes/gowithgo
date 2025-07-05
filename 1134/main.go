package main

import (
	"fmt"
)

func main() {
	var n int
	alcool := 0
	gasolina := 0
	diesel := 0

	fuelLoop: for {
		fmt.Scan(&n)
		switch n {
		case 1:
			alcool += 1
		case 2:
			gasolina += 1
		case 3:
			diesel += 1
		case 4:
			break fuelLoop
		default:
			continue
		}
	}
	
	fmt.Println("MUITO OBRIGADO")
	fmt.Println("Alcool:", alcool)
	fmt.Println("Gasolina:", gasolina)
	fmt.Println("Diesel:", diesel)
}