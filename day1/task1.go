package main

import (
	"fmt"
	"math"
)

func main() {
	cnt := 0
	previous := math.MaxInt32
	for {
		var input int
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			break
		}
		if input > previous {
			cnt++
		}
		previous = input
	}

	fmt.Println(cnt)
}
