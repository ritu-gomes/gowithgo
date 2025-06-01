package main

import (
	"fmt"
)

func main() {
	var value float64
	fmt.Scan(&value)

	// Convert to cents
	totalCents := int(value*100 + 0.5)

	// Notes
	hundreds := totalCents / 10000
	totalCents %= 10000

	fiftys := totalCents / 5000
	totalCents %= 5000

	twentys := totalCents / 2000
	totalCents %= 2000

	tens := totalCents / 1000
	totalCents %= 1000

	fives := totalCents / 500
	totalCents %= 500

	twos := totalCents / 200
	totalCents %= 200

	// Coins
	ones := totalCents / 100
	totalCents %= 100

	pointFiftys := totalCents / 50
	totalCents %= 50

	pointTwentyFives := totalCents / 25
	totalCents %= 25

	pointTens := totalCents / 10
	totalCents %= 10

	pointFives := totalCents / 5
	totalCents %= 5

	pointOnes := totalCents

	// Output
	fmt.Println("NOTAS:")
	fmt.Printf("%d nota(s) de R$ 100.00\n", hundreds)
	fmt.Printf("%d nota(s) de R$ 50.00\n", fiftys)
	fmt.Printf("%d nota(s) de R$ 20.00\n", twentys)
	fmt.Printf("%d nota(s) de R$ 10.00\n", tens)
	fmt.Printf("%d nota(s) de R$ 5.00\n", fives)
	fmt.Printf("%d nota(s) de R$ 2.00\n", twos)
	fmt.Println("MOEDAS:")
	fmt.Printf("%d moeda(s) de R$ 1.00\n", ones)
	fmt.Printf("%d moeda(s) de R$ 0.50\n", pointFiftys)
	fmt.Printf("%d moeda(s) de R$ 0.25\n", pointTwentyFives)
	fmt.Printf("%d moeda(s) de R$ 0.10\n", pointTens)
	fmt.Printf("%d moeda(s) de R$ 0.05\n", pointFives)
	fmt.Printf("%d moeda(s) de R$ 0.01\n", pointOnes)
}
