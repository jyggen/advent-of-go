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
	intSlice, err := utils2.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	result := 0

	for i := 1; i < len(intSlice); i++ {
		if intSlice[i] > intSlice[i-1] {
			result++
		}
	}

	return strconv.Itoa(result), nil
}

func SolvePart2(input string) (string, error) {
	intSlice, err := utils2.ToIntegerSlice(input, "\n")

	if err != nil {
		return "", err
	}

	result := 0

	for i := 0; i < len(intSlice)-3; i++ {
		if intSlice[i+1]+intSlice[i+2]+intSlice[i+3] > intSlice[i]+intSlice[i+1]+intSlice[i+2] {
			result++
		}
	}

	return strconv.Itoa(result), nil
}
