package main

import (
	"fmt"
)

func main()  {
	var month int
	fmt.Scan(&month)
	name := " "

	switch month {
	case 1:
		name = "January"
	case 2:
		name = "February"
	case 3:
		name = "March"
	case 4:
		name = "April"
	case 5:
		name = "May"
	case 6:
		name = "June"
	case 7:
		name = "July"
	case 8:
		name = "August"
	case 9:
		name = "September"
	case 10:
		name = "October"
	case 11:
		name = "November"
	case 12:
		name = "December"
	}

	fmt.Println(name)
	
}