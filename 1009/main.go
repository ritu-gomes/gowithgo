package main

import "fmt" 

func main () {

	var name string
	var fixedSalary, totalSale float64

	fmt.Scanf("%s %f %f", &name, & fixedSalary, &totalSale)

	bonus := (totalSale * 15.00) / 100.00
	totalSalary := bonus + fixedSalary

	fmt.Printf("TOTAL = R$ %.2f\n", totalSalary)

}