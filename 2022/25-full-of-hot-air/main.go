package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func fromSnafu(snafu string) int {
	result := 0

	for i, r := range snafu {
		pow := utils.IntPow(5, len(snafu)-i-1)

		switch r {
		case '2':
			result += pow * 2
		case '1':
			result += pow * 1
		case '0':
			result += pow * 0
		case '-':
			result += pow * -1
		case '=':
			result += pow * -2
		}
	}

	return result
}

func toSnafu(decimal int) string {
	if decimal == 0 {
		return ""
	}

	result := ""
	remainder := (decimal + 2) % 5

	switch remainder {
	case 0:
		result += "="
	case 1:
		result += "-"
	case 2:
		result += "0"
	case 3:
		result += "1"
	case 4:
		result += "2"

	}

	return toSnafu((decimal+2)/5) + result
}

func SolvePart1(input string) (string, error) {
	snafus := utils.ToStringSlice(input, "\n")
	sum := 0

	for _, s := range snafus {
		sum += fromSnafu(s)
	}

	return toSnafu(sum), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(0), nil
}
