package main

import (
	"fmt"
)

func main(){
	const pi float64 = 3.14159
	var R float64
	fmt.Scanln(&R)
	area := pi * R * R
	fmt.Printf("A=%.4f\n",area)
	
}