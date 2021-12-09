package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

type workOrder struct {
	letter      rune
	completedAt int
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
	rules, runes := buildRulesAndRunes(utils.ToStringSlice(input, "\n"))
	result, _ := build(runes, rules, 1, 0)

	return string(result), nil
}

func SolvePart2(input string) (string, error) {
	rules, runes := buildRulesAndRunes(utils.ToStringSlice(input, "\n"))
	_, timeSpent := build(runes, rules, 5, 60)

	return strconv.Itoa(timeSpent), nil
}

func build(runes []rune, rules [][]rune, workers int, stepDur int) ([]rune, int) {
	avail := make([]rune, len(runes))
	availWorkers := workers
	clock := 0
	curr := make([]rune, 0)
	runesLen := len(runes)
	workOrders := make([]*workOrder, 0)

	copy(avail, runes)

	for len(curr) != runesLen {
		built := false

		for k, v := range workOrders {
			if v.completedAt == clock {
				built = true
				curr = append(curr, v.letter)
				workOrders = append(workOrders[:k], workOrders[k+1:]...)
				availWorkers++
			}
		}

		if len(avail) != 0 && (built || len(workOrders) == 0) {
			possibilities := getPossibilities(curr, avail, rules)

			if len(possibilities) > 0 {
				for i := 0; availWorkers > 0; i++ {
					var r rune

					r, possibilities = possibilities[0], possibilities[1:]
					avail = removeRuneFromSlice(avail, r)
					workOrders = append(workOrders, &workOrder{
						letter:      r,
						completedAt: clock + int(r) - 64 + stepDur,
					})

					availWorkers--

					if len(possibilities) == 0 {
						break
					}
				}
			}
		}

		clock++
	}

	return curr, clock - 1
}

func getPossibilities(curr []rune, avail []rune, rules [][]rune) []rune {
	currLen := len(curr)
	possibilities := make([]rune, 0)

	for _, v := range avail {
		newCurr := make([]rune, currLen)
		copy(newCurr, curr)
		newCurr = append(newCurr, v)

		if !isValid(newCurr, rules) {
			continue
		}

		possibilities = append(possibilities, v)
	}

	return possibilities
}

func isValid(str []rune, rules [][]rune) bool {
	for _, rule := range rules {
		i, j := -1, -1

		for k, v := range str {
			if v == rule[0] {
				i = k
			}

			if v == rule[1] {
				j = k
			}
		}

		if j == -1 {
			continue
		}

		if i == -1 || i > j {
			return false
		}
	}

	return true
}

func removeRuneFromSlice(slice []rune, r rune) []rune {
	for k, v := range slice {
		if v == r {
			slice = append(slice[:k], slice[k+1:]...)
			break
		}
	}

	return slice
}

func buildRulesAndRunes(input []string) ([][]rune, []rune) {
	rules := make([][]rune, len(input))
	steps := make(map[rune]bool, 0)

	for i, l := range input {
		rules[i] = []rune{
			rune(l[5]),
			rune(l[36]),
		}

		for _, letter := range rules[i] {
			if _, ok := steps[letter]; !ok {
				steps[letter] = true
			}
		}
	}

	runes := make([]rune, 0, len(steps))

	for k := range steps {
		runes = append(runes, k)
	}

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return rules, runes
}
