package utils

import (
	"strconv"
	"strings"
)

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func ManhattanDistance(x int, y int) int {
	return AbsInt(x) + AbsInt(y)
}

func ToIntegerSlice(input string, separator string) ([]int, error) {
	stringSlice := ToStringSlice(input, separator)
	integerSlice := make([]int, len(stringSlice))

	for i, val := range stringSlice {
		numVal, err := strconv.Atoi(val)

		if err != nil {
			return integerSlice, err
		}

		integerSlice[i] = numVal
	}

	return integerSlice, nil
}

func ToRuneSlice(input string, separator string) [][]rune {
	stringSlice := ToStringSlice(input, separator)
	runeSlice := make([][]rune, len(stringSlice))

	for i, val := range stringSlice {
		runeSlice[i] = []rune(val)
	}

	return runeSlice
}

func ToStringSlice(input string, separator string) []string {
	return strings.Split(strings.TrimSpace(input), separator)
}
