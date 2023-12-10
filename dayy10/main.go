package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type pos struct {
	i int
	j int
}

type route struct {
	from string
	i    int
	j    int
}

var grid [][]rune
var startPos pos

func main() {
	contents, err := os.ReadFile("./input_t")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents[:]), "\n")
	// how long is the line?
	width := len(lines[0])
	height := len(lines)
	grid = make([][]rune, height+1) // why not +2? well the last line is empty string but it's there and counts
	// make the top border
	grid[0] = make([]rune, width+2)
	grid[height] = make([]rune, width+2)
	for i := 0; i < width+2; i++ {
		grid[0][i] = '.'
		grid[height][i] = '.'
	}

	for i, line := range lines {
		// take the trash out
		if len(line) < 1 {
			continue
		}
		// make grid
		// and find S
		grid[i+1] = make([]rune, width+2)
		grid[i+1][0] = '.'
		grid[i+1][width+1] = '.'
		for j, cell := range line {
			grid[i+1][j+1] = cell
			if cell == 'S' {
				startPos.i = i + 1
				startPos.j = j + 1
			}
		}
	}

	currentI := startPos.i
	currentJ := startPos.j
	currentFrom := ""

	var routes []route
	// Figure out where we can go.
	northI := currentI - 1
	northJ := currentJ
	if canComeFromSouth(grid[northI][northJ]) {
		fmt.Println("Can go north from S")
		routes = append(routes, route{
			from: "south",
			i:    northI,
			j:    northJ,
		})
	}
	southI := currentI + 1
	southJ := currentJ
	if canComeFromNorth(grid[southI][southJ]) {
		fmt.Println("Can go south from S")
		routes = append(routes, route{
			from: "north",
			i:    southI,
			j:    southJ,
		})
	}

	westI := currentI
	westJ := currentJ - 1
	if canComeFromEast(grid[westI][westJ]) {
		fmt.Println("Can go west from S")
		routes = append(routes, route{
			from: "east",
			i:    westI,
			j:    westJ,
		})
	}

	eastI := currentI
	eastJ := currentJ + 1
	if canComeFromWest(grid[eastI][eastJ]) {
		fmt.Println("Can go east from S")
		routes = append(routes, route{
			from: "west",
			i:    eastI,
			j:    eastJ,
		})
	}

	// Assuming there are only 2 routes possible
	if len(routes) != 2 {
		panic("Didn't really expect input like this")
	}

	var visited1 []pos // going north
	var visited2 []pos // going south

	for i := 0; i < width*height; i++ {
		if currentI == startPos.i && currentJ == startPos.j && i > 0 {
			fmt.Println("loop")
			break
		}
		if i == 0 {
			currentFrom, currentI, currentJ = visitFrom(&visited1, routes[0].from, routes[0].i, routes[0].j)
		}
		currentFrom, currentI, currentJ = visitFrom(&visited1, currentFrom, currentI, currentJ)
		if currentFrom == "closed" {
			fmt.Println("route closed!", currentI, currentJ)
			break
		}
	}

	for i := 0; i < width*height; i++ {
		if currentI == startPos.i && currentJ == startPos.j && i > 0 {
			fmt.Println("loop")
			break
		}
		if i == 0 {
			currentFrom, currentI, currentJ = visitFrom(&visited2, routes[1].from, routes[1].i, routes[1].j)
		}
		currentFrom, currentI, currentJ = visitFrom(&visited2, currentFrom, currentI, currentJ)
		if currentFrom == "closed" {
			fmt.Println("route closed!", currentI, currentJ)
			break
		}
	}

	fmt.Println(len(visited1))
	fmt.Println(len(visited2))

	for i := 0; i < len(visited1); i++ {
		if visited1[i].i == visited2[i].i && visited1[i].j == visited2[i].j {
			fmt.Println(visited1[i].i, visited2[i].i, visited1[i].j, visited2[i].j)
			fmt.Printf("Answer to part1: %d\n", i-1) // -1 because S is also visited
		}
	}

	// fmt.Printf("Answer to part 2: %d\n", sumPart2)
	modifyGrid(visited1, grid, width, height)
	insideCount := markInsides(grid, width, height)
	printGrid(grid, width, height)

	fmt.Printf("Answer to part 2: %d\n", insideCount)
}

