package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/grid"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"slices"
	"strconv"
	"strings"
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
	rows := utils.ToRuneSlice(input, "\n")
	g := grid.NewGrid(rows, true)
	sum := 0
	var partNumber strings.Builder
	var hasSymbolNeighbour bool

	g.Each(func(c *grid.Cell[rune]) bool {
		if c.Value >= '0' && c.Value <= '9' {
			partNumber.WriteRune(c.Value)

			for _, n := range c.Neighbours() {
				if (n.Value < '0' || n.Value > '9') && n.Value != '.' {
					hasSymbolNeighbour = true
					break
				}
			}
		} else {
			if partNumber.Len() > 0 {
				if hasSymbolNeighbour {
					value, _ := strconv.Atoi(partNumber.String())
					sum += value
				}

				partNumber = strings.Builder{}
				hasSymbolNeighbour = false
			}
		}

		return true
	})

	if partNumber.Len() > 0 {
		if hasSymbolNeighbour {
			value, _ := strconv.Atoi(partNumber.String())
			sum += value
		}

		partNumber = strings.Builder{}
		hasSymbolNeighbour = false
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToRuneSlice(input, "\n")
	g := grid.NewGrid(rows, true)
	gears := make(map[*grid.Cell[rune]][]int)
	neighbours := make([]*grid.Cell[rune], 0)
	sum := 0
	var partNumber strings.Builder

	g.Each(func(c *grid.Cell[rune]) bool {
		if c.Value >= '0' && c.Value <= '9' {
			partNumber.WriteRune(c.Value)

			for _, n := range c.Neighbours() {
				if n.Value == '*' && !slices.Contains(neighbours, n) {
					neighbours = append(neighbours, n)
				}
			}
		} else {
			if partNumber.Len() > 0 {
				if len(neighbours) > 0 {
					value, _ := strconv.Atoi(partNumber.String())

					for _, n := range neighbours {
						if _, ok := gears[n]; !ok {
							gears[n] = make([]int, 0)
						}

						gears[n] = append(gears[n], value)
					}

				}

				partNumber = strings.Builder{}
				neighbours = make([]*grid.Cell[rune], 0)
			}
		}

		return true
	})

	if partNumber.Len() > 0 {
		if len(neighbours) > 0 {
			value, _ := strconv.Atoi(partNumber.String())

			for _, n := range neighbours {
				if _, ok := gears[n]; !ok {
					gears[n] = make([]int, 0)
				}

				gears[n] = append(gears[n], value)
			}

		}

		partNumber = strings.Builder{}
		neighbours = make([]*grid.Cell[rune], 0)
	}

	for _, ints := range gears {
		if len(ints) == 2 {
			sum += ints[0] * ints[1]
		}
	}

	return strconv.Itoa(sum), nil
}
