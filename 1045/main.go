package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	lines := []float64{a, b, c}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i] > lines[j]
	})

	a = lines[0]
	b = lines[1]
	c = lines[2]

	if a >= (b + c) {
		fmt.Println("NAO FORMA TRIANGULO")
		return
	}
	if (a * a) == (b * b) + (c * c) {
		fmt.Println("TRIANGULO RETANGULO")
	}
	if (a * a) > (b * b) + (c * c) {
		fmt.Println("TRIANGULO OBTUSANGULO")
	}
	if (a * a) < (b * b) + (c * c) {
		fmt.Println("TRIANGULO ACUTANGULO")
	}
	if a == b && b == c {
		fmt.Println("TRIANGULO EQUILATERO")
	}
	if (a == b && b != c && a != c) ||
   (b == c && b != a && a != c) ||
   (a == c && a != b && b != c) {
		fmt.Println("TRIANGULO ISOSCELES")
	}

}