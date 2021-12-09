package main

import (
	"container/list"
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"sort"
	"strconv"
)

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

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

func makeGrid(input string) (*grid, error) {
	runeSlice := utils2.ToRuneSlice(input, "\n")
	colLength := len(runeSlice[0])
	rowLength := len(runeSlice)
	g := &grid{
		colLength: colLength,
		grid:      make([]int, len(runeSlice)*colLength),
		rowLength: rowLength,
	}

	var err error

	for i, runes := range runeSlice {
		for j, r := range runes {
			g.grid[(i*colLength)+j], err = strconv.Atoi(string(r))

			if err != nil {
				return g, err
			}
		}
	}

	return g, nil
}

func (g *grid) lowPoints() []int {
	lowPoints := make([]int, 0)

	for k, v := range g.grid {
		x, y := utils2.ToCoordinates(k, g.colLength)
		lower := true

		for _, n := range []int{
			utils2.Neighbour(x, y, utils2.North, g.rowLength, g.colLength),
			utils2.Neighbour(x, y, utils2.East, g.rowLength, g.colLength),
			utils2.Neighbour(x, y, utils2.South, g.rowLength, g.colLength),
			utils2.Neighbour(x, y, utils2.West, g.rowLength, g.colLength),
		} {
			if n == -1 {
				continue
			}

			if v >= g.grid[n] {
				lower = false
				break
			}
		}

		if lower {
			lowPoints = append(lowPoints, k)
		}
	}

	return lowPoints
}

func SolvePart1(input string) (string, error) {
	g, err := makeGrid(input)

	if err != nil {
		return "", nil
	}

	sum := 0

	for _, p := range g.lowPoints() {
		sum += g.grid[p] + 1
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	g, err := makeGrid(input)

	if err != nil {
		return "", nil
	}

	basins := make([]int, 0)

	for _, p := range g.lowPoints() {
		queue := list.New()
		cache := make(map[int]struct{}, 0)
		queue.PushBack(p)

		for queue.Len() > 0 {
			e := queue.Front()
			v := e.Value.(int)

			if _, ok := cache[v]; !ok {
				x, y := utils2.ToCoordinates(v, g.colLength)
				for _, n := range []int{
					utils2.Neighbour(x, y, utils2.North, g.rowLength, g.colLength),
					utils2.Neighbour(x, y, utils2.East, g.rowLength, g.colLength),
					utils2.Neighbour(x, y, utils2.South, g.rowLength, g.colLength),
					utils2.Neighbour(x, y, utils2.West, g.rowLength, g.colLength),
				} {
					if n != -1 && g.grid[n] != 9 && g.grid[n] > g.grid[v] {
						queue.PushBack(n)
					}
				}
			}

			cache[v] = struct{}{}
			queue.Remove(e)
		}

		basins = append(basins, len(cache))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return strconv.Itoa(basins[0] * basins[1] * basins[2]), nil
}
