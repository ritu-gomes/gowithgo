package main

import (
	"fmt"
)

func percentage(salary float64, rate int)(float64){
	increment := (salary * float64(rate)) / 100.0
	return increment
}

func main() {
	var salary float64
	fmt.Scan(&salary)
	tax := 0.0

	if salary >= 0 && salary <= 2000.00 {
		fmt.Println("Isento")
		return
	}else if salary > 2000.00 && salary <= 3000.00 {
		tax = percentage(salary - 2000.00, 8)
	}else if salary > 3000.00 && salary <= 4500.00 {
		tax = percentage(1000.00, 8) + percentage(salary - 3000.00, 18)
	}else if salary > 4500.00 {
		tax = percentage(1000.00, 8) + percentage(1500.00, 18) + percentage(salary - 4500.00, 28)
	}

	fmt.Printf("R$ %.2f\n", tax)
}