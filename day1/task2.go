package main

import (
	"fmt"
)

func main() {
	var buffer [3]int
	fmt.Scanf("%d", &buffer[0])
	fmt.Scanf("%d", &buffer[1])
	fmt.Scanf("%d", &buffer[2])

	cnt := 0
	for i := 0; true; i++ {
		var input int
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			break
		}

		previous := buffer[0] + buffer[1] + buffer[2]
		buffer[i%3] = input
		current := buffer[0] + buffer[1] + buffer[2]

		if current > previous {
			cnt++
		}
	}

	fmt.Println(cnt)
}
