package main

import (
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

const (
	up = iota
	right
	down
	left
)

type cell struct {
	kind    rune
	visited bool
	up      bool
	right   bool
	down    bool
	left    bool
}

type beam struct {
	y         int
	x         int
	direction int
}

func simulate(input [][]rune, beams []*beam) int {
	height := len(input)
	width := len(input[0])
	grid := make([][]*cell, len(input))
	visited := 0

	for y, row := range input {
		grid[y] = make([]*cell, len(row))

		for x, column := range row {
			grid[y][x] = &cell{kind: column, visited: false, up: false, right: false, down: false, left: false}
		}
	}

	for i := 0; i < len(beams); i++ {
		b := beams[i]

	BeamLoop:
		for {
			if !grid[b.y][b.x].visited {
				visited++
				grid[b.y][b.x].visited = true
			}

			switch b.direction {
			case up:
				if grid[b.y][b.x].up {
					break BeamLoop
				}

				grid[b.y][b.x].up = true
			case right:
				if grid[b.y][b.x].right {
					break BeamLoop
				}

				grid[b.y][b.x].right = true
			case down:
				if grid[b.y][b.x].down {
					break BeamLoop
				}

				grid[b.y][b.x].down = true
			case left:
				if grid[b.y][b.x].left {
					break BeamLoop
				}

				grid[b.y][b.x].left = true
			}

			switch grid[b.y][b.x].kind {
			case '/':
				switch b.direction {
				case up:
					b.direction = right
				case right:
					b.direction = up
				case down:
					b.direction = left
				case left:
					b.direction = down
				}
			case '\\':
				switch b.direction {
				case up:
					b.direction = left
				case right:
					b.direction = down
				case down:
					b.direction = right
				case left:
					b.direction = up
				}
			case '|':
				if b.direction == left || b.direction == right {
					b.direction = up
					beams = append(beams, &beam{y: b.y, x: b.x, direction: down})
				}
			case '-':
				if b.direction == up || b.direction == down {
					b.direction = right
					beams = append(beams, &beam{y: b.y, x: b.x, direction: left})
				}
			}

			switch b.direction {
			case up:
				b.y--
			case right:
				b.x++
			case down:
				b.y++
			case left:
				b.x--
			}

			if b.y < 0 || b.y >= height || b.x < 0 || b.x >= width {
				break BeamLoop
			}
		}
	}

	return visited
}

func SolvePart1(input string) (string, error) {
	grid := utils.ToRuneSlice(input, "\n")
	beams := []*beam{
		{x: 0, y: 0, direction: right},
	}

	return strconv.Itoa(simulate(grid, beams)), nil
}

func SolvePart2(input string) (string, error) {
	grid := utils.ToRuneSlice(input, "\n")
	height := len(grid)
	width := len(grid[0])
	best := 0
	possibilities := make([]*beam, 0)
	possibilities = append(
		possibilities,
		&beam{y: 0, x: 0, direction: down},
		&beam{y: 0, x: 0, direction: right},
		&beam{y: height - 1, x: 0, direction: up},
		&beam{y: height - 1, x: 0, direction: right},
		&beam{y: 0, x: width - 1, direction: down},
		&beam{y: 0, x: width - 1, direction: left},
		&beam{y: height - 1, x: width - 1, direction: up},
		&beam{y: height - 1, x: width - 1, direction: left},
	)

	for x := 1; x < width-1; x++ {
		possibilities = append(possibilities, &beam{y: 0, x: x, direction: down}, &beam{y: height - 1, x: x, direction: up})
	}

	for y := 1; y < height-1; y++ {
		possibilities = append(possibilities, &beam{y: y, x: 0, direction: right}, &beam{y: y, x: width - 1, direction: left})
	}

	for _, p := range possibilities {
		beams := []*beam{p}
		best = max(best, simulate(grid, beams))
	}

	return strconv.Itoa(best), nil
}
