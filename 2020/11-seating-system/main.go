package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

const empty = 'L'
const floor = '.'
const occupied = '#'

var directions = [4]int{
	utils2.North,
	utils2.NorthEast,
	utils2.West,
	utils2.NorthWest,
}

var opposites = [4]int{
	utils2.South,
	utils2.SouthWest,
	utils2.East,
	utils2.SouthEast,
}

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(simulate(input, false, 4)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(simulate(input, true, 5)), nil
}

func simulate(input string, los bool, tolerance int) int {
	rows := utils2.ToRuneSlice(input, "\n")
	rowLen := len(rows)
	colLen := len(rows[0])
	gridSize := rowLen * colLen
	grid := make([]rune, gridSize)
	gridCopy := make([]rune, gridSize)
	neighbours := make([][8]int, gridSize)
	result := 0

	for y, cols := range rows {
		offset := colLen * y

		for x, row := range cols {
			pos := offset + x
			grid[pos] = row

			if row == occupied {
				result++
			}

			for d := range neighbours[pos] {
				neighbours[pos][d] = -1
			}

			for i, d := range directions {
				neighbour := utils2.Neighbour(x, y, d, rowLen, colLen)

				if los {
					for {
						if neighbour == -1 || grid[neighbour] != floor {
							break
						}

						neighbour = neighbours[neighbour][d]
					}
				}

				if neighbour != -1 {
					neighbours[pos][d] = neighbour
					neighbours[neighbour][opposites[i]] = pos
				}
			}
		}
	}

	copy(gridCopy, grid)

	for {
		current := result

		for i, g := range grid {
			if g == floor {
				continue
			}

			count := 0

			for _, j := range neighbours[i] {
				if j != -1 && grid[j] == occupied {
					count++
				}
			}

			if g == occupied && count >= tolerance {
				result--
				gridCopy[i] = empty
			} else if g == empty && count == 0 {
				result++
				gridCopy[i] = occupied
			}
		}

		if current == result {
			break
		}

		copy(grid, gridCopy)
	}

	return result
}
