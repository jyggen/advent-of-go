package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"math"
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
	intSlice, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	sort.Ints(intSlice)

	median := intSlice[len(intSlice)/2]

	adjustments := 0

	for _, i := range intSlice {
		diff := i - median

		if diff < 0 {
			diff = -diff
		}

		adjustments += diff
	}

	return strconv.Itoa(adjustments), nil
}

func SolvePart2(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	sum := 0

	for _, v := range intSlice {
		sum += v
	}

	average := float64(sum) / float64(len(intSlice))
	floor := int(math.Floor(average))
	ceil := int(math.Ceil(average))

	best := math.MaxInt

	for _, target := range [2]int{floor, ceil} {
		sum = 0

		for _, v := range intSlice {
			diff := v - target

			if diff < 0 {
				diff = -diff
			}

			for j := 1; j <= diff; j++ {
				sum += j
			}
		}

		if sum < best {
			best = sum
		}
	}

	return strconv.Itoa(best), nil
}
