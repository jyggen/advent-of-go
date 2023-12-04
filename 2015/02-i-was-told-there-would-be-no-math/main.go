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

func SolvePart1(input string) (string, error) {
	presents := parseInput(input)
	paperSize := 0

	for i := 0; i < len(presents); i += 3 {
		sides := []int{
			presents[i] * presents[i+1],
			presents[i+1] * presents[i+2],
			presents[i+2] * presents[i],
		}

		paperSize += sides[0]*2 + sides[1]*2 + sides[2]*2
		paperSize += slices.Min(sides)
	}

	return strconv.Itoa(paperSize), nil
}

func SolvePart2(input string) (string, error) {
	presents := parseInput(input)
	ribbonSize := 0

	for i := 0; i < len(presents); i += 3 {
		ribbonSize += (presents[i]*2 + presents[i+1]*2 + presents[i+2]*2) - (slices.Max(presents[i:i+3]) * 2)
		ribbonSize += presents[i] * presents[i+1] * presents[i+2]
	}

	return strconv.Itoa(ribbonSize), nil
}

func parseInput(input string) []int {
	return utils.ToOptimisticIntSlice(input, true)
}
