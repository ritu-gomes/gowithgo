package main

import (
	"fmt"
)

func animalPercentage(total int, animal int) float64{
	percentage := (float64(animal) * 100.0) / float64(total)
	return percentage
}

func main() {
	var testCases, animalCount int
	fmt.Scan(&testCases)
	var animal string
	total := 0
	rabbits := 0
	rats := 0
	frogs := 0

	for i := 0; i < testCases; i++ {
		fmt.Scan(&animalCount, &animal)
		total = total + animalCount
		if animal == "C" {
			rabbits = rabbits + animalCount
		}else if animal == "R" {
			rats = rats + animalCount
		}else if animal == "S" {
			frogs = frogs + animalCount
		}
	}

	rabbitPercentage := animalPercentage(total, rabbits)
	ratPercentage := animalPercentage(total, rats)
	frogPercentage := animalPercentage(total, frogs)

	fmt.Printf("Total: %d cobaias\n", total)
	fmt.Printf("Total de coelhos: %d\n", rabbits)
	fmt.Printf("Total de ratos: %d\n", rats)
	fmt.Printf("Total de sapos: %d\n", frogs)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", rabbitPercentage)
	fmt.Printf("Percentual de ratos: %.2f %%\n", ratPercentage)
	fmt.Printf("Percentual de sapos: %.2f %%\n", frogPercentage)
}