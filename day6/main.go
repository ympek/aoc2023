package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcWaysToWin(time int, distance int) int {
	waysToWin := 0
	for holdFor := 1; holdFor < time; holdFor++ {
		result := holdFor * (time - holdFor)
		isRecord := result > distance
		if isRecord {
			waysToWin++
		}
	}
	return waysToWin
}

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])

	combinedTimeStr := ""
	combinedDistanceStr := ""
	ansPart1 := 1

	for i := 1; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		dist, _ := strconv.Atoi(distances[i])
		combinedTimeStr += times[i]
		combinedDistanceStr += distances[i]
		ansPart1 *= calcWaysToWin(time, dist)
	}

	combinedTime, _ := strconv.Atoi(combinedTimeStr)
	combinedDistance, _ := strconv.Atoi(combinedDistanceStr)
	ansPart2 := calcWaysToWin(combinedTime, combinedDistance)

	fmt.Printf("Answer to part 1: %d\n", ansPart1)
	fmt.Printf("Answer to part 2: %d\n", ansPart2)
}
