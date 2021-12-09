package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	masses, err := utils2.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	result := 0

	for _, mass := range masses {
		result += (mass / 3) - 2
	}

	return strconv.Itoa(result), nil
}

func SolvePart2(input string) (string, error) {
	masses, err := utils2.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	total := 0

	for _, mass := range masses {
		result := (mass / 3) - 2
		total += result

		for {
			result = (result / 3) - 2

			if result <= 0 {
				break
			}

			total += result
		}
	}

	return strconv.Itoa(total), nil
}
