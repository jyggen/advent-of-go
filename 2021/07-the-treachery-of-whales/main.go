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
	best := math.Min(calcCost(intSlice, floor), calcCost(intSlice, ceil))

	return strconv.Itoa(int(best)), nil
}

func calcCost(positions []int, target int) float64 {
	sum := 0

	for _, v := range positions {
		diff := v - target

		if diff < 0 {
			diff = -diff
		}

		for j := 1; j <= diff; j++ {
			sum += j
		}
	}

	return float64(sum)
}
