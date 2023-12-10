package main

import (
	"fmt"
	"os"
	"slices"
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

func getHeading(r rune) int {
	switch r {
	case '|':
		return utils.North
	case '-':
		return utils.East
	case 'L':
		return utils.South
	case 'J':
		return utils.South
	case '7':
		return utils.East
	case 'F':
		return utils.North
	default:
		panic("unable to determine heading")
	}
}

func replaceSnakeAndGetHeading(maze []rune, at, numRows, numCols int) int {
	x, y := utils.ToCoordinates(at, numCols)
	directions := [2]int{-1, -1}

	for _, ns := range []struct {
		direction int
		accepted  []rune
	}{
		{utils.North, []rune{'|', '7', 'F'}},
		{utils.East, []rune{'-', 'J', '7'}},
		{utils.South, []rune{'|', 'L', 'J'}},
		{utils.West, []rune{'-', 'L', 'F'}},
	} {
		n := utils.Neighbour(x, y, ns.direction, numRows, numCols)

		if n == -1 {
			continue
		}

		if slices.Contains(ns.accepted, maze[n]) {
			if directions[0] == -1 {
				directions[0] = ns.direction
			} else {
				directions[1] = ns.direction

				break
			}
		}
	}

	switch directions {
	case [2]int{utils.North, utils.South}:
		maze[at] = '|'
	case [2]int{utils.East, utils.West}:
		maze[at] = '-'
	case [2]int{utils.North, utils.East}:
		maze[at] = 'L'
	case [2]int{utils.North, utils.West}:
		maze[at] = 'J'
	case [2]int{utils.South, utils.West}:
		maze[at] = '7'
	case [2]int{utils.East, utils.South}:
		maze[at] = 'F'
	default:
		panic("unable to determine snake pipe type")
	}

	return getHeading(maze[at])
}

func next(maze []rune, at int, heading, numRows, numCols int) (int, int) {
	x, y := utils.ToCoordinates(at, numCols)
	direction := -1

	switch maze[at] {
	case '|':
		if heading == utils.North {
			direction = utils.North
		} else {
			direction = utils.South
		}
	case '-':
		if heading == utils.East {
			direction = utils.East
		} else {
			direction = utils.West
		}
	case 'L':
		if heading == utils.South {
			direction = utils.East
		} else {
			direction = utils.North
		}
	case 'J':
		if heading == utils.South {
			direction = utils.West
		} else {
			direction = utils.North
		}
	case '7':
		if heading == utils.North {
			direction = utils.West
		} else {
			direction = utils.South
		}
	case 'F':
		if heading == utils.North {
			direction = utils.East
		} else {
			direction = utils.South
		}
	}

	return utils.Neighbour(x, y, direction, numRows, numCols), direction
}

func SolvePart1(input string) (string, error) {
	lines := utils.ToRuneSlice(input, "\n")
	numCols := len(lines[0])
	numRows := len(lines)
	maze := make([]rune, numCols*numRows)

	for i, l := range lines {
		for j, r := range l {
			maze[(i*numCols)+j] = r
		}
	}

	steps := 1
	start := slices.Index(maze, 'S')
	heading := replaceSnakeAndGetHeading(maze, start, numRows, numCols)
	at := start

	for {
		at, heading = next(maze, at, heading, numRows, numCols)

		if at == start {
			break
		}

		steps++
	}

	return strconv.Itoa(steps / 2), nil
}

func setNorthSouthOf(maze []rune, northOf bool, x, y, numRows, numCols int) {
	if northOf {
		n := utils.Neighbour(x, y, utils.North, numRows, numCols)

		if maze[n] != 'X' {
			maze[n] = '.'
		}
	} else {
		n := utils.Neighbour(x, y, utils.South, numRows, numCols)

		if maze[n] != 'X' {
			maze[n] = '.'
		}
	}
}

func setEastWestOf(maze []rune, eastOf bool, x, y, numRows, numCols int) {
	if eastOf {
		n := utils.Neighbour(x, y, utils.East, numRows, numCols)

		if maze[n] != 'X' {
			maze[n] = '.'
		}
	} else {
		n := utils.Neighbour(x, y, utils.West, numRows, numCols)

		if maze[n] != 'X' {
			maze[n] = '.'
		}
	}
}

func SolvePart2(input string) (string, error) {
	lines := utils.ToRuneSlice(input, "\n")
	numCols := len(lines[0])
	numRows := len(lines)
	maze := make([]rune, numCols*numRows)
	resolved := make([]rune, numCols*numRows)

	for i, l := range lines {
		for j, r := range l {
			maze[(i*numCols)+j] = r
			resolved[(i*numCols)+j] = ' '
		}
	}

	start := slices.Index(maze, 'S')
	heading := replaceSnakeAndGetHeading(maze, start, numRows, numCols)
	at := start
	corner := start
	resolved[at] = 'X'

	for {
		at, heading = next(maze, at, heading, numRows, numCols)
		resolved[at] = 'X'

		if at == start {
			break
		}

		corner = min(at, corner)
	}

	at = corner
	start = corner
	heading = getHeading(maze[at])
	northOf := false
	eastOf := true

	for {
		x, y := utils.ToCoordinates(at, numCols)

		switch maze[at] {
		case '|':
			setEastWestOf(resolved, eastOf, x, y, numRows, numCols)
		case '-':
			setNorthSouthOf(resolved, northOf, x, y, numRows, numCols)
		case 'L':
			if heading == utils.South {
				northOf = eastOf
			}

			if heading == utils.West {
				eastOf = northOf
			}

			setEastWestOf(resolved, eastOf, x, y, numRows, numCols)
			setNorthSouthOf(resolved, northOf, x, y, numRows, numCols)
		case 'J':
			if heading == utils.South {
				northOf = !eastOf
			}

			if heading == utils.East {
				eastOf = !northOf
			}

			setEastWestOf(resolved, eastOf, x, y, numRows, numCols)
			setNorthSouthOf(resolved, northOf, x, y, numRows, numCols)
		case '7':
			if heading == utils.East {
				eastOf = northOf
			}

			if heading == utils.North {
				northOf = eastOf
			}

			setEastWestOf(resolved, eastOf, x, y, numRows, numCols)
			setNorthSouthOf(resolved, northOf, x, y, numRows, numCols)
		case 'F':
			if heading == utils.North {
				northOf = !eastOf
			}

			if heading == utils.West {
				eastOf = !northOf
			}

			setEastWestOf(resolved, eastOf, x, y, numRows, numCols)
			setNorthSouthOf(resolved, northOf, x, y, numRows, numCols)
		}

		at, heading = next(maze, at, heading, numRows, numCols)

		if at == start {
			break
		}
	}

	for k, v := range resolved {
		if v != '.' {
			continue
		}

		x, y := utils.ToCoordinates(k, numCols)
		east := utils.Neighbour(x, y, utils.East, numRows, numCols)
		south := utils.Neighbour(x, y, utils.South, numRows, numCols)

		if east != -1 && resolved[east] == ' ' {
			resolved[east] = '.'
		}

		if south != -1 && resolved[south] == ' ' {
			resolved[south] = '.'
		}
	}

	count := 0

	for _, v := range resolved {
		if v == '.' {
			count++
		}
	}

	return strconv.Itoa(count), nil
}
