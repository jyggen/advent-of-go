package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	count := strings.Count(input[0:strings.Index(input, "\n")], " ")
	integers := utils.ToOptimisticIntSlice(input, true)
	sum := 0

	for k := 0; k < len(integers); k += count + 1 {
		allZero := false

		for i := 0; !allZero; i++ {
			allZero = true

			for j := k + 1; j < k+count+1-i; j++ {
				integers[j-1] = integers[j] - integers[j-1]

				if allZero && integers[j-1] != 0 {
					allZero = false
				}
			}
		}

		value := 0

		for i := k; i < k+count+1; i++ {
			value = integers[i] + value
		}

		sum += value
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	count := strings.Count(input[0:strings.Index(input, "\n")], " ")
	integers := utils.ToOptimisticIntSlice(input, true)
	sum := 0

	for k := 0; k < len(integers); k += count + 1 {
		allZero := false

		for i := 0; !allZero; i++ {
			allZero = true

			for j := k + count - 1; j >= k+i; j-- {
				integers[j+1] -= integers[j]

				if allZero && integers[j+1] != 0 {
					allZero = false
				}
			}
		}

		value := 0

		for i := k + count; i >= k; i-- {
			value = integers[i] - value
		}

		sum += value
	}

	return strconv.Itoa(sum), nil
}
