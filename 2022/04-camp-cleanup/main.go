package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
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
	integers := utils.ToOptimisticIntSlice(input, false)
	overlaps := 0

	for i := 0; i < len(integers); i += 4 {
		if (integers[i] <= integers[i+2] && integers[i+1] >= integers[i+3]) || (integers[i+2] <= integers[i] && integers[i+3] >= integers[i+1]) {
			overlaps++
		}
	}

	return strconv.Itoa(overlaps), nil
}

func SolvePart2(input string) (string, error) {
	integers := utils.ToOptimisticIntSlice(input, false)
	overlaps := 0

	for i := 0; i < len(integers); i += 4 {
		if (integers[i] <= integers[i+2] && integers[i+1] >= integers[i+2]) || (integers[i+2] <= integers[i] && integers[i+3] >= integers[i]) {
			overlaps++
		}
	}

	return strconv.Itoa(overlaps), nil
}
