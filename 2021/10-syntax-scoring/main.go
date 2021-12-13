package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func filterCorrupted(input string) ([][]rune, int) {
	lines := utils.ToRuneSlice(input, "\n")
	score := 0
	incomplete := make([][]rune, 0, len(lines))
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	for _, runes := range lines {
		state := make([]rune, 0)
		corrupted := false

	Loop:
		for _, r := range runes {
			switch r {
			case '(', '[', '{', '<':
				state = append(state, r)
			case ')', ']', '}', '>':
				e := state[len(state)-1]

				if pairs[e] != r {
					score += points[r]
					corrupted = true
					break Loop
				}

				state = state[:len(state)-1]
			}
		}

		if !corrupted {
			incomplete = append(incomplete, runes)
		}
	}

	return incomplete, score
}

func SolvePart1(input string) (string, error) {
	_, score := filterCorrupted(input)

	return strconv.Itoa(score), nil
}

func SolvePart2(input string) (string, error) {
	incomplete, _ := filterCorrupted(input)
	scores := make([]int, 0, len(incomplete))
	points := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	for _, runes := range incomplete {
		state := make([]rune, 0)
		score := 0

		for _, r := range runes {
			switch r {
			case '(', '[', '{', '<':
				state = append(state, r)
			case ')', ']', '}', '>':
				state = state[:len(state)-1]
			}
		}

		for i := len(state) - 1; i >= 0; i-- {
			score *= 5
			score += points[state[i]]
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)

	return strconv.Itoa(scores[len(scores)/2]), nil
}
