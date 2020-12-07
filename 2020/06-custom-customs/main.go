package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"os"
	"strconv"
	"strings"
)

var eolReplacer *strings.Replacer

func init() {
	eolReplacer = strings.NewReplacer("\n\n", ";", "\n", ",")
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
	groups := strings.Split(eolReplacer.Replace(strings.TrimSpace(input)), ";")
	sum := 0

	for _, g := range groups {
		for i := 'a'; i <= 'z'; i++ {
			if strings.Count(g, string(i)) > 0 {
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	groups := strings.Split(eolReplacer.Replace(strings.TrimSpace(input)), ";")
	sum := 0

	for _, g := range groups {
		size := strings.Count(g, ",") + 1

		for i := 'a'; i <= 'z'; i++ {
			if c := strings.Count(g, string(i)); c == size {
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}
