package main

import "fmt"

func main() {
	depth := 0
	horizontal := 0

	for {
		var command string
		var value int

		_, err := fmt.Scanf("%s %d", &command, &value)
		if err != nil {
			break
		}

		switch command {
		case "down":
			depth += value
		case "up":
			depth -= value
		case "forward":
			horizontal += value
		}
	}

	fmt.Println(depth * horizontal)
}
