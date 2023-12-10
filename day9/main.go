package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	sumPart1 := 0
	sumPart2 := 0

	lines := strings.Split(string(contents[:]), "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		valuesAsString := strings.Split(line, " ")
		numValues := len(valuesAsString)
		currentHistory := make([][]int, numValues)
		currentHistory[0] = make([]int, numValues)

		idx := 0
		for i, s := range valuesAsString {
			x, _ := strconv.Atoi(s)
			currentHistory[idx][i] = x
		}

		allZeroes := false
		for !allZeroes {
			numZeroes := 0
			idx++
			numValues--
			currentHistory[idx] = make([]int, numValues)

			for i := 0; i < numValues; i++ {
				a := currentHistory[idx-1][i]
				b := currentHistory[idx-1][i+1]
				diff := a - b
				currentHistory[idx][i] = diff
				if diff == 0 {
					numZeroes++
				}
			}

			allZeroes = numZeroes == numValues
		}

		currentHistory[idx] = append([]int{0}, currentHistory[idx]...)
		currentHistory[idx] = append(currentHistory[idx], 0)
		numValues++
		for idx != 0 {
			idx--

			nextPrediction := currentHistory[idx][numValues-1] - currentHistory[idx+1][numValues]
			currentHistory[idx] = append(currentHistory[idx], nextPrediction)
			prevPrediction := currentHistory[idx][0] + currentHistory[idx+1][0]
			currentHistory[idx] = append([]int{prevPrediction}, currentHistory[idx]...)

			numValues++
		}

		lastIndex := len(currentHistory[0]) - 1
		sumPart1 += currentHistory[0][lastIndex]
		sumPart2 += currentHistory[0][0]
	}

	fmt.Printf("Answer to part 1: %d\n", sumPart1)
	fmt.Printf("Answer to part 2: %d\n", sumPart2)
}
