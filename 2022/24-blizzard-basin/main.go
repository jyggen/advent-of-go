package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"

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

const (
	north = iota
	east
	south
	west
)

type blizzard struct {
	at        [2]int
	direction int
}

func (b *blizzard) whereAtMinute(minute int, rowLength int, colLength int) (y int, x int) {
	switch b.direction {
	case north:
		y = (b.at[0] - minute) % (rowLength - 2)
		x = b.at[1]

		if y < 0 {
			y += rowLength - 2
		}

		y += 1
	case east:
		y = b.at[0]
		x = ((b.at[1] + minute) % (colLength - 2)) + 1
	case south:
		y = ((b.at[0] + minute) % (rowLength - 2)) + 1
		x = b.at[1]
	case west:
		y = b.at[0]
		x = (b.at[1] - minute) % (colLength - 2)

		if x < 0 {
			x += colLength - 2
		}

		x += 1
	}

	return y, x
}

type minute struct {
	minute int
	at     [2]int
}

type valley struct {
	start      [2]int
	goal       [2]int
	rowLength  int
	colLength  int
	yBlizzards [][]*blizzard
	xBlizzards [][]*blizzard
}

func parse(input string) *valley {
	rows := utils.ToRuneSlice(input, "\n")
	rowLength := len(rows)
	colLength := len(rows[0])
	yBlizzards := make([][]*blizzard, rowLength)
	xBlizzards := make([][]*blizzard, colLength)

	for i := range yBlizzards {
		yBlizzards[i] = make([]*blizzard, 0)
	}

	for i := range xBlizzards {
		xBlizzards[i] = make([]*blizzard, 0)
	}

	var start [2]int
	var goal [2]int

	for y, row := range rows {
		for x, r := range row {
			if y == 0 && r == '.' {
				start = [2]int{y, x}
			} else if y == rowLength-1 && r == '.' {
				goal = [2]int{y, x}
			} else {
				switch r {
				case '^':
					xBlizzards[x] = append(xBlizzards[x], &blizzard{[2]int{y - 1, x}, north})
				case '>':
					yBlizzards[y] = append(yBlizzards[y], &blizzard{[2]int{y, x - 1}, east})
				case 'v':
					xBlizzards[x] = append(xBlizzards[x], &blizzard{[2]int{y - 1, x}, south})
				case '<':
					yBlizzards[y] = append(yBlizzards[y], &blizzard{[2]int{y, x - 1}, west})
				}
			}
		}
	}

	return &valley{start, goal, rowLength, colLength, yBlizzards, xBlizzards}
}

func solve(startingMinute int, start [2]int, goal [2]int, v *valley) int {
	queue := []minute{{startingMinute, start}}
	been := make(map[[3]int]struct{}, v.rowLength*v.colLength)

	var q minute

	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]

		if q.at == goal {
			return q.minute
		}

		if _, ok := been[[3]int{q.at[0], q.at[1], q.minute}]; ok {
			continue
		}

		been[[3]int{q.at[0], q.at[1], q.minute}] = struct{}{}

	NeighbourLoop:
		for _, neighbour := range [5][2]int{
			{q.at[0], q.at[1]},
			{q.at[0] - 1, q.at[1]},
			{q.at[0], q.at[1] + 1},
			{q.at[0] + 1, q.at[1]},
			{q.at[0], q.at[1] - 1},
		} {
			if neighbour != goal && neighbour != start {
				if neighbour[0] <= 0 || neighbour[1] <= 0 || neighbour[0] >= v.rowLength-1 || neighbour[1] >= v.colLength-1 {
					continue
				}

				for _, b := range v.yBlizzards[neighbour[0]] {
					y, x := b.whereAtMinute(q.minute+1, v.rowLength, v.colLength)
					if y == neighbour[0] && x == neighbour[1] {
						continue NeighbourLoop
					}
				}

				for _, b := range v.xBlizzards[neighbour[1]] {
					y, x := b.whereAtMinute(q.minute+1, v.rowLength, v.colLength)
					if y == neighbour[0] && x == neighbour[1] {
						continue NeighbourLoop
					}
				}
			}

			queue = append(queue, minute{q.minute + 1, neighbour})
		}
	}

	return -1
}

func SolvePart1(input string) (string, error) {
	v := parse(input)

	return strconv.Itoa(solve(0, v.start, v.goal, v)), nil
}

func SolvePart2(input string) (string, error) {
	v := parse(input)
	m := 0
	start := v.start
	goal := v.goal

	for i := 0; i < 3; i++ {
		m = solve(m, start, goal, v)
		goal, start = start, goal
	}

	return strconv.Itoa(m), nil
}
