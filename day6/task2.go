package main

import (
	"fmt"
	"strconv"
	"strings"
)

const Days = 256

var memory [Days]int64

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

func breed(n int) int64 {
	if n <= 0 {
		return 1
	}
	if memory[n] != 0 {
		return memory[n]
	}
	result := breed(n-7) + breed(n-9)
	memory[n] = result
	return result
}

func main() {
	fishes := readNumbers()

	var cnt int64 = 0
	for _, fish := range fishes {
		cnt += breed(Days - fish)
	}

	fmt.Println(cnt)
}
