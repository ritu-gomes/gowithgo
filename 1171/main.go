package main

import (
	"fmt"
	"sort"
)

func main() {
	var limit int
	var num int
	numbers := []int{}
	numberMap := make(map[int]int)

	fmt.Scanln(&limit)
	for i := 0; i < limit; i++ {
		fmt.Scanln(&num)
		numberMap[num] = numberMap[num] + 1
	}

	for key := range numberMap {
		numbers = append(numbers, key)
	}

	sort.Ints(numbers)

	for _, value := range numbers {
		fmt.Printf("%d aparece %d vez(es)\n", value, numberMap[value] )
	}
	
}