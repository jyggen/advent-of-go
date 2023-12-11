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

func solve(input string, increase int) int {
	lines := utils.ToRuneSlice(input, "\n")
	numRows := len(lines)
	numCols := len(lines[0])
	universe := make([]rune, numRows*numCols)

	for y, l := range lines {
		for x, r := range l {
			universe[(y*numCols)+x] = r
		}
	}

	emptyCols := make([]int, 0)
	emptyRows := make([]int, 0)

	for x := 0; x < numCols; x++ {
		allEmpty := true

		for y := 0; y < numRows; y++ {
			if universe[(numCols*y)+x] != '.' {
				allEmpty = false

				break
			}
		}

		if !allEmpty {
			continue
		}

		emptyCols = append(emptyCols, x)
	}

	for y := 0; y < numRows; y++ {
		allEmpty := true

		for x := 0; x < numCols; x++ {
			if universe[(numCols*y)+x] != '.' {
				allEmpty = false

				break
			}
		}

		if !allEmpty {
			continue
		}

		emptyRows = append(emptyRows, y)
	}

	sum := 0

	for k, v := range universe {
		if v != '#' {
			continue
		}

		x, y := utils.ToCoordinates(k, numCols)

		for k2, v2 := range universe[k+1:] {
			if v2 != '#' {
				continue
			}

			x1, y1 := x, y
			x2, y2 := utils.ToCoordinates(k+1+k2, numCols)

			for _, col := range emptyCols {
				if col <= max(x, x2) && col >= min(x, x2) {
					if x > x2 {
						x1 += increase
					} else {
						x1 -= increase
					}
				}
			}

			for _, row := range emptyRows {
				if row <= max(y, y2) && row >= min(y, y2) {
					if y > y2 {
						y1 += increase
					} else {
						y1 -= increase
					}
				}
			}

			sum += utils.ManhattanDistance(x1-x2, y1-y2)
		}
	}

	return sum
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(solve(input, 1)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(solve(input, 999999)), nil
}
