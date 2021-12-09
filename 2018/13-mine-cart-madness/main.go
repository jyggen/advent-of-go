package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
)

type cart struct {
	crashed      bool
	facing       int
	position     int
	intersection int
}

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

const (
	turnLeft   = 0
	goStraight = 1
	turnRight  = 2
)

var cartUp, cartRight, cartDown, cartLeft, intersection, turnOne, turnTwo, vertical, horizontal rune

func init() {
	cartUp = []rune("^")[0]
	cartRight = []rune(">")[0]
	cartDown = []rune("v")[0]
	cartLeft = []rune("<")[0]
	intersection = []rune("+")[0]
	turnOne = []rune("/")[0]
	turnTwo = []rune("\\")[0]
	vertical = []rune("|")[0]
	horizontal = []rune("-")[0]
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
	rows := strings.Split(input, "\n")
	rowLen := 0

	for _, r := range rows {
		currLen := len(r)

		if currLen > rowLen {
			rowLen = currLen
		}
	}

	tracks, carts := makeTracksAndCarts(rows, rowLen)
	firstCrashAt := 0

TickLoop:
	for {
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].position < carts[j].position
		})

		for i, c1 := range carts {
			c1.Move(tracks, rowLen)

			for j, c2 := range carts {
				if i != j && c1.position == c2.position {
					firstCrashAt = c1.position
					break TickLoop
				}
			}
		}
	}

	return positionToCoords(firstCrashAt, rowLen), nil
}

func SolvePart2(input string) (string, error) {
	rows := strings.Split(input, "\n")
	rowLen := 0

	for _, r := range rows {
		currLen := len(r)

		if currLen > rowLen {
			rowLen = currLen
		}
	}

	tracks, carts := makeTracksAndCarts(rows, rowLen)

TickLoop:
	for {
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].position < carts[j].position
		})

		crashedCarts := make([]int, 0)

		for i, c1 := range carts {
			if c1.crashed {
				continue
			}

			c1.Move(tracks, rowLen)

			for j, c2 := range carts {
				if i != j && !c2.crashed && c1.position == c2.position {
					c1.crashed = true
					c2.crashed = true
					crashedCarts = append(crashedCarts, i, j)
				}
			}
		}

		sort.Slice(crashedCarts, func(i, j int) bool {
			return crashedCarts[i] > crashedCarts[j]
		})

		for _, i := range crashedCarts {
			carts = append(carts[:i], carts[i+1:]...)
		}

		if len(carts) == 1 {
			break TickLoop
		}
	}

	return positionToCoords(carts[0].position, rowLen), nil
}

func makeTracksAndCarts(input []string, rowLen int) ([]rune, []*cart) {
	tracks := make([]rune, rowLen*len(input))
	carts := make([]*cart, 0)

	for i, r := range input {
		parts := strings.Split(padString(r, rowLen), "")
		for j, p := range parts {
			state := []rune(p)[0]
			position := i*rowLen + j
			tracks[position] = state

			if isCart(state) {
				facing := getCartDirection(state)
				carts = append(carts, &cart{
					facing:   facing,
					position: position,
				})

				switch facing {
				case up:
				case down:
					tracks[position] = vertical
					break
				case left:
				case right:
					tracks[position] = horizontal
					break
				}
			}
		}
	}

	return tracks, carts
}

func getCartDirection(input rune) int {
	switch input {
	case cartUp:
		return up
	case cartRight:
		return right
	case cartDown:
		return down
	case cartLeft:
		return left
	}

	return up
}

func isCart(input rune) bool {
	return input == cartUp || input == cartRight || input == cartDown || input == cartLeft
}

func padString(input string, length int) string {
	inputLen := len(input)

	if inputLen >= length {
		return input
	}

	return input + strings.Repeat(" ", length-inputLen)
}

func positionToCoords(position int, rowLen int) string {
	x := position
	y := 0

	for x > rowLen {
		y++
		x -= rowLen
	}

	return fmt.Sprintf("%d,%d", x, y)
}

func (c *cart) Move(tracks []rune, rowLen int) {
	switch c.facing {
	case up:
		c.position = c.position - rowLen
		break
	case right:
		c.position = c.position + 1
		break
	case down:
		c.position = c.position + rowLen
		break
	case left:
		c.position = c.position - 1
		break
	}

	switch tracks[c.position] {
	case intersection:
		switch c.intersection {
		case turnLeft:
			c.facing--
			c.intersection = goStraight
			break
		case goStraight:
			c.intersection = turnRight
			break
		case turnRight:
			c.facing++
			c.intersection = turnLeft
			break
		}

		if c.facing < up {
			c.facing = left
		}

		if c.facing > left {
			c.facing = up
		}
		break
	case turnOne:
		switch c.facing {
		case up:
			c.facing = right
			break
		case right:
			c.facing = up
			break
		case down:
			c.facing = left
			break
		case left:
			c.facing = down
			break
		}
		break
	case turnTwo:
		switch c.facing {
		case up:
			c.facing = left
			break
		case right:
			c.facing = down
			break
		case down:
			c.facing = right
			break
		case left:
			c.facing = up
			break
		}
		break
	}
}
