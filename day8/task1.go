package main

import "fmt"

func main() {
	res := 0

	for {
		var dummy string
		var display [4]string
		_, err := fmt.Scanf(
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &dummy,
			&display[0], &display[1], &display[2], &display[3])
		if err != nil {
			break
		}

		for _, digit := range display {
			l := len(digit)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				res++
			}
		}
	}

	fmt.Println(res)
}
