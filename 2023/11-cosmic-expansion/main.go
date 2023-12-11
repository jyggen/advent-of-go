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

func solve(input string, increase int) int {
	universe := utils.ToRuneSlice(input, "\n")
	galaxies := make([][2]int, 0)
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)

	for y, points := range universe {
		allEmpty := true

		for x, point := range points {
			if point == '#' {
				galaxies = append(galaxies, [2]int{y, x})
				allEmpty = false
			}
		}

		if allEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := range universe[0] {
		allEmpty := true

		for _, point := range universe {
			if point[x] == '#' {
				allEmpty = false

				break
			}
		}

		if allEmpty {
			emptyCols = append(emptyCols, x)
		}
	}

	for i, g := range galaxies {
		for j, y := range emptyRows {
			if y+(j*increase) < g[0] {
				g[0] += increase
			} else {
				break
			}
		}

		for j, x := range emptyCols {
			if x+(j*increase) < g[1] {
				g[1] += increase
			} else {
				break
			}
		}

		galaxies[i] = g
	}

	sum := 0

	for i, g := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			sum += utils.ManhattanDistance(g[1]-g2[1], g[0]-g2[0])
		}
	}

	return sum
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(solve(input, 1)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(solve(input, 999999)), nil
}
