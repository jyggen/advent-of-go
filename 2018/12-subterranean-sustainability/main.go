package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

type rule struct {
	pattern []bool
	next    bool
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
	pots, rules := makePotsAndRules(utils.ToStringSlice(input, "\n"))

	var sum int

	for generation := 1; generation <= 20; generation++ {
		pots = append([]bool{false, false, false, false}, pots...)
		pots = append(pots, false, false, false, false)
		potsLen := len(pots)
		nextGen := make([]bool, potsLen)
		loopMax := potsLen - 2

		for i := 2; i < loopMax; i++ {
			for _, r := range rules {
				if pots[i-2] == r.pattern[0] &&
					pots[i-1] == r.pattern[1] &&
					pots[i] == r.pattern[2] &&
					pots[i+1] == r.pattern[3] &&
					pots[i+2] == r.pattern[4] {
					nextGen[i] = r.next
				}
			}
		}

		pots = nextGen
		sum = 0

		for i, p := range pots {
			if p {
				sum += i - 4*generation
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	pots, rules := makePotsAndRules(utils.ToStringSlice(input, "\n"))

	diffs := make([]int, 10)
	generation := 1
	prevSum := 0

	for {
		pots = append([]bool{false, false, false, false}, pots...)
		pots = append(pots, false, false, false, false)
		potsLen := len(pots)
		nextGen := make([]bool, potsLen)
		loopMax := potsLen - 2

		for i := 2; i < loopMax; i++ {
			for _, r := range rules {
				if pots[i-2] == r.pattern[0] &&
					pots[i-1] == r.pattern[1] &&
					pots[i] == r.pattern[2] &&
					pots[i+1] == r.pattern[3] &&
					pots[i+2] == r.pattern[4] {
					nextGen[i] = r.next
				}
			}
		}

		pots = nextGen
		sum := 0

		for i, p := range pots {
			if p {
				sum += i - 4*generation
			}
		}

		diffs = append(diffs[1:], sum-prevSum)
		last := diffs[0]
		match := true
		prevSum = sum

		for _, d := range diffs {
			if last != d {
				match = false
				break
			}

			last = d
		}

		if match {
			break
		}

		generation++
	}

	return fmt.Sprintf("%.0f", float64(prevSum)+float64(diffs[0])*(50*math.Pow(10, 9)-float64(generation))), nil
}

func makePotsAndRules(input []string) ([]bool, []*rule) {
	initial := strings.Split(input[0][15:], "")
	pots := make([]bool, len(initial))
	rules := make([]*rule, len(input[2:]))

	for i, state := range initial {
		if state == "#" {
			pots[i] = true
		} else {
			pots[i] = false
		}
	}

	for i, r := range input[2:] {
		patternParts := strings.Split(r[:5], "")
		pattern := make([]bool, len(patternParts))

		for j, p := range patternParts {
			if p == "#" {
				pattern[j] = true
			} else {
				pattern[j] = false
			}
		}

		next := true

		if string(r[9]) == "." {
			next = false
		}

		rules[i] = &rule{
			pattern: pattern,
			next:    next,
		}
	}

	return pots, rules
}
