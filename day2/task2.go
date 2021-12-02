package main

import "fmt"

func main() {
	aim := 0
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
			aim += value
		case "up":
			aim -= value
		case "forward":
			depth += aim * value
			horizontal += value
		}
	}

	fmt.Println(depth * horizontal)
}
