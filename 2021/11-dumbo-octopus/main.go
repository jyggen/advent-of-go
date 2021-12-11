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

func simulate(g *grid.Grid) int {
	//queue := list.New()
	flashes := 0
	queue := make([]*grid.Cell, 0, g.Size())

	g.Each(func(c *grid.Cell) bool {
		c.Value = c.Value + 1

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

			for _, n := range c.Neighbours() {
				queue = append(queue, n)
			}
		}

		queue[0] = queue[len(queue)-1]
		queue[len(queue)-1] = nil
		queue = queue[:len(queue)-1]
	}

	return flashes
}

func makeGrid(input string) (*grid.Grid, error) {
	stringSlice := utils.ToStringSlice(input, "\n")
	values := make([][]int, len(stringSlice))

	for i, s := range stringSlice {
		numbers, err := utils.ToIntegerSlice(s, "")

		if err != nil {
			return nil, err
		}

		values[i] = numbers
	}

	return grid.NewGrid(values, true), nil
}

func SolvePart1(input string) (string, error) {
	g, err := makeGrid(input)

	if err != nil {
		return "", err
	}

	flashes := 0

	for i := 0; i < 100; i++ {
		flashes += simulate(g)
	}

	return strconv.Itoa(flashes), nil
}

func SolvePart2(input string) (string, error) {
	g, err := makeGrid(input)

	if err != nil {
		return "", err
	}

	for i := 0; ; i++ {
		if simulate(g) == g.Size() {
			return strconv.Itoa(i + 1), nil
		}
	}

	return strconv.Itoa(0), nil
}
