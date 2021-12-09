package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input [][]int

	for {
		var line string
		if _, err := fmt.Scanf("%s", &line); err != nil {
			break
		}

		numbers := make([]int, len(line))
		for i, c := range line {
			numbers[i], _ = strconv.Atoi(string(c))
		}

		input = append(input, numbers)
	}

	res := 0
	for i, row := range input {
		for j, number := range row {
			isLowest := true

			if (i > 0 && input[i-1][j] <= number) ||
				(i < len(input)-1 && input[i+1][j] <= number) ||
				(j > 0 && input[i][j-1] <= number) ||
				(j < len(row)-1 && input[i][j+1] <= number) {
				isLowest = false
			}

			if isLowest {
				res += 1 + number
			}
		}
	}
	fmt.Println(res)
}
