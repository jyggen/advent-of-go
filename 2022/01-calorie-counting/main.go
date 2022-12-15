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

func solve(input string) []int {
	elves := utils.ToStringSlice(input, "\n\n")
	sums := make([]int, len(elves))

	for i, elf := range elves {
		for _, value := range utils.ToOptimisticIntSlice(elf, true) {
			sums[i] += value
		}
	}

	sort.Ints(sums)

	return sums
}

func SolvePart1(input string) (string, error) {
	sums := solve(input)

	return strconv.Itoa(sums[len(sums)-1]), nil
}

func SolvePart2(input string) (string, error) {
	sums := solve(input)

	return strconv.Itoa(sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]), nil
}
