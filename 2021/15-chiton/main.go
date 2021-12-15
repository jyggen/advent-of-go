package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/jyggen/advent-of-go/internal/grid"
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

func createGridAndGraph(input string, multiplier int) (grid.Grid, *dijkstra.Graph) {
	rows := utils.ToStringSlice(input, "\n")
	colLength := len(rows[0])
	rowLength := len(rows)
	data := make([][]int, rowLength*multiplier)
	graph := dijkstra.NewGraph()

	for i, r := range rows {
		ints, _ := utils.ToIntegerSlice(r, "")

		for j := 0; j < multiplier; j++ {
			index := i + (rowLength * j)
			newInts := make([]int, colLength*multiplier)

			for k := 0; k < multiplier; k++ {
				for l, v := range ints {
					intIdx := l + (colLength * k)
					newInts[intIdx] = v + j + k

					if newInts[intIdx] > 9 {
						newInts[intIdx] -= 9
					}

					graph.AddVertex(intIdx + ((rowLength * multiplier) * index))
				}
			}

			data[index] = newInts
		}
	}

	g := *grid.NewGrid(data, false)

	g.Each(func(c *grid.Cell) bool {
		ourId := c.Y() + (c.X() * colLength * multiplier)
		for _, n := range c.Neighbours() {
			theirId := n.Y() + (n.X() * colLength * multiplier)
			graph.AddArc(ourId, theirId, int64(n.Value))
		}

		return true
	})

	return g, graph
}

func SolvePart1(input string) (string, error) {
	g, graph := createGridAndGraph(input, 1)
	best, err := graph.Shortest(g.CellAtTopLeft().ID(), g.CellAtBottomRight().ID())

	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(best.Distance)), nil
}

func SolvePart2(input string) (string, error) {
	g, graph := createGridAndGraph(input, 5)
	best, err := graph.Shortest(g.CellAtTopLeft().ID(), g.CellAtBottomRight().ID())

	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(best.Distance)), nil
}
