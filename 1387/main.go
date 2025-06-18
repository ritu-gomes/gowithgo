package main

import (
	"fmt"
)

func main() {
	var sons, daughters int

	for {
		fmt.Scanf("%d %d\n", &sons, &daughters)
		if sons == 0 && daughters == 0 {
			break
		}else{
			fmt.Println(sons + daughters)
		}
	}

}