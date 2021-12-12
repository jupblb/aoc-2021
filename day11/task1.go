package main

import (
	"fmt"
)

func mark(board *[10][10]int, i int, j int) {
	if i < 0 || j < 0 || i > 9 || j > 9 || (*board)[i][j] == 0 {
		return
	}
	(*board)[i][j]++
}

func main() {
	var board [10][10]int
	for i := 0; i < 10; i++ {
		var input string
		fmt.Scanf("%s", &input)
		for j := 0; j < 10; j++ {
			board[i][j] = int(input[j]-'0') % 10
		}
	}

	cnt := 0
	for step := 0; step < 100; step++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				board[i][j]++
			}
		}

		changed := true
		for changed {
			changed = false
			for i := 0; i < 10; i++ {
				for j := 0; j < 10; j++ {
					if board[i][j] >= 10 {
						cnt ++
						changed = true
						board[i][j] = 0
						mark(&board, i-1, j-1)
						mark(&board, i-1, j)
						mark(&board, i-1, j+1)
						mark(&board, i, j-1)
						mark(&board, i, j+1)
						mark(&board, i+1, j-1)
						mark(&board, i+1, j)
						mark(&board, i+1, j+1)
					}
				}
			}
		}
	}
	fmt.Println(cnt)
}
