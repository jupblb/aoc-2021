package main

import (
	"fmt"
	"math"
)

const Steps = 40

func main() {
	var polymer string
	fmt.Scanf("%s\n\n", &polymer)

	fragments := make(map[string]int)
	for i := 1; i < len(polymer); i++ {
		pair := string(polymer[i-1]) + string(polymer[i])
		fragments[pair]++
	}

	rules := make(map[string]rune)
	for {
		var from string
		var to rune
		if _, err := fmt.Scanf("%s -> %c\n", &from, &to); err != nil {
			break
		}
		rules[from] = to
	}

	for step := 0; step < Steps; step++ {
		newFragments := make(map[string]int)

		for pair, occ := range fragments {
			if to, ok := rules[pair]; ok {
				left := string(pair[0]) + string(to)
				right := string(to) + string(pair[1])
				newFragments[left] += occ
				newFragments[right] += occ
			} else {
				newFragments[pair] += occ
			}
		}

		fragments = newFragments
	}

	var cnt [256]int64
	cnt[int(polymer[0])]++
	cnt[int(polymer[len(polymer)-1])]++
	for pair, occ := range fragments {
		cnt[int(pair[0])] += int64(occ)
		cnt[int(pair[1])] += int64(occ)
	}

	minOcc, maxOcc := int64(math.MaxInt64), int64(0)
	for _, occ := range cnt {
		occ /= 2
		if occ < minOcc && occ != 0 {
			minOcc = occ
		}
		if occ > maxOcc {
			maxOcc = occ
		}
	}

	fmt.Println(maxOcc - minOcc)
}