func visitFrom(visited *[]pos, from string, currentI, currentJ int) (string, int, int) {
	isRouteClosed := true
	*visited = append(*visited, pos{i: currentI, j: currentJ})

	northI := currentI - 1
	northJ := currentJ
	if from != "north" && canGoNorth(grid[currentI][currentJ]) && canComeFromSouth(grid[northI][northJ]) {
		isRouteClosed = false
		return "south", northI, northJ
	}
	southI := currentI + 1
	southJ := currentJ
	if from != "south" && canGoSouth(grid[currentI][currentJ]) && canComeFromNorth(grid[southI][southJ]) {
		isRouteClosed = false
		return "north", southI, southJ
	}

	westI := currentI
	westJ := currentJ - 1
	if from != "west" && canGoWest(grid[currentI][currentJ]) && canComeFromEast(grid[westI][westJ]) {
		isRouteClosed = false
		return "east", westI, westJ
	}

	eastI := currentI
	eastJ := currentJ + 1
	if from != "east" && canGoEast(grid[currentI][currentJ]) && canComeFromWest(grid[eastI][eastJ]) {
		isRouteClosed = false
		return "west", eastI, eastJ
	}

	if isRouteClosed {
		fmt.Println("CLOSED ROUTE AT", currentI, currentJ)
		// fmt.Printf("Came from: %s. north: %c south: %c west: %c east: %c\n", from, grid[northI][northJ], grid[southI][southJ], grid[westI][westJ], grid[eastI][eastJ])
		return "closed", currentI, currentJ
	}
	return "closed", 0, 0
}

func canComeFromNorth(r rune) bool {
	return r == 'S' || r == '|' || r == 'L' || r == 'J'
}

func canComeFromSouth(r rune) bool {
	return r == 'S' || r == '|' || r == '7' || r == 'F'
}

func canComeFromWest(r rune) bool {
	return r == 'S' || r == '-' || r == 'J' || r == '7'
}

func canComeFromEast(r rune) bool {
	return r == 'S' || r == '-' || r == 'L' || r == 'F'
}

func canGoNorth(r rune) bool {
	return r == 'S' || r == '|' || r == 'L' || r == 'J'
}

func canGoSouth(r rune) bool {
	return r == 'S' || r == '|' || r == '7' || r == 'F'
}

func canGoWest(r rune) bool {
	return r == 'S' || r == '-' || r == '7' || r == 'J'
}

func canGoEast(r rune) bool {
	return r == 'S' || r == '-' || r == 'F' || r == 'L'
}

func printGrid(grid [][]rune, width int, height int) {
	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}

func modifyGrid(visited []pos, grid [][]rune, width int, height int) {
	// for _, p := range visited {
	// 	if grid[p.i][p.j] == 'S' {
	// 		continue
	// 	}
	// 	grid[p.i][p.j] = 'x'
	// }

	// for i := 0; i < width+2; i++ {
	// 	grid[0][i] = 'x'
	// 	grid[height][i] = 'x'
	// }

	// for i := 0; i < height; i++ {
	// 	grid[i+1][0] = 'x'
	// 	grid[i+1][width+1] = 'x'
	// }

	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			if grid[i][j] == 'S' {
				continue
			}
			if !slices.Contains(visited, pos{i: i, j: j}) {
				grid[i][j] = 'x'
			}
		}
	}
}

func markInsides(grid [][]rune, width, height int) int {
	result := 0
	for i := 1; i <= height-1; i++ {
		for j := 1; j <= width-1; j++ {
			// moore again
			if grid[i][j] == 'x' {
				continue
			}
			count := 0
			if grid[i-1][j-1] == 'x' {
				count++
			}
			if grid[i-1][j] == 'x' {
				count++
			}
			if grid[i-1][j+1] == 'x' {
				count++
			}
			if grid[i][j-1] == 'x' {
				count++
			}
			if grid[i][j+1] == 'x' {
				count++
			}
			if grid[i+1][j-1] == 'x' {
				count++
			}
			if grid[i+1][j] == 'x' {
				count++
			}
			if grid[i+1][j+1] == 'x' {
				count++
			}
			if count == 8 {
				grid[i][j] = '`'
				result++
			}
		}
	}
	return result
}
