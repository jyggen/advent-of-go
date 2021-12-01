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
	intSlice, err := utils.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	result := 0

	for i := 1; i < len(intSlice); i++ {
		if intSlice[i] > intSlice[i-1] {
			result++
		}
	}

	return strconv.Itoa(result), err
}

func SolvePart2(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	result := 0

	for i := 0; i < len(intSlice)-3; i++ {
		if intSlice[i+1]+intSlice[i+2]+intSlice[i+3] > intSlice[i]+intSlice[i+1]+intSlice[i+2] {
			result++
		}
	}

	return strconv.Itoa(result), err
}
