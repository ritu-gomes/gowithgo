package main

import "fmt"

func main() {
	var startDay, startHour, startMinute, startSecond int
	var endDay, endHour, endMinute, endSecond int

	fmt.Scanf("Dia %d\n", &startDay)
	fmt.Scanf("%d : %d : %d\n", &startHour, &startMinute, &startSecond)

	fmt.Scanf("Dia %d\n", &endDay)
	fmt.Scanf("%d : %d : %d", &endHour, &endMinute, &endSecond)

	// Convert both times to total seconds
	start := startDay*86400 + startHour*3600 + startMinute*60 + startSecond
	end := endDay*86400 + endHour*3600 + endMinute*60 + endSecond

	diff := end - start

	days := diff / 86400
	diff = diff % 86400

	hours := diff / 3600
	diff = diff % 3600

	minutes := diff / 60
	seconds := diff % 60

	fmt.Printf("%d dia(s)\n", days)
	fmt.Printf("%d hora(s)\n", hours)
	fmt.Printf("%d minuto(s)\n", minutes)
	fmt.Printf("%d segundo(s)\n", seconds)
}