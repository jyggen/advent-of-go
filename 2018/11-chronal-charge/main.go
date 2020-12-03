package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"os"
	"strconv"
	"strings"
)

const size = 300

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	serialNumber, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		return "", err
	}

	grid := makeGrid(serialNumber)

	maxValue := 0
	maxPos := ""

	for x := 1; x < size-2; x++ {
		for y := 1; y < size-2; y++ {
			value := 0
			for _, i := range grid[x : x+3] {
				for _, j := range i[y : y+3] {
					value += j
				}
			}

			if value > maxValue {
				maxPos = fmt.Sprintf("%d,%d", x, y)
				maxValue = value
			}
		}
	}

	return maxPos, nil
}

func SolvePart2(input string) (string, error) {
	serialNumber, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		return "", err
	}

	grid := makeGrid(serialNumber)

	maxPower := 0
	maxCoords := ""

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			gridSize := 0

			if x < y {
				gridSize = x + 1
			} else {
				gridSize = y + 1
			}

			if gridSize > 32 {
				gridSize = 32
			}

			for i := 1; i <= gridSize; i++ {
				gridPower := calculateGridValue(grid, x, y, i)

				if gridPower > maxPower {
					maxCoords = fmt.Sprintf("%d,%d,%d", x-i+1, y-i+1, i)
					maxPower = gridPower
				}
			}
		}
	}

	return maxCoords, nil
}

func makeGrid(input int) [][]int {
	grid := make([][]int, size)

	for x := 0; x < size; x++ {
		grid[x] = make([]int, size)

		for y := 0; y < size; y++ {
			grid[x][y] = getPowerLevel(x, y, input)
		}
	}

	return grid
}

func calculateGridValue(grid [][]int, startX int, startY int, size int) int {
	value := 0

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			value += grid[startX-x][startY-y]
		}
	}

	return value
}

func getPowerLevel(x int, y int, serialNumber int) int {
	rackId := x + 10
	powerLevel := rackId * y
	powerLevel += serialNumber
	powerLevel *= rackId

	if powerLevel < 100 {
		powerLevel = 0
	} else {
		levelStr := strconv.Itoa(powerLevel)
		levelLen := len(levelStr)
		powerLevel, _ = strconv.Atoi(levelStr[levelLen-3 : levelLen-2])
	}

	return powerLevel - 5
}
