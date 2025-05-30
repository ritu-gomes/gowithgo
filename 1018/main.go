package main

import (
	"fmt"
	"math"
)

func main() {
	var value int
	fmt.Scanln(&value)
	var remainder int = 0

	hundreds := int(math.Floor(float64(value)/100))
	remainder = value % 100

	fiftys := int(math.Floor(float64(remainder)/50))
	remainder = remainder % 50

	twentys := int(math.Floor(float64(remainder)/20))
	remainder = remainder % 20

	tens := int(math.Floor(float64(remainder)/10))
	remainder = remainder % 10

	fives := int(math.Floor(float64(remainder)/5))
	remainder = remainder % 5

	twos := int(math.Floor(float64(remainder)/2))
	remainder = remainder % 2

	ones := int(remainder)

	fmt.Println(value)
	fmt.Printf("%d nota(s) de R$ 100,00\n", hundreds)
	fmt.Printf("%d nota(s) de R$ 50,00\n", fiftys)
	fmt.Printf("%d nota(s) de R$ 20,00\n", twentys)
	fmt.Printf("%d nota(s) de R$ 10,00\n", tens)
	fmt.Printf("%d nota(s) de R$ 5,00\n", fives)
	fmt.Printf("%d nota(s) de R$ 2,00\n", twos)
	fmt.Printf("%d nota(s) de R$ 1,00\n", ones)
}