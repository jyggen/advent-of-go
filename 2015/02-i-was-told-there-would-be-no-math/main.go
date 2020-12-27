package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"sort"
	"strconv"
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

	for _, present := range presents {
		sides := []int{
			present[0] * present[1],
			present[1] * present[2],
			present[2] * present[0],
		}

		paperSize += sides[0]*2 + sides[1]*2 + sides[2]*2
		paperSize += sides[0]
	}

	return strconv.Itoa(paperSize), nil
}

func SolvePart2(input string) (string, error) {
	presents := parseInput(input)
	ribbonSize := 0

	for _, present := range presents {
		ribbonSize += present[0]*2 + present[1]*2
		ribbonSize += present[0] * present[1] * present[2]
	}

	return strconv.Itoa(ribbonSize), nil
}

func parseInput(input string) [][]int {
	rows := utils.ToStringSlice(input, "\n")
	presents := make([][]int, 0, len(rows))

	for _, present := range rows {
		dimensionStrings := utils.ToStringSlice(present, "x")
		dimensions := make([]int, 0, len(dimensionStrings))

		for _, value := range dimensionStrings {
			intValue, _ := strconv.Atoi(value)

			dimensions = append(dimensions, intValue)
		}

		sort.Ints(dimensions) // lazy

		presents = append(presents, dimensions)
	}

	return presents
}
