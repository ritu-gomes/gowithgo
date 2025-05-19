package main

import (
	"fmt"
)
 
func main() {

    var number, hours int
	var perHour float64

	fmt.Scanf("%d\n%d\n%f", &number, &hours, &perHour)

	salary := float64(hours) * perHour

	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f", number, salary)
	

}
