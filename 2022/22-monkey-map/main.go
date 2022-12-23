package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/jyggen/advent-of-go/internal/solver"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type cellLink struct {
	to     *cell
	facing int
}

type cell struct {
	y        int
	x        int
	walkable bool
	north    *cellLink
	east     *cellLink
	west     *cellLink
	south    *cellLink
}

const (
	EAST = iota
	SOUTH
	WEST
	NORTH
)

func parse(input string) (*cell, [][]*cell, string) {
	parts := strings.Split(input, "\n\n")
	rows := strings.Split(parts[0], "\n")
	rowLength := len(rows)
	grid := make([][]*cell, rowLength)

	for y, r := range rows {
		grid[y] = make([]*cell, len(r))

		for x := range grid[y] {
			if r[x] != ' ' {
				grid[y][x] = &cell{}
			}
		}
	}

	for y, row := range rows {
		colLength := len(grid[y])

		for x, r := range row {
			if r == ' ' {
				continue
			}

			c := grid[y][x]
			c.x = x + 1
			c.y = y + 1

			if r == '.' {
				c.walkable = true
			}

			north, east, south, west := y, x, y, x

			for north == y || x >= len(rows[north]) || rows[north][x] == ' ' {
				north--

				if north == -1 {
					north = rowLength - 1
				}
			}

			for east == x || rows[y][east] == ' ' {
				east++

				if east == colLength {
					east = 0
				}
			}

			for south == y || x >= len(rows[south]) || rows[south][x] == ' ' {
				south++

				if south == rowLength {
					south = 0
				}
			}

			for west == x || rows[y][west] == ' ' {
				west--

				if west == -1 {
					west = colLength - 1
				}
			}

			c.north = &cellLink{grid[north][x], NORTH}
			c.east = &cellLink{grid[y][east], EAST}
			c.south = &cellLink{grid[south][x], SOUTH}
			c.west = &cellLink{grid[y][west], WEST}
		}
	}

	var position *cell

	for _, c := range grid[0] {
		if c != nil && c.walkable {
			position = c
			break
		}
	}

	return position, grid, strings.TrimSpace(parts[1])
}

func solve(instructions string, position *cell) int {
	number := make([]rune, 0)
	facing := EAST
	for _, c := range instructions {
		if unicode.IsDigit(c) {
			number = append(number, c)
		} else {
			position, facing = walk(position, facing, number)

			switch c {
			case 'L':
				facing--

				if facing == -1 {
					facing = 3
				}
			case 'R':
				facing = (facing + 1) % 4
			}

			number = make([]rune, 0)
		}
	}

	if len(number) > 0 {
		position, facing = walk(position, facing, number)
	}

	return (1000 * position.y) + (position.x * 4) + facing
}

func walk(position *cell, facing int, number []rune) (*cell, int) {
	steps, _ := strconv.Atoi(string(number))

WalkLoop:
	for steps > 0 {
		switch facing {
		case NORTH:
			if !position.north.to.walkable {
				break WalkLoop
			}

			position, facing = position.north.to, position.north.facing
		case EAST:
			if !position.east.to.walkable {
				break WalkLoop
			}

			position, facing = position.east.to, position.east.facing
		case SOUTH:
			if !position.south.to.walkable {
				break WalkLoop
			}

			position, facing = position.south.to, position.south.facing
		case WEST:
			if !position.west.to.walkable {
				break WalkLoop
			}

			position, facing = position.west.to, position.west.facing
		}

		steps--
	}

	return position, facing
}

func SolvePart1(input string) (string, error) {
	position, _, instructions := parse(input)

	return strconv.Itoa(solve(instructions, position)), nil
}

const cubeSize = 50

func SolvePart2(input string) (string, error) {
	position, grid, instructions := parse(input)
	currentCube := 0
	cubeNoLookup := make(map[int]int)
	cubes := [6][cubeSize][cubeSize]*cell{}

	for _, row := range grid {
		for _, c := range row {
			if c == nil {
				continue
			}

			cubeNumber := ((c.x - 1) / cubeSize) + (((c.y - 1) / cubeSize) * 10)

			if _, ok := cubeNoLookup[cubeNumber]; !ok {
				cubeNoLookup[cubeNumber] = currentCube
				currentCube++
			}

			cubes[cubeNoLookup[cubeNumber]][(c.y-1)%cubeSize][(c.x-1)%cubeSize] = c
		}
	}

	// @todo: hardcoded to my input, for now...
	for a, b := 0, cubeSize-1; a < cubeSize; a, b = a+1, b-1 {
		// 1 NORTH -> 6 WEST
		cubes[0][0][a].north.to = cubes[5][a][0]
		cubes[0][0][a].north.facing = EAST
		cubes[5][a][0].west.to = cubes[0][0][a]
		cubes[5][a][0].west.facing = SOUTH

		// 1 WEST -> 4 WEST
		cubes[0][a][0].west.to = cubes[3][b][0]
		cubes[0][a][0].west.facing = EAST
		cubes[3][b][0].west.to = cubes[0][a][0]
		cubes[3][b][0].west.facing = EAST

		// 2 NORTH -> 6 SOUTH
		cubes[1][0][a].north.to = cubes[5][cubeSize-1][a]
		cubes[1][0][a].north.facing = NORTH
		cubes[5][cubeSize-1][a].south.to = cubes[1][0][a]
		cubes[5][cubeSize-1][a].south.facing = SOUTH

		// 2 EAST -> 5 EAST
		cubes[1][a][cubeSize-1].east.to = cubes[4][b][cubeSize-1]
		cubes[1][a][cubeSize-1].east.facing = WEST
		cubes[4][b][cubeSize-1].east.to = cubes[1][a][cubeSize-1]
		cubes[4][b][cubeSize-1].east.facing = WEST

		// 2 SOUTH -> 3 EAST
		cubes[1][cubeSize-1][a].south.to = cubes[2][a][cubeSize-1]
		cubes[1][cubeSize-1][a].south.facing = WEST
		cubes[2][a][cubeSize-1].east.to = cubes[1][cubeSize-1][a]
		cubes[2][a][cubeSize-1].east.facing = NORTH

		// 3 WEST -> 4 NORTH
		cubes[2][a][0].west.to = cubes[3][0][a]
		cubes[2][a][0].west.facing = SOUTH
		cubes[3][0][a].north.to = cubes[2][a][0]
		cubes[3][0][a].north.facing = EAST

		// 5 SOUTH -> 6 EAST
		cubes[4][cubeSize-1][a].south.to = cubes[5][a][cubeSize-1]
		cubes[4][cubeSize-1][a].south.facing = WEST
		cubes[5][a][cubeSize-1].east.to = cubes[4][cubeSize-1][a]
		cubes[5][a][cubeSize-1].east.facing = NORTH
	}

	return strconv.Itoa(solve(instructions, position)), nil
}
