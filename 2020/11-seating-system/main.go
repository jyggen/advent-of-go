package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"math"
	"os"
	"strconv"
)

const empty = 'L'
const floor = '.'
const occupied = '#'

type change struct {
	cell int
	kind rune
}

var directions = [8]int{
	utils.North,
	utils.NorthEast,
	utils.East,
	utils.SouthEast,
	utils.South,
	utils.SouthWest,
	utils.West,
	utils.NorthWest,
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(simulate(input, 1, 4)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(simulate(input, math.MaxUint8, 5)), nil
}

func simulate(input string, los int, tolerance int) int {
	rows := utils.ToRuneSlice(input, "\n")
	rowLen := len(rows)
	colLen := len(rows[0])
	gridSize := rowLen * colLen
	grid := make([]rune, gridSize)
	neighbours := make([][8]int, gridSize)

	for y, cols := range rows {
		offset := colLen * y

		for x, row := range cols {
			pos := offset + x
			grid[pos] = row

			for _, d := range directions {
				neighbours[pos][d] = utils.Neighbour(pos, d, rowLen, colLen)
			}
		}
	}

	for {
		pending := make([]*change, 0)

		for i, g := range grid {
			if g == floor {
				continue
			}

			count := 0

			for _, d := range directions {
				j := i

				for k := 0; k < los; k++ {
					j = neighbours[j][d]

					if j == -1 {
						break
					}

					if grid[j] == floor {
						continue
					}

					if grid[j] == occupied {
						count++
					}

					break
				}
			}

			switch g {
			case empty:
				if count == 0 {
					pending = append(pending, &change{cell: i, kind: occupied})
				}
			case occupied:
				if count >= tolerance {
					pending = append(pending, &change{cell: i, kind: empty})
				}
			}
		}

		if len(pending) == 0 {
			break
		}

		for _, p := range pending {
			grid[p.cell] = p.kind
		}
	}

	result := 0

	for _, v := range grid {
		if v == occupied {
			result++
		}
	}

	return result
}
