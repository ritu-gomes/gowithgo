package main

import (
	"fmt"
)

func main() {
	var startHour, startMinut, endHour, endMinut int
	fmt.Scan(&startHour, &startMinut, &endHour, &endMinut)
	hour := 0
	minut := 0

	if startHour < endHour {
		if startMinut < endMinut {
			hour = endHour - startHour
			minut = endMinut - startMinut
		}
		if startMinut > endMinut {
			hour = (endHour - startHour) - 1
			minut = (60 - startMinut) + endMinut
		}
	}else if startHour > endHour {
		if startMinut < endMinut {
			hour = (24 - startHour) + endHour
			minut = endHour - startHour
		}
		if startMinut > endMinut {
			hour = ((24 - startHour) + endHour) - 1
			minut = (60 - startMinut) + endMinut
		}
	}else {
		if startMinut < endMinut {
			minut = endMinut - startMinut
		}else if startMinut > endMinut {
			hour = 23
			minut = (60 - startMinut) + endMinut
		}else {
			hour = 24
		}
	}

	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", hour, minut)
}