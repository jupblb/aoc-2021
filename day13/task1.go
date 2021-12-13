package main

import (
	"fmt"
	"os"
)

const Size = 2000

func main() {
	var manual [Size][Size]bool

	for {
		var x, y int
		if _, err := fmt.Scanf("%d,%d", &x, &y); err != nil {
			break
		}
		manual[y][x] = true
	}

	var dimension rune
	var position int
	fmt.Scanf("fold along %c=%d", &dimension, &position)

	switch dimension {
	case 'x':
		for y := 0; y < Size; y++ {
			for x := 0; x < position; x++ {
				manual[y][x] = manual[y][x] || manual[y][2*position-x]
			}
		}
		for y := 0; y < Size; y++ {
			for x := position; x < Size; x++ {
				manual[y][x] = false
			}
		}
	case 'y':
		for y := 0; y < position; y++ {
			for x := 0; x < Size; x++ {
				manual[y][x] = manual[y][x] || manual[2*position-y][x]
			}
		}
		for y := position; y < Size; y++ {
			for x := 0; x < Size; x++ {
				manual[y][x] = false
			}
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid dimension:", dimension)
		return
	}

	cnt := 0
	for _, row := range manual {
		for _, cell := range row {
			if cell {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
