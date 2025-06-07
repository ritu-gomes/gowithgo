package main

import (
	"fmt"
	"math"
)

func main() {
	var number1, number2, number3, number4, number5 float64
	fmt.Scan(&number1, &number2, &number3, &number4, &number5)

	evens := 0
	odds := 0
	positives := 0
	negetives := 0

	allNumbers := []float64{number1, number2, number3, number4, number5}

	for _, value := range allNumbers{
		if math.Mod(value, 2) == 0 {
			evens = evens +1
		}else{
			odds = odds + 1
		}

		if value > 0 {
			positives = positives + 1
		}else if value < 0 {
			negetives = negetives + 1
		}
	}

	fmt.Println(evens, "valor(es) par(es)")
	fmt.Println(odds, "valor(es) impar(es)")
	fmt.Println(positives, "valor(es) positivo(s)")
	fmt.Println(negetives, "valor(es) negativo(s)")

}