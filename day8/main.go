package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	left  string
	right string
}

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents[:]), "\n")

	nodes := make(map[string]node)
	var instructions string
	var startingNodes []string

	for i, line := range lines {
		if len(line) < 1 {
			continue
		}
		if i == 0 {
			instructions = line
			continue
		}
		definition := strings.Split(line, "=")
		key := strings.TrimSpace(definition[0])
		hops := strings.Split(strings.TrimSpace(definition[1]), ",")

		if key[2] == 'A' {
			// ends with A
			startingNodes = append(startingNodes, key)
		}
		currentNode := node{
			left:  strings.TrimSpace(strings.TrimPrefix(hops[0], "(")),
			right: strings.TrimSpace(strings.TrimSuffix(hops[1], ")")),
		}
		nodes[key] = currentNode
	}

	steps := 0
	stepsToZZZ := 0
	count := 0

	targetCount := len(startingNodes) // 6
	stepsRequired := make([]int, targetCount)

	for count != targetCount {
		for _, c := range instructions {
			steps++
			if c == 'L' {
				for j, n := range startingNodes {
					if nodes[n].left[2] == 'Z' {
						if stepsRequired[j] != 0 {
							stepsRequired[j] = steps
							count++
							if nodes[n].left[0] == 'Z' && nodes[n].left[1] == 'Z' {
								stepsToZZZ = steps
							}
						}
					}
					startingNodes[j] = nodes[n].left
				}
			}
			if c == 'R' {
				for j, n := range startingNodes {
					if nodes[n].right[2] == 'Z' {
						stepsRequired[j] = steps
						count++
						if nodes[n].right[0] == 'Z' && nodes[n].right[1] == 'Z' {
							stepsToZZZ = steps
						}
					}
					startingNodes[j] = nodes[n].right
				}
			}
		}
	}

	fmt.Printf("Answer to part 1: %d\n", stepsToZZZ)

	a := stepsRequired[0]
	b := stepsRequired[1]
	lcm := a * b / findGCD(a, b)

	for _, num := range stepsRequired[2:] {
		lcm = lcm * num / findGCD(lcm, num)
	}
	fmt.Printf("Answer to part 2: %d\n", lcm)
}

func findGCD(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}
