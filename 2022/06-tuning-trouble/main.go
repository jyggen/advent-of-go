package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
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

func solve(input string, length int) int {
	withoutLast := length - 1

Loop:
	for i := withoutLast; i < len(input); {
		for j := 0; j < withoutLast; j++ {
			char := input[i-j]

			for k := j + 1; k < length; k++ {
				if char == input[i-k] {
					i = i - k + length
					continue Loop
				}
			}
		}
		return i + 1
	}

	return -1
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(solve(input, 4)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(solve(input, 14)), nil
}
