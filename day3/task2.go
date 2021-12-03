package main

import (
	"fmt"
	"os"
	"strconv"
)

func matching(number string, rating []rune, idx int) bool {
	for i := 0; i < idx; i++ {
		if rune(number[i]) != rating[i] {
			return false
		}
	}
	return true
}

func main() {
	var input []string

	for {
		var number string
		_, err := fmt.Scanf("%s", &number)
		if err != nil {
			break
		}
		input = append(input, number)
	}

	length := len(input)
	width := len(input[0])
	oxygenRating := make([]rune, width)
	co2Rating := make([]rune, width)

	for i := 0; i < width; i++ {
		cntOxygen := 0
		cntMatchingOxygen := 0

		cntCo2 := 0
		cntMatchingCo2 := 0

		for j := 0; j < length; j++ {
			if matching(input[j], oxygenRating, i) {
				cntMatchingOxygen++
				switch input[j][i] {
				case '0':
				case '1':
					cntOxygen++
				default:
					fmt.Fprintln(os.Stderr, "Invalid input")
					os.Exit(1)
				}
			}

			if matching(input[j], co2Rating, i) {
				cntMatchingCo2++
				switch input[j][i] {
				case '0':
					cntCo2++
				case '1':
				default:
					fmt.Fprintln(os.Stderr, "Invalid input")
					os.Exit(1)
				}
			}
		}

		if cntOxygen >= cntMatchingOxygen-cntOxygen {
			oxygenRating[i] = '1'
		} else {
			oxygenRating[i] = '0'
		}

		if (cntCo2 > cntMatchingCo2-cntCo2 || cntCo2 == 0) && cntMatchingCo2-cntCo2 != 0 {
			co2Rating[i] = '1'
		} else {
			co2Rating[i] = '0'
		}
	}

	oxygenRatingDec, _ := strconv.ParseInt(string(oxygenRating), 2, 64)
	co2RatingDec, _ := strconv.ParseInt(string(co2Rating), 2, 64)

	fmt.Println(oxygenRatingDec*co2RatingDec)
}
