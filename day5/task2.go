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

func (l Line) step() (int, int) {
	sgn := func(n int) int {
		if n > 0 {
			return 1
		}
		if n < 0 {
			return -1
		}
		return 0
	}
	return sgn(l.p2.x - l.p1.x), sgn(l.p2.y - l.p1.y)
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
		dx, dy := line.step()
		for x, y := line.p1.x, line.p1.y; x != line.p2.x || y != line.p2.y; x, y = x+dx, y+dy {
			board[y][x]++
		}
		board[line.p2.y][line.p2.x]++
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
