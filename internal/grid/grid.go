package grid

import (
	"github.com/beefsack/go-astar"
	"github.com/jyggen/advent-of-go/internal/utils"
)

func NewGrid(values [][]int, allowDiagonalNeighbours bool) *Grid {
	g := &Grid{
		cells:                   make([][]*Cell, len(values)),
		allowDiagonalNeighbours: allowDiagonalNeighbours,
	}

	for x, rows := range values {
		g.cells[x] = make([]*Cell, len(rows))

		for y, v := range rows {
			g.cells[x][y] = &Cell{
				x:     x,
				y:     y,
				grid:  g,
				Value: v,
			}
		}
	}

	return g
}

type Grid struct {
	cells                   [][]*Cell
	allowDiagonalNeighbours bool
}

func (g Grid) CellAt(x int, y int) *Cell {
	if x < 0 || x >= len(g.cells) {
		return nil
	}

	if y < 0 || y >= len(g.cells[x]) {
		return nil
	}

	return g.cells[x][y]
}

func (g Grid) Each(callback func(c *Cell) bool) {
	for _, cells := range g.cells {
		for _, c := range cells {
			shouldContinue := callback(c)

			if !shouldContinue {
				return
			}
		}
	}
}

func (g Grid) Size() int {
	return len(g.cells) * len(g.cells[0])
}

type Cell struct {
	x          int
	y          int
	grid       *Grid
	neighbours []*Cell
	Value      int
}

func (c Cell) Neighbours() []*Cell {
	if c.neighbours == nil {
		neighbours := make([]*Cell, 0, 8)
		coordsList := make([][2]int, 0, 8)
		coordsList = append(coordsList, [][2]int{
			{c.x, c.y - 1}, // N
			{c.x + 1, c.y}, // E
			{c.x, c.y + 1}, // S
			{c.x - 1, c.y}, // W
		}...)

		if c.grid.allowDiagonalNeighbours {
			coordsList = append(coordsList, [][2]int{
				{c.x + 1, c.y - 1}, // NE
				{c.x + 1, c.y + 1}, // SE
				{c.x - 1, c.y + 1}, // SW
				{c.x - 1, c.y - 1}, // NW
			}...)
		}

		for _, coords := range coordsList {
			n := c.grid.CellAt(coords[0], coords[1])

			if n == nil {
				continue
			}

			neighbours = append(neighbours, n)
		}

		c.neighbours = neighbours
	}

	return c.neighbours
}

func (c Cell) X() int {
	return c.x
}

func (c Cell) Y() int {
	return c.y
}

func (c Cell) PathNeighbors() []astar.Pather {
	neighbours := c.Neighbours()
	returnSlice := make([]astar.Pather, len(neighbours))

	for i, n := range neighbours {
		returnSlice[i] = n
	}

	return returnSlice
}

func (c Cell) PathNeighborCost(to astar.Pather) float64 {
	return float64(to.(*Cell).Value)
}

func (c Cell) PathEstimatedCost(to astar.Pather) float64 {
	return float64(utils.ManhattanDistance(c.X()+to.(*Cell).X(), c.Y()+to.(*Cell).Y()))
}
