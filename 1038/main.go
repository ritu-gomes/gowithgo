package main

import (
	"fmt"
)

func main() {
	var xCode, yQuantity int
	fmt.Scan(&xCode, &yQuantity)

	switch xCode{
	case 1:
		price := 4.00 * float64(yQuantity)
		fmt.Printf("Total: R$ %.2f\n", price)
	case 2:
		price := 4.50 * float64(yQuantity)
		fmt.Printf("Total: R$ %.2f\n", price)
	case 3:
		price := 5.00 * float64(yQuantity)
		fmt.Printf("Total: R$ %.2f\n", price)
	case 4:
		price := 2.00 * float64(yQuantity)
		fmt.Printf("Total: R$ %.2f\n", price)
	case 5:
		price := 1.50 * float64(yQuantity)
		fmt.Printf("Total: R$ %.2f\n", price)

	}
}