package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type IntHeap []int

const MaxSize = 500

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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

func main() {
	var cave [][]int
	for {
		if row, ok := readNumbers(); ok {
			cave = append(cave, row)
		} else {
			break
		}
	}

	originalSize := len(cave)
	for i := 0; i < originalSize; i++ {
		for j := originalSize; j < originalSize*5; j++ {
			cave[i] = append(cave[i], cave[i][j-originalSize]+1)
			if cave[i][j] == 10 {
				cave[i][j] = 1
			}
		}
	}

	for i := originalSize; i < originalSize*5; i++ {
		var newRow []int
		for j := 0; j < originalSize*5; j++ {
			newRow = append(newRow, cave[i-originalSize][j]+1)
			if newRow[j] == 10 {
				newRow[j] = 1
			}
		}
		cave = append(cave, newRow)
	}

	queue := &IntHeap{0}
	var visited [MaxSize][MaxSize]bool

	for queue.Len() > 0 {
		front := heap.Pop(queue).(int)
		x := front % MaxSize
		y := (front / MaxSize) % MaxSize
		w := (front / MaxSize) / MaxSize

		if x < 0 || y < 0 || x >= len(cave) || y >= len(cave) || visited[x][y] {
			continue
		}
		visited[x][y] = true

		if x == len(cave)-1 && y == len(cave)-1 {
			fmt.Println(w + cave[x][y])
			return
		}

		heap.Push(queue, (x-1) + (y) * 500 + (w + cave[x][y]) * 500 * 500)
		heap.Push(queue, (x+1) + (y) * 500 + (w + cave[x][y]) * 500 * 500)
		heap.Push(queue, (x) + (y-1) * 500 + (w + cave[x][y]) * 500 * 500)
		heap.Push(queue, (x) + (y+1) * 500 + (w + cave[x][y]) * 500 * 500)
	}
}
