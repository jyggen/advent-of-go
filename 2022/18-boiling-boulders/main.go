package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	coordinates := utils.ToOptimisticIntSlice(input, false)
	size := 0

	for _, c := range coordinates {
		size = utils.MaxInt(c, size)
	}

	size += 1
	grid := make([][][]bool, size)

	for x := range grid {
		grid[x] = make([][]bool, size)

		for y := range grid[x] {
			grid[x][y] = make([]bool, size)
		}
	}

	for i := 0; i < len(coordinates); i += 3 {
		grid[coordinates[i]][coordinates[i+1]][coordinates[i+2]] = true
	}

	sides := 0

	for i := 0; i < len(coordinates); i += 3 {
		x, y, z := coordinates[i], coordinates[i+1], coordinates[i+2]

		for _, neighbour := range [][3]int{
			{x - 1, y, z},
			{x + 1, y, z},
			{x, y - 1, z},
			{x, y + 1, z},
			{x, y, z - 1},
			{x, y, z + 1},
		} {
			x2, y2, z2 := neighbour[0], neighbour[1], neighbour[2]

			if x2 == -1 || x2 == size || y2 == -1 || y2 == size || z2 == -1 || z2 == size || !grid[x2][y2][z2] {
				sides++
			}
		}
	}

	return strconv.Itoa(sides), nil
}

func SolvePart2(input string) (string, error) {
	coordinates := utils.ToOptimisticIntSlice(input, false)
	size := 0

	for _, c := range coordinates {
		size = utils.MaxInt(c, size)
	}

	size += 1
	grid := make([][][]bool, size)

	for x := range grid {
		grid[x] = make([][]bool, size)

		for y := range grid[x] {
			grid[x][y] = make([]bool, size)
		}
	}

	for i := 0; i < len(coordinates); i += 3 {
		grid[coordinates[i]][coordinates[i+1]][coordinates[i+2]] = true
	}

	sides := 0
	knownReachable := make(map[[3]int]struct{})
	knownNotReachable := make(map[[3]int]struct{})

	for i := 0; i < len(coordinates); i += 3 {
		x, y, z := coordinates[i], coordinates[i+1], coordinates[i+2]

		for _, neighbour := range [][3]int{
			{x - 1, y, z},
			{x + 1, y, z},
			{x, y - 1, z},
			{x, y + 1, z},
			{x, y, z - 1},
			{x, y, z + 1},
		} {
			x2, y2, z2 := neighbour[0], neighbour[1], neighbour[2]

			if x2 == -1 || x2 == size || y2 == -1 || y2 == size || z2 == -1 || z2 == size {
				sides++
			} else if !grid[x2][y2][z2] {
				reachable, been := isReachable(grid, x2, y2, z2, size, knownReachable, knownNotReachable)

				if reachable {
					sides++
					for k, v := range been {
						knownReachable[k] = v
					}
				} else {
					for k, v := range been {
						knownNotReachable[k] = v
					}
				}
			}
		}
	}

	return strconv.Itoa(sides), nil
}

func isReachable(grid [][][]bool, x int, y int, z int, size int, knownReachable map[[3]int]struct{}, knownNotReachable map[[3]int]struct{}) (bool, map[[3]int]struct{}) {
	been := make(map[[3]int]struct{}, 0)
	queue := [][3]int{
		{x, y, z},
	}

	var current [3]int

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if _, ok := knownReachable[current]; ok {
			return true, been
		}

		if _, ok := knownNotReachable[current]; ok {
			return false, been
		}

		if _, ok := been[current]; ok {
			continue
		}

		been[current] = struct{}{}
		x2, y2, z2 := current[0], current[1], current[2]

		for _, neighbour := range [][3]int{
			{x2 - 1, y2, z2},
			{x2 + 1, y2, z2},
			{x2, y2 - 1, z2},
			{x2, y2 + 1, z2},
			{x2, y2, z2 - 1},
			{x2, y2, z2 + 1},
		} {
			x3, y3, z3 := neighbour[0], neighbour[1], neighbour[2]

			if x3 == -1 || x3 == size || y3 == -1 || y3 == size || z3 == -1 || z3 == size {
				return true, been
			}

			if !grid[x3][y3][z3] {
				queue = append(queue, neighbour)
			}
		}
	}

	return false, been
}
