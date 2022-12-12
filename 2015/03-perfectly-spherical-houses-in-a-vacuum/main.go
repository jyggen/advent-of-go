package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

func SolvePart1(input string) (string, error) {
	x, y := 0, 0
	visited := map[int]struct{}{
		(y * 1000) + x: {},
	}

	for _, c := range strings.TrimSpace(input) {
		switch c {
		case '^':
			y--
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		}

		visited[(y*1000)+x] = struct{}{}
	}

	return strconv.Itoa(len(visited)), nil
}

func SolvePart2(input string) (string, error) {
	coords := [2][2]int{{0, 0}, {0, 0}}
	current := 0
	visited := map[int]struct{}{
		(coords[0][0] * 1000) + coords[0][1]: {},
	}

	for _, c := range strings.TrimSpace(input) {
		switch c {
		case '^':
			coords[current][0]--
		case '>':
			coords[current][1]++
		case 'v':
			coords[current][0]++
		case '<':
			coords[current][1]--
		}

		visited[(coords[current][0]*1000)+coords[current][1]] = struct{}{}
		current = (current + 1) % 2
	}

	return strconv.Itoa(len(visited)), nil
}
