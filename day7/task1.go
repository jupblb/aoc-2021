package main

import (
	"fmt"
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

func median(arr []int) int {
	return (arr[len(arr)/2-1] + arr[len(arr)/2]) / 2
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	positions := readNumbers()
	medPosition := median(positions)

	res := 0
	for _, position := range positions {
		res += abs(medPosition - position)
	}

	fmt.Println(res)
}
