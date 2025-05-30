package main

import (
	"fmt"
)
 
func main() {

    const Pi = 3.14159
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	areaOfTriangle := (1/2.0) * A * C

	areaOfCircle := Pi * C * C

	areaOfTrapezium := (1/2.0) * (A + B) * C

	areaOfSquare := B * B

	areaOfRectangle := A * B

	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n", areaOfTriangle, areaOfCircle, areaOfTrapezium, areaOfSquare, areaOfRectangle)

	// 12.7 10.4 15.2
// 	TRIANGULO: 96.520
// CIRCULO: 725.833
// TRAPEZIO: 175.560
// QUADRADO: 108.160
// RETANGULO: 132.080

}
