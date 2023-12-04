package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"slices"
	"strconv"
	"strings"
)

var spaceReplacer *strings.Replacer

func init() {
	spaceReplacer = strings.NewReplacer("   ", " ", "  ", " ")
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type card struct {
	number         int
	winningNumbers int
}

func SolvePart1(input string) (string, error) {
	input = spaceReplacer.Replace(input)
	rows := utils.ToStringSlice(input, "\n")
	sum := 0

	for _, r := range rows {
		value := 0
		colonIndex := strings.Index(r, ": ")
		pipeIndex := strings.Index(r, " | ")
		winning := utils.ToOptimisticIntSlice(r[colonIndex+2:pipeIndex], false)
		all := utils.ToOptimisticIntSlice(r[pipeIndex+3:], false)

		for _, w := range winning {
			if slices.Contains(all, w) {
				if value == 0 {
					value = 1
				} else {
					value *= 2
				}
			}
		}

		sum += value
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	input = spaceReplacer.Replace(input)
	rows := utils.ToStringSlice(input, "\n")
	cards := make([]*card, 0, len(rows))

	for i, r := range rows {
		colonIndex := strings.Index(r, ": ")
		pipeIndex := strings.Index(r, " | ")
		winning := utils.ToOptimisticIntSlice(r[colonIndex+2:pipeIndex], false)
		all := utils.ToOptimisticIntSlice(r[pipeIndex+3:], false)
		c := &card{winningNumbers: 0, number: i}

		for _, w := range winning {
			if slices.Contains(all, w) {
				c.winningNumbers++
			}
		}

		cards = append(cards, c)
	}

	for i := 0; i < len(cards); i++ {
		for j := 1; j <= cards[i].winningNumbers; j++ {
			if i+j == len(cards) {
				break
			}

			cards = append(cards, cards[cards[i].number+j])
		}
	}

	return strconv.Itoa(len(cards)), nil
}
