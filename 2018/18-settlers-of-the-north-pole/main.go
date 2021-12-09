package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

const open = '.'
const tree = '|'
const lumberyard = '#'

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(simulate(input, 10)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(simulate(input, 1000000000)), nil
}

func simulate(input string, times int) int {
	rows := utils2.ToRuneSlice(input, "\n")

	rowLen := len(rows)
	colLen := len(rows[0])
	gridSize := rowLen * colLen
	grid := make([]rune, gridSize)

	for i, cols := range rows {
		offset := colLen * i

		for j, row := range cols {
			grid[offset+j] = row

		}
	}

	seen := make(map[string]int, 0)
	iteration := 0
	optimize := true

	for {
		iteration++

		newGrid := make([]rune, gridSize)

		copy(newGrid, grid)

		for i, _ := range grid {
			y := i / colLen
			x := i - (colLen * y)
			above, right, below, left := y != 0, x != colLen-1, y != rowLen-1, x != 0
			n, s := i-colLen, i+colLen
			neighbours := make([]int, 0)
			counts := map[rune]int{
				open:       0,
				tree:       0,
				lumberyard: 0,
			}

			if above {
				if left {
					neighbours = append(neighbours, n-1)
				}

				neighbours = append(neighbours, n)

				if right {
					neighbours = append(neighbours, n+1)
				}
			}

			if right {
				neighbours = append(neighbours, i+1)
			}

			if below {
				if right {
					neighbours = append(neighbours, s+1)
				}

				neighbours = append(neighbours, s)

				if left {
					neighbours = append(neighbours, s-1)
				}
			}

			if left {
				neighbours = append(neighbours, i-1)
			}

			for _, j := range neighbours {
				counts[grid[j]]++
			}

			switch grid[i] {
			case open:
				if counts[tree] >= 3 {
					newGrid[i] = tree
				}
			case tree:
				if counts[lumberyard] >= 3 {
					newGrid[i] = lumberyard
				}
			case lumberyard:
				if counts[lumberyard] < 1 || counts[tree] < 1 {
					newGrid[i] = open
				}
			}
		}

		copy(grid, newGrid)

		if iteration == times {
			break
		}

		if !optimize {
			continue
		}

		h := fmt.Sprintf("%q", grid)

		if v, ok := seen[h]; ok {
			diff := iteration - v

			for {
				iteration += diff

				if iteration > times {
					iteration -= diff
					optimize = false
					break
				}
			}
		}

		seen[h] = iteration
	}

	trees := 0
	lumberyards := 0

	for _, v := range grid {
		if v == tree {
			trees++
		} else if v == lumberyard {
			lumberyards++
		}
	}

	return trees * lumberyards
}
