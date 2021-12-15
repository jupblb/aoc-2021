package main

import (
	"fmt"
	"strings"
)

const Steps = 10

func main() {
	var polymer string
	fmt.Scanf("%s\n\n", &polymer)

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
		var newPolymer strings.Builder
		for i := 1; i < len(polymer); i++ {
			newPolymer.WriteByte(polymer[i-1])

			input := string(polymer[i-1]) + string(polymer[i])
			if to, ok := rules[input]; ok {
				newPolymer.WriteRune(to)
			}
		}
		newPolymer.WriteByte(polymer[len(polymer)-1])
		polymer = newPolymer.String()
	}

	var cnt [256]int
	for _, c := range polymer {
		cnt[int(c)]++
	}

	minOcc, maxOcc := len(polymer), 0
	for _, occ := range cnt {
		if occ < minOcc && occ != 0 {
			minOcc = occ
		}
		if occ > maxOcc {
			maxOcc = occ
		}
	}
	fmt.Println(maxOcc - minOcc)
}
