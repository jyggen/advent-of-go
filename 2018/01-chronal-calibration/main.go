package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
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
	deltas, err := utils.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	frequency := 0

	for _, delta := range deltas {
		frequency = frequency + delta
	}

	return strconv.Itoa(frequency), nil
}

func SolvePart2(input string) (string, error) {
	deltas, err := utils.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	numOfDeltas := len(deltas)
	frequency := 0
	i := 0
	seen := map[int]bool{
		0: true,
	}

	for {
		frequency += deltas[i]

		if _, ok := seen[frequency]; ok {
			break
		}

		seen[frequency] = true
		i = (i + 1) % numOfDeltas
	}

	return strconv.Itoa(frequency), nil
}
