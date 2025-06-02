package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3, n4 float64
	fmt.Scan(&n1, &n2, &n3, &n4)

	avarage := ((n1*2) + (n2*3) + (n3*4) + (n4*1)) / 10.00
	fmt.Printf("Media: %.1f\n", avarage)

	if avarage >= 7.0 {
		fmt.Println("Aluno aprovado.")
	}else if avarage < 5.0 {
		fmt.Println("Aluno reprovado.")
	}else {
		fmt.Println("Aluno em exame.")
		var score float64
		fmt.Scan(&score)

		fmt.Printf("Nota do exame: %.1f\n", score)
		newAverage := (avarage + score) / 2.0

		if newAverage >= 5.0 {
			fmt.Println("Aluno aprovado.")
		}else{
			fmt.Println("Aluno reprovado.")
		}

		fmt.Printf("Media final: %.1f\n", newAverage)
	}
}