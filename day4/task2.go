package main

import (
	"fmt"
	"strconv"
	"strings"
)

const BoardSize = 5

type BoardCell struct {
	value    int
	isMarked bool
}

func readNumbers() []int {
	var firstLine string
	fmt.Scanf("%s\n\n", &firstLine)

	splitFirstLine := strings.Split(firstLine, ",")
	numbers := make([]int, len(splitFirstLine))
	for i, number := range splitFirstLine {
		numbers[i], _ = strconv.Atoi(number)
	}

	return numbers
}

func readBoard() ([BoardSize][BoardSize]BoardCell, bool) {
	var board [BoardSize][BoardSize]BoardCell

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			fmt.Scanf("%d", &board[i][j].value)
		}
	}

	var newline rune
	_, err := fmt.Scanf("%c", &newline)

	return board, err == nil
}

func readBoards() [][BoardSize][BoardSize]BoardCell {
	var boards [][BoardSize][BoardSize]BoardCell

	for {
		board, hasNext := readBoard()
		boards = append(boards, board)
		if !hasNext {
			break
		}
	}

	return boards
}

func mark(number int, boards *[][BoardSize][BoardSize]BoardCell) {
	for b := 0; b < len(*boards); b++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if (*boards)[b][i][j].value == number {
					(*boards)[b][i][j].isMarked = true
				}
			}
		}
	}
}

func countUnmarked(board [BoardSize][BoardSize]BoardCell) int {
	cnt := 0

	for _, row := range board {
		for _, cell := range row {
			if !cell.isMarked {
				cnt += cell.value
			}
		}
	}

	return cnt
}

func findWinning(boards [][BoardSize][BoardSize]BoardCell, hasWon []bool) (int, int, bool) {
	for b, board := range boards {
		if !hasWon[b] {
			for i := 0; i < BoardSize; i++ {
				isRowBingo := true
				isColumnBingo := true

				for j := 0; j < BoardSize; j++ {
					if !board[i][j].isMarked {
						isRowBingo = false
					}
					if !board[j][i].isMarked {
						isColumnBingo = false
					}
				}

				if isRowBingo || isColumnBingo {
					return countUnmarked(board), b, true
				}
			}
		}
	}

	return 0, -1, false
}

func main() {
	numbers := readNumbers()
	boards := readBoards()
	hasWon := make([]bool, len(boards))
	lastWin := 0

	for _, number := range numbers {
		mark(number, &boards)

		for {
			cnt, board, isFound := findWinning(boards, hasWon)
			if isFound {
				hasWon[board] = true
				lastWin = cnt * number
			} else {
				break
			}
		}
	}

	fmt.Println(lastWin)
}
