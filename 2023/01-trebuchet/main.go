package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"
)

const chars = "123456789"

// better replacements courtesy of /u/pred
var replaces = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func solve(r string) (int, error) {
	var first, last uint8

	first = r[strings.IndexAny(r, chars)]
	last = r[strings.LastIndexAny(r, chars)]

	return strconv.Atoi(string(first) + string(last))
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	sum := 0

	for _, r := range rows {
		number, err := solve(r)

		if err != nil {
			return "", err
		}

		sum += number
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	sum := 0

	for _, r := range rows {
		for from, to := range replaces {
			r = strings.Replace(r, from, to, -1)
		}

		number, err := solve(r)

		if err != nil {
			return "", err
		}

		sum += number
	}

	return strconv.Itoa(sum), nil
}
