package main

import (
	"fmt"
	"math"
)

func main() {
	var ageInDays int
	fmt.Scan(&ageInDays)

	years := int(math.Floor(float64(ageInDays) / 365))
	ageInDays = ageInDays % 365

	months := int(math.Floor(float64(ageInDays) / 30))
	ageInDays = ageInDays % 30

	days := int(ageInDays)

	fmt.Println(years,"ano(s)")
	fmt.Println(months,"mes(es)")
	fmt.Println(days,"dia(s)")
}