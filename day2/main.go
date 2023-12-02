package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hand struct {
	red   int
	green int
	blue  int
}

type game struct {
	id    int
	hands []hand
}

func getGamesFromInputFile() []game {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents[:]), "\n")

	var games []game
	for i, line := range lines {
		var currentGame game
		currentGame.id = i + 1

		if len(line) < 1 { // remember to take out the trash
			continue
		}

		gameDescription := strings.Split(line, ":")[1]
		handsInGame := strings.Split(gameDescription, ";")
		for _, h := range handsInGame {
			var currentHand hand
			colorsInHand := strings.Split(h, ",")
			for _, c := range colorsInHand {
				pair := strings.Split(strings.TrimSpace(c), " ")
				count := pair[0]
				color := pair[1]
				if color == "green" {
					currentHand.green, _ = strconv.Atoi(count)
				}
				if color == "red" {
					currentHand.red, _ = strconv.Atoi(count)
				}
				if color == "blue" {
					currentHand.blue, _ = strconv.Atoi(count)
				}
			}
			currentGame.hands = append(currentGame.hands, currentHand)
		}
		games = append(games, currentGame)
	}
	return games
}

func main() {
	games := getGamesFromInputFile()

	sumPart1 := 0
	sumPart2 := 0

	for _, g := range games {
		isPossible := true
		fewestPossibleReds := 0
		fewestPossibleGreens := 0
		fewestPossibleBlues := 0
		for _, h := range g.hands {
			// which games would have been possible if the bag contained only
			// 12 red cubes, 13 green cubes, and 14 blue cubes?
			if h.red > 12 || h.green > 13 || h.blue > 14 {
				isPossible = false
			}
			if h.red > fewestPossibleReds {
				fewestPossibleReds = h.red
			}
			if h.green > fewestPossibleGreens {
				fewestPossibleGreens = h.green
			}
			if h.blue > fewestPossibleBlues {
				fewestPossibleBlues = h.blue
			}
		}

		if isPossible {
			sumPart1 += g.id
		}

		sumPart2 += fewestPossibleReds * fewestPossibleBlues * fewestPossibleGreens
	}

	fmt.Printf("Answer to part 1: %d\n", sumPart1)
	fmt.Printf("Answer to part 2: %d\n", sumPart2)
}
