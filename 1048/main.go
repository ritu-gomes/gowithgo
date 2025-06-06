package main

import (
	"fmt"
)

func percentage(salary float64, rate int)(float64, float64){
	increment := (salary * float64(rate)) / 100.0
	newSalary := salary + increment

	return newSalary, increment
}

func main() {
	var salary float64
	fmt.Scan(&salary)

	newSalary := 0.0
	increment := 0.0
	rate := 0

	if salary >= 0 && salary <= 400.0 {
		rate = 15
		newSalary, increment = percentage(salary,rate)
	}
	if salary >= 400.01 && salary <= 800.0 {
		rate = 12
		newSalary, increment = percentage(salary,rate)
	}
	if salary >= 800.01 && salary <= 1200.0 {
		rate = 10
		newSalary, increment = percentage(salary,rate)
	}
	if salary >= 1200.01 && salary <= 2000.0 {
		rate = 7
		newSalary, increment = percentage(salary,rate)
	}
	if salary > 2000 {
		rate = 4
		newSalary, increment = percentage(salary,rate)
	}

	fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %d %%\n", newSalary, increment, rate)
}