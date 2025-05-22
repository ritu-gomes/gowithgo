package main

import "fmt"

func main() {
	var code1, unit1, code2, unit2 int
	var perUnitPrice1, perUnitPrice2 float64

	fmt.Scan(&code1, &unit1, & perUnitPrice1)
	fmt.Scan(&code2, &unit2, & perUnitPrice2)

	totalPayable := (float64(unit1) * perUnitPrice1) + (float64(unit2) * perUnitPrice2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totalPayable)

	//12 1 5.30
// 16 2 5.10
// 13 2 15.30
// 161 4 5.20


}