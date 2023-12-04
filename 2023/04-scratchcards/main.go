package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
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
	winningNumbers int
	count          int
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

	for _, r := range rows {
		colonIndex := strings.Index(r, ": ")
		pipeIndex := strings.Index(r, " | ")
		winning := utils.ToOptimisticIntSlice(r[colonIndex+2:pipeIndex], false)
		all := utils.ToOptimisticIntSlice(r[pipeIndex+3:], false)
		c := &card{winningNumbers: 0, count: 1}

		for _, w := range winning {
			if slices.Contains(all, w) {
				c.winningNumbers++
			}
		}

		cards = append(cards, c)
	}

	for i, c := range cards {
		for j := 1; j <= c.winningNumbers; j++ {
			if i+j == len(cards) {
				break
			}

			cards[i+j].count += c.count
		}
	}

	sum := 0

	for _, c := range cards {
		sum += c.count
	}

	return strconv.Itoa(sum), nil
}
