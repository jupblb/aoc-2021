package main

import (
	"fmt"
	"sort"
)

func main() {
	var res []int

	for {
		var input string
		if _, err := fmt.Scanf("%s", &input); err != nil {
			break
		}

		stack := make([]rune, len(input))
		stackIt := -1
		isValid := true

	loop:
		for _, brace := range input {
			switch brace {
			case ')':
				if stack[stackIt] == '(' {
					stackIt--
				} else {
					isValid = false
					break loop
				}
			case ']':
				if stack[stackIt] == '[' {
					stackIt--
				} else {
					isValid = false
					break loop
				}
			case '}':
				if stack[stackIt] == '{' {
					stackIt--
				} else {
					isValid = false
					break loop
				}
			case '>':
				if stack[stackIt] == '<' {
					stackIt--
				} else {
					isValid = false
					break loop
				}
			default:
				stackIt++
				stack[stackIt] = brace
			}
		}

		if isValid {
			score := 0
			for stackIt >= 0 {
				score *= 5

				switch stack[stackIt] {
				case '(':
					score += 1
				case '[':
					score += 2
				case '{':
					score += 3
				case '<':
					score += 4
				}

				stackIt--
			}
			res = append(res, score)
		}
	}

	sort.Ints(res)
	fmt.Println(res[len(res)/2])
}
