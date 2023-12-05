package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

type claim struct {
	id int
	x1 int
	y1 int
	x2 int
	y2 int
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
	claims := makeClaims(input)

	var maxX, maxY int

	for _, c := range claims {
		if c.y2 > maxY {
			maxY = c.y2
		}

		if c.x2 > maxX {
			maxX = c.x2
		}
	}

	seen := make([]int, (maxX+1)*(maxY+1))
	numOfDupes := 0

	for _, c := range claims {
		for i := c.x1; i <= c.x2; i++ {
			for j := c.y1; j <= c.y2; j++ {
				seen[i+(maxY*j)]++

				if seen[i+(maxY*j)] == 2 {
					numOfDupes++
				}
			}
		}
	}

	return strconv.Itoa(numOfDupes), nil
}

func SolvePart2(input string) (string, error) {
	claims := makeClaims(input)

Loop:
	for _, c1 := range claims {
		for _, c2 := range claims {
			if c1 == c2 {
				continue
			}

			for _, check := range []struct {
				x int
				y int
				c *claim
			}{
				{x: c1.x1, y: c1.y1, c: c2},
				{x: c1.x2, y: c1.y1, c: c2},
				{x: c1.x1, y: c1.y2, c: c2},
				{x: c1.x2, y: c1.y2, c: c2},
				{x: c2.x1, y: c2.y1, c: c1},
				{x: c2.x2, y: c2.y1, c: c1},
				{x: c2.x1, y: c2.y2, c: c1},
				{x: c2.x2, y: c2.y2, c: c1},
			} {
				if check.x >= check.c.x1 && check.x <= check.c.x2 && check.y >= check.c.y1 && check.y <= check.c.y2 {
					continue Loop
				}
			}
		}

		return strconv.Itoa(c1.id), nil
	}

	return "", errors.New("unable to solve with the provided input")
}

func makeClaims(input string) []*claim {
	ints := utils.ToOptimisticIntSlice(input, false)
	claims := make([]*claim, 0, len(ints)/5)

	for i := 0; i < len(ints); i += 5 {
		claims = append(claims, &claim{
			id: ints[i],
			x1: ints[i+1],
			y1: ints[i+2],
			x2: ints[i+1] + ints[i+3] - 1,
			y2: ints[i+2] + ints[i+4] - 1,
		})
	}

	return claims
}
