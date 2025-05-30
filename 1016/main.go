package main

import "fmt"

func main() {
	var distance int
	fmt.Scanln(&distance)

	time := distance * 2
	fmt.Println(time, "minutos")
}