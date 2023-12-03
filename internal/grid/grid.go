package grid

func NewGrid[V any](values [][]V, allowDiagonalNeighbours bool) *Grid[V] {
	g := &Grid[V]{
		colLength:               len(values[0]),
		rowLength:               len(values),
		cells:                   make([][]*Cell[V], len(values)),
		allowDiagonalNeighbours: allowDiagonalNeighbours,
	}

	for x, rows := range values {
		g.cells[x] = make([]*Cell[V], len(rows))

		for y, v := range rows {
			g.cells[x][y] = &Cell[V]{
				x:     x,
				y:     y,
				grid:  g,
				Value: v,
			}
		}
	}

	return g
}

type Grid[V any] struct {
	colLength               int
	rowLength               int
	cells                   [][]*Cell[V]
	allowDiagonalNeighbours bool
}

func (g Grid[V]) CellAt(x int, y int) *Cell[V] {
	if x < 0 || x >= len(g.cells) {
		return nil
	}

	if y < 0 || y >= len(g.cells[x]) {
		return nil
	}

	return g.cells[x][y]
}

func (g Grid[V]) CellAtTopLeft() *Cell[V] {
	return g.cells[0][0]
}

func (g Grid[V]) CellAtBottomRight() *Cell[V] {
	return g.cells[g.rowLength-1][g.colLength-1]
}

func (g Grid[V]) Each(callback func(c *Cell[V]) bool) {
	for _, cells := range g.cells {
		for _, c := range cells {
			shouldContinue := callback(c)

			if !shouldContinue {
				return
			}
		}
	}
}

func (g Grid[V]) Size() int {
	return len(g.cells) * len(g.cells[0])
}

type Cell[V any] struct {
	x          int
	y          int
	grid       *Grid[V]
	neighbours []*Cell[V]
	Value      V
}

func (c Cell[V]) Neighbours() []*Cell[V] {
	if c.neighbours == nil {
		neighbours := make([]*Cell[V], 0, 8)
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

func (c Cell[V]) ID() int {
	return c.x + (c.y * c.grid.colLength)
}

func (c Cell[V]) X() int {
	return c.x
}

func (c Cell[V]) Y() int {
	return c.y
}
