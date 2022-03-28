package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readNumbers() ([]int, bool) {
	var line string
	if _, err := fmt.Scanf("%s", &line); err != nil {
		return nil, false
	}

	splitFirstLine := strings.Split(line, "")
	numbers := make([]int, len(splitFirstLine))
	for i, number := range splitFirstLine {
		numbers[i], _ = strconv.Atoi(number)
	}

	return numbers, true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	var cave [][]int

	for {
		if row, ok := readNumbers(); ok {
			cave = append(cave, row)
		} else {
			break
		}
	}

	for i := 1; i < len(cave); i++ {
		cave[i][0] += cave[i-1][0]
		cave[0][i] += cave[0][i-1]
	}

	for i := 1; i < len(cave); i++ {
		for j := 1; j < len(cave); j++ {
			cave[i][j] += min(cave[i-1][j], cave[i][j-1])
		}
	}

	fmt.Println(cave[len(cave)-1][len(cave)-1] - cave[0][0])
}
