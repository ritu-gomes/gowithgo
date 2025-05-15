package main

import (
	"fmt"
)
 
func main() {
	var A int
	fmt.Scanln(&A)
	var B int
	fmt.Scanln(&B)
	sum := A+B
	fmt.Println("X =",sum)
}