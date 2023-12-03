package main

import (
	"fmt"
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

func simulate(g *grid.Grid[int], queue []*grid.Cell[int]) int {
	flashes := 0

	g.Each(func(c *grid.Cell[int]) bool {
		c.Value++

		if c.Value > 9 {
			queue = append(queue, c)
		}

		return true
	})

	for len(queue) > 0 {
		c := queue[0]

		if c.Value != 0 {
			c.Value = c.Value + 1
		}

		if c.Value > 9 {
			flashes++
			c.Value = 0
			queue = append(queue, c.Neighbours()...)
		}

		queue[0] = queue[len(queue)-1]
		queue[len(queue)-1] = nil
		queue = queue[:len(queue)-1]
	}

	return flashes
}

func makeGrid(input string) (*grid.Grid[int], error) {
	stringSlice := utils.ToStringSlice(input, "\n")
	values := make([][]int, len(stringSlice))

	var err error

	for i, s := range stringSlice {
		values[i], err = utils.ToIntegerSlice(s, "")

		if err != nil {
			return nil, err
		}
	}

	return grid.NewGrid(values, true), nil
}

func SolvePart1(input string) (string, error) {
	g, err := makeGrid(input)

	if err != nil {
		return "", err
	}

	flashes := 0
	queue := make([]*grid.Cell[int], 0, g.Size())

	for i := 0; i < 100; i++ {
		flashes += simulate(g, queue)
	}

	return strconv.Itoa(flashes), nil
}

func SolvePart2(input string) (string, error) {
	g, err := makeGrid(input)

	if err != nil {
		return "", err
	}

	queue := make([]*grid.Cell[int], 0, g.Size())

	for i := 0; ; i++ {
		if simulate(g, queue) == g.Size() {
			return strconv.Itoa(i + 1), nil
		}
	}

	return strconv.Itoa(0), nil
}
