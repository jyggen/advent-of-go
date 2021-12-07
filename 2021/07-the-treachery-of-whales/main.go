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
	cost := 0

	for _, i := range intSlice {
		diff := i - median

		if diff < 0 {
			diff = -diff
		}

		cost += diff
	}

	return strconv.Itoa(cost), nil
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

	floor := sum / len(intSlice)
	ceil := floor + 1
	best := math.MaxInt

	for _, v := range [2]int{
		calcCost(intSlice, floor),
		calcCost(intSlice, ceil),
	} {
		if v < best {
			best = v
		}
	}

	return strconv.Itoa(best), nil
}

func calcCost(positions []int, target int) int {
	cost := 0

	for _, v := range positions {
		diff := v - target

		if diff < 0 {
			diff = -diff
		}

		cost += diff * (diff + 1) / 2
	}

	return cost
}
