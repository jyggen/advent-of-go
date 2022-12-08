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

type item struct {
	value   int
	visible bool
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	grid := make([][]item, len(rows))

	for i, row := range rows {
		cols, _ := utils.ToIntegerSlice(row, "")
		grid[i] = make([]item, len(cols))

		for j, col := range cols {
			grid[i][j] = item{col, false}
		}
	}

	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			visible := true

			for k := i - 1; k >= 0; k-- {
				if grid[i][j].value <= grid[k][j].value {
					visible = false
					break
				}
			}

			if visible {
				grid[i][j].visible = true
				sum++
				continue
			}

			visible = true

			for k := i + 1; k < len(grid); k++ {
				if grid[i][j].value <= grid[k][j].value {
					visible = false
					break
				}
			}

			if visible {
				grid[i][j].visible = true
				sum++
				continue
			}

			visible = true

			for k := j - 1; k >= 0; k-- {
				if grid[i][j].value <= grid[i][k].value {
					visible = false
					break
				}
			}

			if visible {
				grid[i][j].visible = true
				sum++
				continue
			}

			visible = true

			for k := j + 1; k < len(grid[i]); k++ {
				if grid[i][j].value <= grid[i][k].value {
					visible = false
					break
				}
			}

			if visible {
				grid[i][j].visible = true
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	grid := make([][]item, len(rows))

	for i, row := range rows {
		cols, _ := utils.ToIntegerSlice(row, "")
		grid[i] = make([]item, len(cols))

		for j, col := range cols {
			grid[i][j] = item{col, false}
		}
	}

	bestScore := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			currentScore := 1
			iterations := 0

			for k := i - 1; k >= 0; k-- {
				iterations++
				if grid[i][j].value <= grid[k][j].value {
					break
				}
			}

			currentScore *= iterations
			iterations = 0

			for k := i + 1; k < len(grid); k++ {
				iterations++
				if grid[i][j].value <= grid[k][j].value {
					break
				}
			}

			currentScore *= iterations
			iterations = 0

			for k := j - 1; k >= 0; k-- {
				iterations++
				if grid[i][j].value <= grid[i][k].value {
					break
				}
			}

			currentScore *= iterations
			iterations = 0

			for k := j + 1; k < len(grid[i]); k++ {
				iterations++
				if grid[i][j].value <= grid[i][k].value {
					break
				}
			}
			currentScore *= iterations

			if currentScore > bestScore {
				bestScore = currentScore
			}
		}
	}

	return strconv.Itoa(bestScore), nil
}
