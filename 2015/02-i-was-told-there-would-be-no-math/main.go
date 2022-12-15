package main

import (
	"fmt"
	"os"
	"sort"
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

func SolvePart1(input string) (string, error) {
	presents := parseInput(input)
	paperSize := 0

	for i := 0; i < len(presents); i += 3 {
		dimensions := make([]int, 3)
		copy(dimensions, presents[i:i+3])
		sort.Ints(dimensions)

		sides := []int{
			dimensions[0] * dimensions[1],
			dimensions[1] * dimensions[2],
			dimensions[2] * dimensions[0],
		}

		paperSize += sides[0]*2 + sides[1]*2 + sides[2]*2
		paperSize += sides[0]
	}

	return strconv.Itoa(paperSize), nil
}

func SolvePart2(input string) (string, error) {
	presents := parseInput(input)
	ribbonSize := 0

	for i := 0; i < len(presents); i += 3 {
		dimensions := make([]int, 3)
		copy(dimensions, presents[i:i+3])
		sort.Ints(dimensions)

		ribbonSize += dimensions[0]*2 + dimensions[1]*2
		ribbonSize += dimensions[0] * dimensions[1] * dimensions[2]
	}

	return strconv.Itoa(ribbonSize), nil
}

func parseInput(input string) []int {
	return utils.ToOptimisticIntSlice(input, true)
}
