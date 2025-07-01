package main

import (
	"fmt"
)

func main() {
	var password int
	for {
		fmt.Scan(&password)
		if password == 2002 {
			fmt.Println("Acesso Permitido")
			break
		}else{
			fmt.Println("Senha Invalida")
		}
	}
}