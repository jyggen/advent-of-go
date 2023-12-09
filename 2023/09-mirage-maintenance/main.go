package main

import (
	"fmt"
	"os"
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

func solve(input string, beginning bool) int {
	lines := utils.ToStringSlice(input, "\n")
	sum := 0

	for _, l := range lines {
		sequences := make([][]int, 0, 2)
		sequences = append(sequences, utils.ToOptimisticIntSlice(l, true))

		for i := 0; i < len(sequences); i++ {
			newSequence := make([]int, len(sequences[i])-1)
			allZero := true

			for j := 1; j < len(sequences[i]); j++ {
				newSequence[j-1] = sequences[i][j] - sequences[i][j-1]

				if allZero && newSequence[j-1] != 0 {
					allZero = false
				}
			}

			sequences = append(sequences, newSequence)

			if allZero {
				break
			}
		}

		value := 0

		if beginning {
			for i := len(sequences) - 2; i >= 0; i-- {
				value = sequences[i][0] - value
			}
		} else {
			for i := len(sequences) - 2; i >= 0; i-- {
				value = sequences[i][len(sequences[i])-1] + value
			}
		}

		sum += value
	}

	return sum
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(solve(input, false)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(solve(input, true)), nil
}
