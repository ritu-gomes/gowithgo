package main

import (
	"fmt"
)

func location(dPointx int, dPointy int, x int, y int) string {
	result := " "

	if x == dPointx || y == dPointy {
		result = "divisa"
	}else if x > dPointx && y > dPointy {
		result = "NE"
	}else if x < dPointx && y > dPointy {
		result = "NO"
	}else if x < dPointx && y < dPointy {
		result = "SO"
	}else if x > dPointx && y < dPointy {
		result = "SE"
	}

	return result
}

func main() {
	var k int
	for {
		fmt.Scanln(&k)
		if k == 0 {
			break
		}else{
			var m,n int
			fmt.Scanf("%d %d\n", &m, &n)
			for i := 0; i < k; i++ {
				var x, y int
				fmt.Scanf("%d %d\n", &x, &y)
				homeIn := location(m,n,x,y)
				fmt.Println(homeIn)
			}
		}
	}
}