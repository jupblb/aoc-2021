package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readNumbers() []int {
	var firstLine string
	fmt.Scanf("%s", &firstLine)

	splitFirstLine := strings.Split(firstLine, ",")
	numbers := make([]int, len(splitFirstLine))
	for i, number := range splitFirstLine {
		numbers[i], _ = strconv.Atoi(number)
	}

	return numbers
}

func main() {
	fishes := readNumbers()

	for day := 0; day < 80; day++ {
		var newFishes []int
		for _, fish := range fishes {
			if fish == 0 {
				newFishes = append(newFishes, 6, 8)
			} else {
				newFishes = append(newFishes, fish-1)
			}
		}
		fishes = newFishes
	}

	fmt.Println(len(fishes))
}
