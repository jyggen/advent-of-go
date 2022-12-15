package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func createGrid(input string, withFloor bool) ([]bool, int) {
	minX := 0
	maxX := 1000
	minY := 0
	maxY := 0
	coordinatePairs := make([][][2]int, 0)

	for _, r := range utils.ToStringSlice(input, "\n") {
		coordinates := make([][2]int, 0)
		for _, p := range utils.ToStringSlice(r, " -> ") {
			c := utils.ToStringSlice(p, ",")
			x, _ := strconv.Atoi(c[0])
			y, _ := strconv.Atoi(c[1])
			maxY = utils.MaxInt(maxY, y)
			coordinates = append(coordinates, [2]int{x, y})
		}

		coordinatePairs = append(coordinatePairs, coordinates)
	}

	rowLength := (maxY + 3) - minY
	columnLength := maxX - minX
	grid := make([]bool, columnLength*rowLength)

	for _, cp := range coordinatePairs {
		for i := 1; i < len(cp); i++ {
			if cp[i][0] == cp[i-1][0] {
				for y := utils.MinInt(cp[i][1], cp[i-1][1]); y <= utils.MaxInt(cp[i][1], cp[i-1][1]); y++ {
					grid[utils.FromCoordinates(cp[i][0]-minX, y-minY, columnLength)] = true
				}
			} else {
				for x := utils.MinInt(cp[i][0], cp[i-1][0]); x <= utils.MaxInt(cp[i][0], cp[i-1][0]); x++ {
					grid[utils.FromCoordinates(x-minX, cp[i][1]-minY, columnLength)] = true
				}
			}
		}
	}

	if withFloor {
		for i := len(grid) - columnLength; i < len(grid); i++ {
			grid[i] = true
		}
	}

	return grid, columnLength
}

func sandPhysics(grid []bool, columnLength int, sandX int, sandY int) int {
	atRest := 0
	gridLength := len(grid)
	sandIndex := utils.FromCoordinates(sandX, sandY, columnLength)
	topY := gridLength / columnLength
	var possibilities [3]int

Loop:
	for {

		possibilities[0] = sandIndex + columnLength
		possibilities[1] = possibilities[0] - 1
		possibilities[2] = possibilities[0] + 1

		for _, possibility := range possibilities {
			if possibility >= gridLength {
				break Loop
			}

			if !grid[possibility] {
				sandIndex = possibility
				continue Loop
			}
		}

		if grid[sandIndex] {
			break Loop
		}

		grid[sandIndex] = true
		topY = utils.MinInt(topY, sandIndex/columnLength)
		sandIndex = utils.FromCoordinates(sandX, utils.MaxInt(0, topY-1), columnLength)
		atRest++
	}

	return atRest
}

func SolvePart1(input string) (string, error) {
	grid, columnLength := createGrid(input, false)

	return strconv.Itoa(sandPhysics(grid, columnLength, 500, 0)), nil
}

func SolvePart2(input string) (string, error) {
	grid, columnLength := createGrid(input, true)

	return strconv.Itoa(sandPhysics(grid, columnLength, 500, 0)), nil
}
