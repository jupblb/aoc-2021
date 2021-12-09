package main

import (
	"fmt"
	"sort"
	"strconv"
)

func dfs(arr *[][]int, x int, y int, prev int) int {
	if x < 0 || x >= len(*arr) || y < 0 || y >= len((*arr)[x]) {
		return 0
	}

	curr := (*arr)[x][y]
	if curr < prev || curr >= 9 {
		return 0
	}
	(*arr)[x][y] = 10

	return 1 + dfs(arr, x-1, y, curr) + dfs(arr, x+1, y, curr) +
		dfs(arr, x, y-1, curr) + dfs(arr, x, y+1, curr)
}

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

	var res []int

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
				res = append(res, dfs(&input, i, j, number-1))
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	fmt.Println(res[0] * res[1] * res[2])
}
