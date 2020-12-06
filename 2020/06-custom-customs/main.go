package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
	"strings"
)

var eolReplacer *strings.Replacer

func init() {
	eolReplacer = strings.NewReplacer("\n", "")
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
	groups := utils.ToStringSlice(input, "\n\n")
	sum := 0

	for _, g := range groups {
		g = eolReplacer.Replace(g)

		for i := 'a'; i <= 'z'; i++ {
			if strings.Count(g, string(i)) > 0 {
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	groups := utils.ToStringSlice(input, "\n\n")
	sum := 0

	for _, g := range groups {
		size := strings.Count(g, "\n") + 1
		g = eolReplacer.Replace(g)

		for i := 'a'; i <= 'z'; i++ {
			if c := strings.Count(g, string(i)); c == size {
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}
