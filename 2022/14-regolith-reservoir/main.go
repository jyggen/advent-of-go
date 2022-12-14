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

const (
	EMPTY = iota
	WALL
	SAND
)

func isOutOfRange(coordinates [2]int, columnLength int, rowLength int) bool {
	return coordinates[0] < 0 || coordinates[1] < 0 || coordinates[0] == columnLength || coordinates[1] == rowLength
}

func sandPhysics(grid []int, columnLength int, rowLength int, sandX int, sandY int) int {
	atRest := 0

Outer:
	for {
		sand := [2]int{sandX, sandY}

	Inner:
		for {
			sandIndex := utils.FromCoordinates(sand[0], sand[1], columnLength)

			if isOutOfRange([2]int{sand[0], sand[1] + 1}, columnLength, rowLength) {
				break Outer
			}

			switch grid[sandIndex+columnLength] {
			case EMPTY:
				sand[1]++
				continue Inner
			}

			if sand[0]-1 < 0 || sand[1]+1 < 0 || sand[0]-1 == columnLength || sand[1]+1 == rowLength {
				break Outer
			}

			switch grid[sandIndex+columnLength-1] {
			case EMPTY:
				sand[1]++
				sand[0]--
				continue Inner
			}

			if sand[0]+1 < 0 || sand[1]+1 < 0 || sand[0]+1 == columnLength || sand[1]+1 == rowLength {
				break Outer
			}

			switch grid[sandIndex+columnLength+1] {
			case EMPTY:
				sand[1]++
				sand[0]++
				continue Inner
			}

			if grid[sandIndex] == SAND {
				break Outer
			}

			grid[sandIndex] = SAND
			atRest++
			break Inner
		}
	}

	return atRest
}

func SolvePart1(input string) (string, error) {
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

	rowLength := (maxY + 1) - minY
	columnLength := maxX - minX
	grid := make([]int, columnLength*rowLength)

	for _, cp := range coordinatePairs {
		for i := 1; i < len(cp); i++ {
			if cp[i][0] == cp[i-1][0] {
				for y := utils.MinInt(cp[i][1], cp[i-1][1]); y <= utils.MaxInt(cp[i][1], cp[i-1][1]); y++ {
					grid[utils.FromCoordinates(cp[i][0]-minX, y-minY, columnLength)] = WALL
				}
			} else {
				for x := utils.MinInt(cp[i][0], cp[i-1][0]); x <= utils.MaxInt(cp[i][0], cp[i-1][0]); x++ {
					grid[utils.FromCoordinates(x-minX, cp[i][1]-minY, columnLength)] = WALL
				}
			}
		}
	}

	return strconv.Itoa(sandPhysics(grid, columnLength, rowLength, 500, 0)), nil
}

func SolvePart2(input string) (string, error) {
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
	grid := make([]int, columnLength*rowLength)

	for i := len(grid) - columnLength; i < len(grid); i++ {
		grid[i] = WALL
	}

	for _, cp := range coordinatePairs {
		for i := 1; i < len(cp); i++ {
			if cp[i][0] == cp[i-1][0] {
				for y := utils.MinInt(cp[i][1], cp[i-1][1]); y <= utils.MaxInt(cp[i][1], cp[i-1][1]); y++ {
					grid[utils.FromCoordinates(cp[i][0]-minX, y-minY, columnLength)] = WALL
				}
			} else {
				for x := utils.MinInt(cp[i][0], cp[i-1][0]); x <= utils.MaxInt(cp[i][0], cp[i-1][0]); x++ {
					grid[utils.FromCoordinates(x-minX, cp[i][1]-minY, columnLength)] = WALL
				}
			}
		}
	}

	return strconv.Itoa(sandPhysics(grid, columnLength, rowLength, 500, 0)), nil
}
