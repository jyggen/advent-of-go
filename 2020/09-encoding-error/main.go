package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
)

const preamble = 25

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	numbers, err := utils.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	return strconv.Itoa(preambleSolver(numbers)), nil
}

func SolvePart2(input string) (string, error) {
	numbers, err := utils.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	expected := preambleSolver(numbers)
	i := len(numbers)

	var min, max int

ContiguousLoop:
	for {
		i--

		if numbers[i] > expected {
			continue
		}

		sum, j := numbers[i], i
		min, max = sum, sum

		for {
			j--

			sum += numbers[j]

			if sum > expected {
				break
			}

			if numbers[j] > max {
				max = numbers[j]
			}

			if numbers[j] < min {
				min = numbers[j]
			}

			if sum == expected {
				break ContiguousLoop
			}
		}
	}

	return strconv.Itoa(min + max), nil
}

func preambleSolver(numbers []int) int {
	i := preamble

	for {
		found := false

	PreambleLoop:
		for j := i - preamble; j < i; j++ {
			if numbers[j] > numbers[i] {
				continue
			}

			for k := j + 1; k <= i; k++ {
				if numbers[j]+numbers[k] == numbers[i] {
					found = true
					break PreambleLoop
				}
			}
		}

		if !found {
			break
		}

		i++
	}

	return numbers[i]
}
