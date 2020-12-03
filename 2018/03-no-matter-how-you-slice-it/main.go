package main

import (
	"errors"
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"regexp"
	"strconv"
)

type claim struct {
	id int
	x1 int
	y1 int
	x2 int
	y2 int
}

var inputRegex = regexp.MustCompile("^#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)$")

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	claims, err := makeClaims(utils.ToStringSlice(input, "\n"))

	if err != nil {
		return "", err
	}

	seen := make(map[string]int, 0)

	for _, c := range claims {
		for i := c.x1; i <= c.x2; i++ {
			for j := c.y1; j <= c.y2; j++ {
				seen[fmt.Sprintf("%dx%d", i, j)]++
			}
		}
	}

	numOfDupes := 0

	for _, num := range seen {
		if num > 1 {
			numOfDupes++
		}
	}

	return strconv.Itoa(numOfDupes), nil
}

func SolvePart2(input string) (string, error) {
	claims, err := makeClaims(utils.ToStringSlice(input, "\n"))

	if err != nil {
		return "", err
	}

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

func makeClaims(input []string) ([]*claim, error) {
	claims := make([]*claim, len(input))

	for i, c := range input {
		match := inputRegex.FindStringSubmatch(c)
		if len(match) == 0 {
			return claims, errors.New("unable to find matches")
		}
		id, err := strconv.Atoi(match[1])

		if err != nil {
			return claims, err
		}

		left, err := strconv.Atoi(match[2])

		if err != nil {
			return claims, err
		}

		top, err := strconv.Atoi(match[3])

		if err != nil {
			return claims, err
		}

		width, err := strconv.Atoi(match[4])

		if err != nil {
			return claims, err
		}

		height, err := strconv.Atoi(match[5])

		if err != nil {
			return claims, err
		}

		claims[i] = &claim{
			id: id,
			x1: left,
			y1: top,
			x2: left + width - 1,
			y2: top + height - 1,
		}
	}

	return claims, nil
}
