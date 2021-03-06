package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func sortInt(input int) int {

	swapped, _ := strconv.Atoi(bubbleSort(strconv.Itoa(input)))
	return swapped

}

func bubbleSort(word string) string {
	wordtable := strings.Split(word, "")
	for j := 0; j < len(word); j++ {

		for i := 0; i < len(word)-1; i++ {
			if wordtable[i] < wordtable[i+1] {
				temp := wordtable[i]
				wordtable[i] = wordtable[i+1]
				wordtable[i+1] = temp
			}
		}
	}
	return strings.Join(wordtable, "")
}

func intLength(n int) int {
	return 1 + int(math.Log10(float64(n)))
}

func isPandigital(n int) bool {

	height := intLength(n)

	output := 0

	for i := 1; i < height+1; i++ {
		current := 1
		for j := 1; j < i; j++ {
			current *= 10
		}
		output += (current * i)
	}

	return output == sortInt(n)
}

func concatenInt(a int, b int) int {
	answer, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return answer
}

func main() {

	winner := 1

	for n := 1; n < 99999; n++ {
		for count := 1; count < 9; count++ {

			pandigit := n

			for i := 2; i <= count; i++ {
				pandigit = concatenInt(pandigit, i*n)
			}

			if pandigit > winner && isPandigital(pandigit) {
				fmt.Println(pandigit)
				winner = pandigit

			}

		}
	}
}
