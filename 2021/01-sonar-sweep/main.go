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

func SolvePart1(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, "\n")
	if err != nil {
		return "", err
	}

	result := 0

	for i, j := 1, 0; i < len(intSlice); i, j = i+1, j+1 {
		if intSlice[i] > intSlice[j] {
			result++
		}
	}

	return strconv.Itoa(result), nil
}

func SolvePart2(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, "\n")
	if err != nil {
		return "", err
	}

	result := 0

	for i, j, k, l := 3, 2, 1, 0; i < len(intSlice); i, j, k, l = i+1, j+1, k+1, l+1 {
		if intSlice[j]+intSlice[k]+intSlice[l] < intSlice[i]+intSlice[j]+intSlice[k] {
			result++
		}
	}

	return strconv.Itoa(result), nil
}
