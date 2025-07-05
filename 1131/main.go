package main

import (
	"fmt"
)

func main() {
	var x, y int
	count := 0
	inter := 0
	gremio := 0
	draws := 0

	for {
		fmt.Scan(&x, &y)
		if x > y {
			inter += 1
		}else if x < y {
			gremio += 1
		}else{
			draws += 1
		}
		count += 1
		fmt.Println("Novo grenal (1-sim 2-nao)")
		var reEntry int
		fmt.Scan(&reEntry)
		if reEntry == 1 {
			continue
		}else if reEntry == 2 {
			break
		}
	}

	fmt.Println(count ,"grenais")
	fmt.Printf("Inter:%d\n", inter)
	fmt.Printf("Gremio:%d\n", gremio)
	fmt.Printf("Empates:%d\n", draws)
	if inter > gremio {
		fmt.Println("Inter venceu mais")
	}else if inter < gremio {
		fmt.Println("Gremio venceu mais")
	}else {
		fmt.Println("Não houve vencedor")
	}
}