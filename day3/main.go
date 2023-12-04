package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type maybeGear struct {
	partNumbers []int
}

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	// make the data easier to work with
	input := transformInput(string(contents))

	rows := strings.Split(input, "\n")
	rowLength := len(rows[0])

	sumPart1 := 0
	sumPart2 := 0

	var maybeGears = make(map[string]maybeGear)

	for i := 1; i < len(rows)-1; i++ {
		currentNumber := ""
		gearAdjacency := make(map[string]bool) // FYI this is one way of doing a Set in Go
		isCurrentNumberSelected := false
		for j := 1; j < rowLength; j++ {
			cell := string(rows[i][j])

			if cell == "." || isSymbol(cell) {
				if isCurrentNumberSelected {
					isCurrentNumberSelected = false
					x, _ := strconv.Atoi(currentNumber)
					sumPart1 += x
					for k := range gearAdjacency {
						addToMaybeGears(maybeGears, k, x)
					}
				}
				currentNumber = ""
				gearAdjacency = make(map[string]bool)
				continue
			}

			// we have a digit
			currentNumber += cell

			// check moore neighbourhood basically
			// this is so ugly xD lol
			if isSymbol(string(rows[i-1][j-1])) {
				isCurrentNumberSelected = true
				if isAsterisk(string(rows[i-1][j-1])) {
					addToGearAdjacencySet(gearAdjacency, makeKey(i-1, j-1))
				}
			}
			if isSymbol(string(rows[i-1][j])) {
				isCurrentNumberSelected = true
				if isAsterisk(string(rows[i-1][j])) {
					addToGearAdjacencySet(gearAdjacency, makeKey(i-1, j))
				}
			}
			if isSymbol(string(rows[i-1][j+1])) {
				isCurrentNumberSelected = true
				if isAsterisk(string(rows[i-1][j+1])) {
					addToGearAdjacencySet(gearAdjacency, makeKey(i-1, j+1))
				}
			}
			if isSymbol(string(rows[i][j-1])) {
				isCurrentNumberSelected = true
				if isAsterisk(string(rows[i][j-1])) {
					addToGearAdjacencySet(gearAdjacency, makeKey(i, j-1))
				}
			}
			if isSymbol(string(rows[i][j+1])) {
				isCurrentNumberSelected = true
				if isAsterisk(string(rows[i][j+1])) {
					addToGearAdjacencySet(gearAdjacency, makeKey(i, j+1))
				}
			}
			if isSymbol(string(rows[i+1][j-1])) {
				isCurrentNumberSelected = true
				if isAsterisk(string(rows[i+1][j-1])) {
					addToGearAdjacencySet(gearAdjacency, makeKey(i+1, j-1))
				}
			}
			if isSymbol(string(rows[i+1][j])) {
				isCurrentNumberSelected = true
				if isAsterisk(string(rows[i+1][j])) {
					addToGearAdjacencySet(gearAdjacency, makeKey(i+1, j))
				}
			}
			if isSymbol(string(rows[i+1][j+1])) {
				isCurrentNumberSelected = true
				if isAsterisk(string(rows[i+1][j+1])) {
					addToGearAdjacencySet(gearAdjacency, makeKey(i+1, j+1))
				}
			}
		}
	}

	for _, mg := range maybeGears {
		if len(mg.partNumbers) == 2 {
			sumPart2 += mg.partNumbers[0] * mg.partNumbers[1]
		}
	}

	fmt.Printf("Answer to part 1: %d\n", sumPart1)
	fmt.Printf("Answer to part 2: %d\n", sumPart2)
}

func transformInput(input string) string {
	inputRows := strings.Split(input, "\n")
	rowLength := len(inputRows[0])
	emptyRow := strings.Repeat(".", rowLength+2)

	var rows []string
	rows = append(rows, emptyRow)
	for _, row := range inputRows {
		if len(row) < 1 {
			continue
		}
		rows = append(rows, strings.Join([]string{".", row, "."}, "")) // not the most efficient way surely xD
	}
	rows = append(rows, emptyRow)
	return strings.Join(rows, "\n")
}

func isSymbol(s string) bool {
	// yeah, anything not being period or digit will do
	r, _ := utf8.DecodeRuneInString(s)
	return !unicode.IsDigit(r) && r != '.'
}

func isAsterisk(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return r == '*'
}

func makeKey(i int, j int) string {
	return fmt.Sprintf("%d_%d", i, j)
}

func addToMaybeGears(maybeGears map[string]maybeGear, key string, currentNumber int) {
	mg, exists := maybeGears[key]
	if !exists {
		maybeGears[key] = maybeGear{
			partNumbers: []int{currentNumber},
		}
	} else {
		mg.partNumbers = append(mg.partNumbers, currentNumber)
		maybeGears[key] = mg
	}
}

func addToGearAdjacencySet(set map[string]bool, key string) {
	_, exists := set[key]
	if !exists {
		set[key] = true
	}
}
