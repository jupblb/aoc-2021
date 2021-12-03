package main

import (
	"fmt"
	"os"
	"strconv"
)

const WordSize = 12

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
	gammaRate := make([]rune, WordSize)
	epsilonRate := make([]rune, WordSize)

	for i := 0; i < WordSize; i++ {
		cnt := 0
		for j := 0; j < length; j++ {
			switch input[j][i] {
			case '0':
			case '1':
				cnt++
			default:
				fmt.Fprintln(os.Stderr, "Invalid input")
				os.Exit(1)
			}
		}
		if cnt > length/2 {
			gammaRate[i] = '1'
			epsilonRate[i] = '0'
		} else {
			gammaRate[i] = '0'
			epsilonRate[i] = '1'
		}
	}

	gammaRateDec, _ := strconv.ParseInt(string(gammaRate), 2, 64)
	epsilonRateDec, _ := strconv.ParseInt(string(epsilonRate), 2, 64)

	fmt.Println(gammaRateDec * epsilonRateDec)
}
