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
	boundaries, err := utils2.ToIntegerSlice(input, "-")

	if err != nil {
		return "", nil
	}

	found := 0

RangeLoop:
	for i := boundaries[0]; i <= boundaries[1]; i++ {
		digits := []rune(strconv.Itoa(i))
		adjacent := false

		for j, digit := range digits {
			if j == 0 {
				continue
			}

			if digit < digits[j-1] {
				continue RangeLoop
			}

			if digit == digits[j-1] {
				adjacent = true
			}
		}

		if !adjacent {
			continue
		}

		found++
	}

	return strconv.Itoa(found), nil
}

func SolvePart2(input string) (string, error) {
	boundaries, err := utils2.ToIntegerSlice(input, "-")

	if err != nil {
		return "", nil
	}

	found := 0

RangeLoop:
	for i := boundaries[0]; i <= boundaries[1]; i++ {
		digits := []rune(strconv.Itoa(i))
		adjacent := false

		var streak int
		var lastRepeat rune

		for j, digit := range digits {
			if j == 0 {
				continue
			}

			if digit < digits[j-1] {
				continue RangeLoop
			}

			if digit == digits[j-1] {
				if lastRepeat == digit {
					streak++
				} else {
					streak = 2
				}

				lastRepeat = digit
			} else if streak == 2 {
				adjacent = true
			}
		}

		if !adjacent && streak != 2 {
			continue
		}

		found++
	}

	return strconv.Itoa(found), nil
}
