package main

import (
	"container/list"
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

type grid struct {
	colLength int
	grid      []int
	rowLength int
}

func (g *grid) simulate() int {
	queue := list.New()
	flashes := 0

	for j := 0; j < len(g.grid); j++ {
		g.grid[j]++

		if g.grid[j] > 9 {
			queue.PushBack(j)
		}
	}

	for queue.Len() > 0 {
		e := queue.Front()
		j := e.Value.(int)

		if g.grid[j] != 0 {
			g.grid[j]++
		}

		if g.grid[j] > 9 {
			flashes++

			g.grid[j] = 0
			x, y := utils.ToCoordinates(j, g.colLength)

			for _, k := range []int{
				utils.Neighbour(x, y, utils.North, g.rowLength, g.colLength),
				utils.Neighbour(x, y, utils.NorthEast, g.rowLength, g.colLength),
				utils.Neighbour(x, y, utils.East, g.rowLength, g.colLength),
				utils.Neighbour(x, y, utils.SouthEast, g.rowLength, g.colLength),
				utils.Neighbour(x, y, utils.South, g.rowLength, g.colLength),
				utils.Neighbour(x, y, utils.SouthWest, g.rowLength, g.colLength),
				utils.Neighbour(x, y, utils.West, g.rowLength, g.colLength),
				utils.Neighbour(x, y, utils.NorthWest, g.rowLength, g.colLength),
			} {
				if k != -1 {
					queue.PushBack(k)
				}
			}
		}

		queue.Remove(e)
	}

	return flashes
}

func makeGrid(input string) (*grid, error) {
	stringSlice := utils.ToStringSlice(input, "\n")
	colLength := len(stringSlice[0])
	rowLength := len(stringSlice)
	g := &grid{
		colLength: colLength,
		grid:      make([]int, len(stringSlice)*colLength),
		rowLength: rowLength,
	}

	for i, s := range stringSlice {
		ints, err := utils.ToIntegerSlice(s, "")

		if err != nil {
			return g, err
		}

		for j, digit := range ints {
			g.grid[(i*colLength)+j] = digit
		}
	}

	return g, nil
}

func SolvePart1(input string) (string, error) {
	g, err := makeGrid(input)

	if err != nil {
		return "", err
	}

	flashes := 0

	for i := 0; i < 100; i++ {
		flashes += g.simulate()
	}

	return strconv.Itoa(flashes), nil
}

func SolvePart2(input string) (string, error) {
	g, err := makeGrid(input)

	if err != nil {
		return "", err
	}

	for i := 0; ; i++ {
		g.simulate()

		allFlashed := true

		for _, v := range g.grid {
			if v != 0 {
				allFlashed = false
				break
			}
		}

		if allFlashed {
			return strconv.Itoa(i + 1), nil
		}
	}

	return strconv.Itoa(0), nil
}
