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

const (
	north     = iota
	northEast = iota
	east      = iota
	southEast = iota
	south     = iota
	southWest = iota
	west      = iota
	northWest = iota
)

type change struct {
	cell *cell
	kind rune
}

type cell struct {
	id    int
	kind  rune
	north *cell
	east  *cell
	south *cell
	west  *cell
}

var directions [8]int

func init() {
	directions = [8]int{north, northEast, east, southEast, south, southWest, west, northWest}
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

func NewGrid(input [][]rune) []*cell {
	rowLen := len(input)
	colLen := len(input[0])
	gridSize := rowLen * colLen
	grid := make([]*cell, gridSize)

	for y, cols := range input {
		offset := colLen * y

		for x, row := range cols {
			pos := offset + x
			n, w := pos-colLen, pos-1
			grid[pos] = &cell{
				id:   pos,
				kind: row,
			}

			if y != 0 {
				grid[pos].north = grid[n]
				grid[n].south = grid[pos]
			}

			if x != 0 {
				grid[pos].west = grid[w]
				grid[w].east = grid[pos]
			}
		}
	}

	return grid
}

func (c *cell) neighbour(direction int) *cell {
	switch direction {
	case north:
		return c.north
	case northEast:
		n := c.north

		if n == nil {
			return nil
		}

		return n.east
	case east:
		return c.east
	case southEast:
		s := c.south

		if s == nil {
			return nil
		}

		return s.east
	case south:
		return c.south
	case southWest:
		s := c.south

		if s == nil {
			return nil
		}

		return s.west
	case west:
		return c.west
	case northWest:
		n := c.north

		if n == nil {
			return nil
		}

		return n.west
	default:
		return nil
	}
}

func simulate(input string, los int, tolerance int) int {
	rows := utils.ToRuneSlice(input, "\n")
	grid := NewGrid(rows)

	for {
		pending := make([]*change, 0)

		for _, g := range grid {
			if g.kind == floor {
				continue
			}

			count := 0

			for _, d := range directions {
				n := g

				for i := 0; i < los; i++ {
					n = n.neighbour(d)

					if n == nil {
						break
					}

					if n.kind == floor {
						continue
					}

					if n.kind == occupied {
						count++
					}

					break
				}
			}

			switch g.kind {
			case empty:
				if count == 0 {
					pending = append(pending, &change{cell: g, kind: occupied})
				}
			case occupied:
				if count >= tolerance {
					pending = append(pending, &change{cell: g, kind: empty})
				}
			}
		}

		if len(pending) == 0 {
			break
		}

		for _, p := range pending {
			p.cell.kind = p.kind
		}
	}

	result := 0

	for _, v := range grid {
		if v.kind == occupied {
			result++
		}
	}

	return result
}
