package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"github.com/mitchellh/hashstructure/v2"
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

type cell struct {
	Kind    rune // exported to make it hashable
	pending rune
	coords  string
	north   *cell
	east    *cell
	south   *cell
	west    *cell
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
				coords: fmt.Sprint(x, "x", y),
				Kind:   row,
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

	var last uint64

	for {
		for _, g := range grid {
			if g.Kind == floor {
				continue
			}

			counts := map[rune]int{
				empty:    0,
				floor:    0,
				occupied: 0,
			}

			for _, d := range directions {
				n := g

				for i := 0; i < los; i++ {
					n = n.neighbour(d)

					if n == nil {
						break
					}

					if n.Kind == floor {
						continue
					}

					counts[n.Kind]++
					break
				}
			}

			switch g.Kind {
			case empty:
				if counts[occupied] == 0 {
					g.pending = occupied
				}
			case occupied:
				if counts[occupied] >= tolerance {
					g.pending = empty
				}
			}
		}

		for _, g := range grid {
			if g.pending != 0 {
				g.Kind = g.pending
				g.pending = 0
			}
		}

		hash, _ := hashstructure.Hash(grid, hashstructure.FormatV2, nil)

		if hash == last {
			break
		}

		last = hash
	}

	result := 0

	for _, v := range grid {
		if v.Kind == occupied {
			result++
		}
	}

	return result
}

func printGrid(grid []*cell) {
	for _, g := range grid {
		fmt.Print(string(g.Kind))

		if g.neighbour(east) == nil {
			fmt.Println()
		}
	}

	fmt.Println()
}
