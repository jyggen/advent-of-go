package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/utils"
	"math"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func solve(grid []rune, rowLength int, colLength int, end int, queue [][2]int) int {
	been := make(map[int]int, 0)
	best := math.MaxInt

	var current [2]int

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if current[0] == end {
			if current[1] < best {
				best = current[1]
			}
			continue
		}

		if current[1] > best {
			continue
		}

		if v, ok := been[current[0]]; ok && v < current[1] {
			continue
		}

		been[current[0]] = been[current[1]]
		x, y := utils.ToCoordinates(current[0], colLength)

		for _, d := range []int{
			utils.North,
			utils.East,
			utils.South,
			utils.West,
		} {
			n := utils.Neighbour(x, y, d, rowLength, colLength)

			if n == -1 {
				continue
			}

			if grid[n] <= grid[current[0]]+1 {
				queue = append(queue, [2]int{n, current[1] + 1})
			}
		}
	}

	return best
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToRuneSlice(input, "\n")
	rowLength := len(rows)
	colLength := len(rows[0])
	grid := make([]rune, 0, rowLength*colLength)

	for _, r := range rows {
		grid = append(grid, r...)
	}

	start := -1
	end := -1

	for i, g := range grid {
		if g == 'S' {
			start = i
			grid[i] = 'a'
		}

		if g == 'E' {
			end = i
			grid[i] = 'z'
		}

		if start != -1 && end != -1 {
			break
		}
	}

	best := solve(grid, rowLength, colLength, end, [][2]int{{start, 0}})

	return strconv.Itoa(best), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToRuneSlice(input, "\n")
	rowLength := len(rows)
	colLength := len(rows[0])
	grid := make([]rune, 0, rowLength*colLength)

	for _, r := range rows {
		grid = append(grid, r...)
	}

	end := -1
	queue := make([][2]int, 0)

	for i, g := range grid {
		if g == 'S' {
			queue = append(queue, [2]int{i, 0})
			grid[i] = 'a'
		} else if g == 'a' {
			queue = append(queue, [2]int{i, 0})
		} else if g == 'E' {
			end = i
			grid[i] = 'z'
		}
	}

	best := solve(grid, rowLength, colLength, end, queue)

	return strconv.Itoa(best), nil
}
