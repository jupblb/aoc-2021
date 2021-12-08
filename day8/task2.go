package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://stackoverflow.com/a/22689818
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// poor levenshtein substitute
func diffStr(a, b string) (int, int) {
	diffA := 0
	for _, charA := range a {
		isInB := false
		for _, charB := range b {
			if charA == charB {
				isInB = true
			}
		}
		if !isInB {
			diffA++
		}
	}

	diffB := 0
	for _, charB := range b {
		isInA := false
		for _, charA := range a {
			if charA == charB {
				isInA = true
			}
		}
		if !isInA {
			diffB++
		}
	}

	return diffA, diffB
}

func main() {
	res := 0

	for {
		var digits [10]string
		var display [4]string

		_, err := fmt.Scanf(
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&digits[0], &digits[1], &digits[2], &digits[3], &digits[4],
			&digits[5], &digits[6], &digits[7], &digits[8], &digits[9],
			&display[0], &display[1], &display[2], &display[3])
		if err != nil {
			break
		}

		for i := 0; i < 10; i++ {
			digits[i] = sortString(digits[i])
		}
		for i := 0; i < 4; i++ {
			display[i] = sortString(display[i])
		}

		var numbers [10]string

		for _, digit := range digits {
			l := len(digit)
			if l == 2 {
				numbers[1] = digit
			}
			if l == 4 {
				numbers[4] = digit
			}
			if l == 3 {
				numbers[7] = digit
			}
			if l == 7 {
				numbers[8] = digit
			}
		}

		for _, digit := range digits {
			diffOneA, diffOneB := diffStr(digit, numbers[1])
			diffFourA, diffFourB := diffStr(digit, numbers[4])
			diffSevenA, diffSevenB := diffStr(digit, numbers[7])
			_, diffEightB := diffStr(digit, numbers[8])

			if diffOneA == 3 && diffOneB == 0 &&
				diffFourA == 2 && diffFourB == 1 &&
				diffSevenA == 2 && diffSevenB == 0 &&
				diffEightB == 2 {
				numbers[3] = digit
			}
			if diffOneA == 4 && diffOneB == 1 &&
				diffFourA == 2 && diffFourB == 1 &&
				diffSevenA == 3 && diffSevenB == 1 &&
				diffEightB == 2 {
				numbers[5] = digit
			}
			if diffOneA == 5 && diffOneB == 1 &&
				diffFourA == 3 && diffFourB == 1 &&
				diffSevenA == 4 && diffSevenB == 1 &&
				diffEightB == 1 {
				numbers[6] = digit
			}
			if diffOneA == 4 && diffOneB == 1 &&
				diffFourA == 3 && diffFourB == 2 &&
				diffSevenA == 3 && diffSevenB == 1 &&
				diffEightB == 2 {
				numbers[2] = digit
			}
			if diffOneA == 4 && diffOneB == 0 &&
				diffFourA == 2 && diffFourB == 0 &&
				diffSevenA == 3 && diffSevenB == 0 &&
				diffEightB == 1 {
				numbers[9] = digit
			}
		}

		partialRes := 0
		for i, digit := range display {
			mult := 1
			for j := 0; j < 3-i; j++ {
				mult *= 10
			}

			for j := 0; j < 10; j++ {
				if digit == numbers[j] {
					partialRes += mult * j
				}
			}
		}
		res += partialRes
	}

	fmt.Println(res)
}
