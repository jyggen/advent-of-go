package main

import (
	"errors"
	"fmt"
	"github.com/beefsack/go-astar"
	"github.com/jyggen/advent-of-go/internal/grid"
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
	rows := utils.ToStringSlice(input, "\n")
	data := make([][]int, len(rows))

	for i, r := range rows {
		ints, _ := utils.ToIntegerSlice(r, "")
		data[i] = ints
	}

	g := *grid.NewGrid(data, false)
	start := g.CellAt(0, 0)
	goal := g.CellAt(len(data)-1, len(data[0])-1)

	_, distance, found := astar.Path(start, goal)

	if !found {
		return "", errors.New("no path found")
	}

	return strconv.Itoa(int(distance)), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	data := make([][]int, len(rows)*5)

	for i, r := range rows {
		ints, _ := utils.ToIntegerSlice(r, "")

		for j := 0; j < 5; j++ {
			index := i + (len(rows) * j)
			for k := 0; k < 5; k++ {
				newInts := make([]int, len(ints))

				for l, v := range ints {
					newInts[l] = v + j + k

					if newInts[l] > 9 {
						newInts[l] -= 9
					}
				}

				data[index] = append(data[index], newInts...)
			}
		}
	}

	g := *grid.NewGrid(data, false)
	start := g.CellAt(0, 0)
	goal := g.CellAt(len(data)-1, len(data[0])-1)

	_, distance, found := astar.Path(start, goal)

	if !found {
		return "", errors.New("no path found")
	}

	return strconv.Itoa(int(distance)), nil
}
