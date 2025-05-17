package main

import (
	"fmt"
)
 
func main() {

    var A, B float64
	fmt.Scanf("%f\n%f",&A, &B)
	avrg := (A * 3.5 + B * 7.5) / 11.0
	fmt.Printf("MEDIA = %.5f\n",avrg)

}
