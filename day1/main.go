package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var subs = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func substituteFirstAndLastWordsToDigits(haystack string) string {
	tmp := haystack
	firstWordIndex := len(tmp)
	firstWord := ""
	lastWordIndex := -1
	lastWord := ""

	for key := range subs {
		index := strings.Index(tmp, key)
		if index != -1 {
			if index < firstWordIndex {
				firstWordIndex = index
				firstWord = key
			}
		}
		lastIndex := strings.LastIndex(tmp, key)
		if lastIndex != -1 {
			if lastIndex > lastWordIndex {
				lastWordIndex = lastIndex
				lastWord = key
			}
		}
	}

	if firstWord == "" {
		// no words - skip
		return haystack
	}

	tmp = tmp[:firstWordIndex] + subs[firstWord] + tmp[firstWordIndex+1:]
	tmp = tmp[:lastWordIndex] + subs[lastWord] + tmp[lastWordIndex+1:]

	return tmp
}

func calcCalibrationValue(line string) int {
	firstFound := false
	first := 0
	last := 0

	for _, runeVal := range line {
		if unicode.IsDigit(runeVal) {
			if !firstFound {
				first = int(runeVal - '0')
				firstFound = true
			}
			last = int(runeVal - '0')
		}
	}

	return first*10 + last
}

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents[:]), "\n")
	sumPart1 := 0
	sumPart2 := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		sumPart1 += calcCalibrationValue(line)
		parsedLine := substituteFirstAndLastWordsToDigits(line)
		sumPart2 += calcCalibrationValue(parsedLine)
	}

	fmt.Printf("Answer to part 1: %d\n", sumPart1)
	fmt.Printf("Answer to part 2: %d\n", sumPart2)
}
