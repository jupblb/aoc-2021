package main

import "fmt"

func main() {
	res := 0

	for {
		var input string
		if _, err := fmt.Scanf("%s", &input); err != nil {
			break
		}

		stack := make([]rune, len(input))
		stackIt := -1

	loop:
		for _, brace := range input {
			switch brace {
			case ')':
				if stack[stackIt] == '(' {
					stackIt--
				} else {
					res += 3
					break loop
				}
			case ']':
				if stack[stackIt] == '[' {
					stackIt--
				} else {
					res += 57
					break loop
				}
			case '}':
				if stack[stackIt] == '{' {
					stackIt--
				} else {
					res += 1197
					break loop
				}
			case '>':
				if stack[stackIt] == '<' {
					stackIt--
				} else {
					res += 25137
					break loop
				}
			default:
				stackIt++
				stack[stackIt] = brace
			}
		}
	}

	fmt.Println(res)
}
