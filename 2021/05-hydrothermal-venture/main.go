package main

import (
	"fmt"
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

type line struct {
	x pair
	y pair
}

func (l line) IsStraight() bool {
	return l.x.IsStraight() || l.y.IsStraight()
}

type pair struct {
	from int
	to   int
}

func (p pair) Check(v int) bool {
	if p.from < p.to {
		return v <= p.to
	}

	return v >= p.to
}

func (p pair) Next(v int) int {
	if p.from < p.to {
		return v + 1
	}

	return v - 1
}

func (p pair) IsStraight() bool {
	return p.to == p.from
}

func SolvePart1(input string) (string, error) {
	grid, lines, err := makeGrid(input)
	if err != nil {
		return "", err
	}

	sum := 0

	for _, l := range lines {
		if !l.IsStraight() {
			continue
		}

		for x := l.x.from; l.x.Check(x); x = l.x.Next(x) {
			for y := l.y.from; l.y.Check(y); y = l.y.Next(y) {
				grid[y][x]++

				if grid[y][x] == 2 {
					sum++
				}
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	grid, lines, err := makeGrid(input)
	if err != nil {
		return "", err
	}

	sum := 0

	for _, l := range lines {
		if !l.IsStraight() {
			for x, y := l.x.from, l.y.from; l.x.Check(x) && l.y.Check(y); x, y = l.x.Next(x), l.y.Next(y) {
				grid[y][x]++

				if grid[y][x] == 2 {
					sum++
				}
			}
		} else {
			for x := l.x.from; l.x.Check(x); x = l.x.Next(x) {
				for y := l.y.from; l.y.Check(y); y = l.y.Next(y) {
					grid[y][x]++

					if grid[y][x] == 2 {
						sum++
					}
				}
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func makeGrid(input string) ([][]int, []line, error) {
	integers := utils.ToOptimisticIntSlice(input)
	lines := make([]line, 0, len(integers)/4)
	maxX := 0
	maxY := 0

	for i, j, k, l := 0, 1, 2, 3; l < len(integers); i, j, k, l = i+4, j+4, k+4, l+4 {
		x1, y1, x2, y2 := integers[i], integers[j], integers[k], integers[l]

		if x1 > maxX {
			maxX = x1
		}

		if x2 > maxX {
			maxX = x2
		}

		if y1 > maxY {
			maxY = y1
		}

		if y2 > maxY {
			maxY = y2
		}

		lines = append(lines, line{pair{from: x1, to: x2}, pair{from: y1, to: y2}})
	}

	grid := make([][]int, maxY+1)

	for y := range grid {
		grid[y] = make([]int, maxX+1)
	}

	return grid, lines, nil
}
