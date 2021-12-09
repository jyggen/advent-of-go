package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var lookup = make([]byte, 127)

func init() {
	upper := byte(65)
	lower := byte(97)

	for i := 0; i < 26; i++ {
		lookup[upper] = lower
		lookup[lower] = upper

		upper++
		lower++
	}
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(react([]byte(strings.TrimSpace(input)))), nil
}

func SolvePart2(input string) (string, error) {
	polymer := []byte(strings.TrimSpace(input))
	variants := make([][]byte, 26)

	for i := 0; i < 26; i++ {
		char := byte(i + 65)
		variants[i] = make([]byte, 0)

		for _, r := range polymer {
			if r != char && r != lookup[char] {
				variants[i] = append(variants[i], r)
			}
		}
	}

	min := 0

	for _, variant := range variants {
		result := react(variant)

		if min == 0 || result < min {
			min = result
		}
	}

	return strconv.Itoa(min), nil
}

func react(polymer []byte) int {
	polymerLen := len(polymer) - 1

	for i := 0; i < polymerLen; i++ {
		if polymer[i] == lookup[polymer[i+1]] {
			polymer = append(polymer[:i], polymer[i+2:]...)
			polymerLen -= 2
			i -= 2

			if i == -2 {
				i = -1
			}
		}
	}

	return polymerLen + 1
}
