package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Line struct {
	p1, p2 Point
}

const BoardSize = 1000

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	var lines []Line
	for {
		var x1, y1, x2, y2 int
		_, err := fmt.Scanf("%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			break
		}
		lines = append(lines, Line{Point{x1, y1}, Point{x2, y2}})
	}

	var board [BoardSize][BoardSize]int
	for _, line := range lines {
		if line.p1.x == line.p2.x {
			y1 := min(line.p1.y, line.p2.y)
			y2 := max(line.p1.y, line.p2.y)
			for i := y1; i <= y2; i++ {
				board[i][line.p1.x]++
			}
		} else if line.p1.y == line.p2.y {
			x1 := min(line.p1.x, line.p2.x)
			x2 := max(line.p1.x, line.p2.x)
			for i := x1; i <= x2; i++ {
				board[line.p1.y][i]++
			}
		}
	}

	cnt := 0
	for _, row := range board {
		for _, cell := range row {
			if cell >= 2 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
