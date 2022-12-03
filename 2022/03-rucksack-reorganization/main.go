package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func uniqueSorted(input []rune) []rune {
	output := make([]rune, 0, len(input))
	exists := make(map[rune]struct{})

	for _, r := range input {
		if _, ok := exists[r]; ok {
			continue
		}

		output = append(output, r)
		exists[r] = struct{}{}
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i] < output[j]
	})

	return output
}

func SolvePart1(input string) (string, error) {
	rucksacks := utils.ToStringSlice(input, "\n")
	priorities := 0

	for _, rucksack := range rucksacks {
		first := rucksack[:len(rucksack)/2]
		second := rucksack[len(rucksack)/2:]
		item := first[strings.IndexAny(first, second)]
		priorities += int(item)

		if item >= 'a' {
			priorities -= 96
		} else {
			priorities -= 38
		}
	}

	return strconv.Itoa(priorities), nil
}

func SolvePart2(input string) (string, error) {
	rucksacks := utils.ToStringSlice(input, "\n")
	priorities := 0

	for i := 0; i < len(rucksacks); i += 3 {
		first := rucksacks[i]
		second := rucksacks[i+1]
		third := rucksacks[i+2]
		idx := 0

		var item uint8

		for {
			match := strings.IndexAny(first[idx:], second)

			if strings.IndexRune(third, rune(first[match+idx])) != -1 {
				item = first[match+idx]
				break
			} else {
				idx += match + 1
			}
		}

		priorities += int(item)

		if item >= 'a' {
			priorities -= 96
		} else {
			priorities -= 38
		}
	}

	return strconv.Itoa(priorities), nil
}
