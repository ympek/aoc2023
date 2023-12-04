package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	sumPart1 := 0

	lines := strings.Split(string(contents), "\n")

	copies := make([]int, len(lines))
	for i := 0; i < len(copies); i++ {
		copies[i] = 0
	}

	for i, line := range lines {
		if len(line) < 1 {
			continue
		}

		card := strings.Split(line[9:], "|")
		winningNumbers := strings.Fields(card[0])
		ourNumbers := strings.Fields(card[1])

		var matches []string
		for _, number := range ourNumbers {
			for _, winningNumber := range winningNumbers {
				if number == winningNumber {
					matches = append(matches, winningNumber)
				}
			}
		}
		matchingNumbers := len(matches)
		copies[i] += 1
		if matchingNumbers > 0 {
			points := math.Pow(float64(2), float64(matchingNumbers-1))
			sumPart1 += int(points)

			for j := 0; j < copies[i]; j++ {
				for k := i + 1; k <= i+matchingNumbers; k++ {
					copies[k] += 1
				}
			}
		}
	}

	sumPart2 := 0
	for _, c := range copies {
		sumPart2 += c
	}

	fmt.Printf("Answer to part 1: %d\n", sumPart1)
	fmt.Printf("Answer to part 2: %d\n", sumPart2)
}
