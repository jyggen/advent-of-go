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
			grid[i][j] = item{value: col}
		}
	}

	sum := 0

	for r := 0; r < 4; r++ {
		for i := 0; i < len(grid); i++ {
			highest := -1
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j].value <= highest {
					continue
				}

				highest = grid[i][j].value

				if !grid[i][j].visible {
					grid[i][j].visible = true
					sum++
				}
			}
		}

		grid = rotateMatrix(grid)
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
			grid[i][j] = item{value: col, visible: false}
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

func rotateMatrix(matrix [][]item) [][]item {

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}
