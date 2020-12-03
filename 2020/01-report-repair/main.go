package main

import (
	"errors"
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
)

const expectedSum = 2020

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	result1, err := resolve(intSlice, expectedSum, 2, make([]int, 0))

	return strconv.Itoa(result1), err
}

func SolvePart2(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	result1, err := resolve(intSlice, expectedSum, 3, make([]int, 0))

	return strconv.Itoa(result1), err
}

func resolve(input []int, expected int, iterations int, attempt []int) (int, error) {
	for i, val := range input {
		newAttempt := append(attempt, val)
		attemptSum := 0

		for _, val := range newAttempt {
			attemptSum += val
		}

		if attemptSum > expected {
			continue
		}

		if iterations > 1 {
			returnVal, err := resolve(input[i+1:], expected, iterations-1, newAttempt)

			if err == nil {
				return returnVal, nil
			}
		} else {
			if attemptSum == expected {
				result := 1

				for _, val := range newAttempt {
					result = result * val
				}

				return result, nil
			}
		}
	}

	return 0, errors.New("the attempt does not result in the expected number")
}
