package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

type coordinate struct {
	x     int
	y     int
	total int
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
	result, _, err := solve(input)

	return result, err
}

func SolvePart2(input string) (string, error) {
	_, result, err := solve(input)

	return result, err
}

func solve(input string) (string, string, error) {
	coords, err := makeCoords(utils.ToStringSlice(input, "\n"))
	if err != nil {
		return "", "", err
	}

	lowX, lowY, highX, highY := 10000, 10000, 0, 0

	for _, coord := range coords {
		if coord.x < lowX {
			lowX = coord.x
		}

		if coord.x > highX {
			highX = coord.x
		}

		if coord.y < lowY {
			lowY = coord.y
		}

		if coord.y > highY {
			highY = coord.y
		}
	}

	best := 0
	size := 0

	for x := lowX; x <= highX; x++ {
		for y := lowY; y <= highY; y++ {
			closestDistance := 10000
			closestIndex := -1
			totalDistance := 0

			for i, c := range coords {
				distance := utils.AbsInt(c.x-x) + utils.AbsInt(c.y-y)
				totalDistance += distance

				if distance < closestDistance {
					closestDistance = distance
					closestIndex = i
				}
			}

			if totalDistance < 10000 {
				size++
			}

			coords[closestIndex].total++
		}
	}

	for _, c := range coords {
		if c.total > best {
			best = c.total
		}
	}

	return strconv.Itoa(best), strconv.Itoa(size), nil
}

func makeCoords(input []string) ([]*coordinate, error) {
	coords := make([]*coordinate, len(input))

	for i, l := range input {
		split := utils.ToStringSlice(l, ", ")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			return coords, err
		}

		y, err := strconv.Atoi(split[1])
		if err != nil {
			return coords, err
		}

		coords[i] = &coordinate{
			x:     x,
			y:     y,
			total: 0,
		}
	}

	return coords, nil
}
