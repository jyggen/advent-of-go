package main

import (
	"fmt"
	"os"
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

const (
	ROCK = iota + 1
	PAPER
	SCISSOR
)

func SolvePart1(input string) (string, error) {
	rounds := utils.ToRuneSlice(input, "\n")
	score := 0

	for _, round := range rounds {
		theirs := round[0] - 64
		ours := round[2] - 87

		switch {
		case theirs == ours:
			score += 3 + int(ours)
		case theirs == ROCK && ours == PAPER:
			score += 6 + PAPER
		case theirs == PAPER && ours == SCISSOR:
			score += 6 + SCISSOR
		case theirs == SCISSOR && ours == ROCK:
			score += 6 + ROCK
		default:
			score += int(ours)
		}
	}

	return strconv.Itoa(score), nil
}

func SolvePart2(input string) (string, error) {
	rounds := utils.ToRuneSlice(input, "\n")
	score := 0

	for _, round := range rounds {
		theirs := round[0] - 64
		outcome := round[2]

		switch outcome {
		case 'X':
			switch theirs {
			case ROCK:
				score += SCISSOR
			case PAPER:
				score += ROCK
			case SCISSOR:
				score += PAPER
			}
		case 'Y':
			score += int(theirs) + 3
		case 'Z':
			switch theirs {
			case ROCK:
				score += PAPER + 6
			case PAPER:
				score += SCISSOR + 6
			case SCISSOR:
				score += ROCK + 6
			}
		}
	}

	return strconv.Itoa(score), nil
}
