package main

import (
	"fmt"
)

func main() {
	var code int
	fmt.Scan(&code)

	message := " "

	switch code{
	case 61:
		message = "Brasilia"
	case 71:
		message = "Salvador"
	case 11:
		message = "Sao Paulo"
	case 21:
		message = "Rio de Janeiro"
	case 32:
		message = "Juiz de Fora"
	case 19:
		message = "Campinas"
	case 27:
		message = "Vitoria"
	case 31:
		message = "Belo Horizonte"
	default:
		message = "DDD nao cadastrado"
	}

	fmt.Println(message)
}