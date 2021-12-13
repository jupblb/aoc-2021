package main

import (
	"fmt"
	"math"
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

	minX, minY := math.MaxInt32, math.MaxInt32

	for {
		var dimension rune
		var position int
		if _, err := fmt.Scanf("fold along %c=%d", &dimension, &position); err != nil {
			break
		}

		switch dimension {
		case 'x':
			if position < minX {
				minX = position
			}
			for y := 0; y < Size; y++ {
				for x := 0; x < position; x++ {
					manual[y][x] = manual[y][x] || manual[y][2*position-x]
				}
			}
		case 'y':
			if position < minY {
				minY = position
			}
			for y := 0; y < position; y++ {
				for x := 0; x < Size; x++ {
					manual[y][x] = manual[y][x] || manual[2*position-y][x]
				}
			}
		default:
			fmt.Fprintln(os.Stderr, "Invalid dimension:", dimension)
			return
		}
	}

	for y := 0; y < minY; y++ {
		for x := 0; x < minX; x++ {
			if manual[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
