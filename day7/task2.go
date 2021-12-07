package main

import (
	"fmt"
	"math"
	"sort"
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
	sort.Ints(numbers)

	return numbers
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	positions := readNumbers()
	minRes := math.MaxInt32

	for mid := 0; mid <= positions[len(positions)-1]; mid++ {
		res := 0

		for _, position := range positions {
			dist := abs(mid - position)
			res += (dist * (dist + 1)) / 2
		}

		if res < minRes {
			minRes = res
		}
	}

	fmt.Println(minRes)
}
