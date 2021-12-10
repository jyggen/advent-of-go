package main

import (
	"container/list"
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

	for _, runes := range lines {
		state := list.New()
		corrupted := false

		for _, r := range runes {
			switch r {
			case '(':
				state.PushBack(0)
			case ')':
				e := state.Back()

				if e.Value.(int) != 0 {
					score += 3
					corrupted = true
					break
				}

				state.Remove(e)
			case '[':
				state.PushBack(1)
			case ']':
				e := state.Back()

				if e.Value.(int) != 1 {
					score += 57
					corrupted = true
					break
				}

				state.Remove(e)
			case '{':
				state.PushBack(2)
			case '}':
				e := state.Back()

				if e.Value.(int) != 2 {
					score += 1197
					corrupted = true
					break
				}

				state.Remove(e)
			case '<':
				state.PushBack(3)
			case '>':
				e := state.Back()

				if e.Value.(int) != 3 {
					score += 25137
					corrupted = true
					break
				}

				state.Remove(e)
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

	for _, runes := range incomplete {
		state := list.New()
		score := 0

		for _, r := range runes {
			switch r {
			case '(':
				state.PushBack(0)
			case ')':
				e := state.Back()
				state.Remove(e)
			case '[':
				state.PushBack(1)
			case ']':
				e := state.Back()
				state.Remove(e)
			case '{':
				state.PushBack(2)
			case '}':
				e := state.Back()
				state.Remove(e)
			case '<':
				state.PushBack(3)
			case '>':
				e := state.Back()
				state.Remove(e)
			}
		}

		for state.Len() > 0 {
			e := state.Back()
			score *= 5

			switch e.Value.(int) {
			case 0:
				score += 1
			case 1:
				score += 2
			case 2:
				score += 3
			case 3:
				score += 4
			}

			state.Remove(e)
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)

	return strconv.Itoa(scores[len(scores)/2]), nil
}
