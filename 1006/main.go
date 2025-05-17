package main

import (
	"fmt"
)
 
func main() {

    var A, B, C float64
	fmt.Scanf("%f\n%f\n%f",&A, &B, &C)
	avrg := (A * 2 + B * 3 + C * 5) / 10.0
	fmt.Printf("MEDIA = %.1f\n",avrg)

}