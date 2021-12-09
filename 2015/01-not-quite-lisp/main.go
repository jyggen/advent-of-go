package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func SolvePart1(input string) (string, error) {
	up := strings.Count(input, "(")
	down := strings.Count(input, ")")

	return strconv.Itoa(up - down), nil
}

func SolvePart2(input string) (string, error) {
	var floor int

	for k, v := range input {
		if v == '(' {
			floor++
		} else {
			floor--
		}

		if floor < 0 {
			return strconv.Itoa(k + 1), nil
		}
	}

	return strconv.Itoa(-1), errors.New("unable to find solution")
}
