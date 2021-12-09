package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

const (
	p1win = iota
	p2win = iota
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	player1, player2 := parse(input)
	winner := play(player1, player2)

	return strconv.Itoa(score(winner)), nil
}

func SolvePart2(input string) (string, error) {
	player1, player2 := parse(input)
	_, winner := playRecurse(player1, player2, true)

	return strconv.Itoa(score(winner)), nil
}

func score(deck []int) int {
	multiplier := len(deck)
	sum := 0

	for _, c := range deck {
		sum += c * multiplier
		multiplier--
	}

	return sum
}

func play(p1deck []int, p2deck []int) []int {
	for {
		if len(p1deck) == 0 {
			return p2deck
		}

		if len(p2deck) == 0 {
			return p1deck
		}

		var p1card, p2card int

		p1card, p1deck = p1deck[0], p1deck[1:]
		p2card, p2deck = p2deck[0], p2deck[1:]

		if p1card > p2card {
			p1deck = append(p1deck, p1card, p2card)
		} else {
			p2deck = append(p2deck, p2card, p1card)
		}
	}
}

func playRecurse(p1deck []int, p2deck []int, root bool) (int, []int) {
	if !root {
		p1high := utils.MaxIntSlice(p1deck)
		p2high := utils.MaxIntSlice(p2deck)

		if p1high > p2high {
			return p1win, p1deck
		}
	}

	cache := make(map[string]bool, 0)

	for {
		id := cacheKey(p1deck, p2deck)

		if _, ok := cache[id]; ok {
			return p1win, p1deck
		}

		cache[id] = true

		var p1card, p2card int
		var winner int

		p1card, p1deck = p1deck[0], p1deck[1:]
		p2card, p2deck = p2deck[0], p2deck[1:]

		if len(p1deck) >= p1card && len(p2deck) >= p2card {
			p1new := make([]int, p1card)
			p2new := make([]int, p2card)

			copy(p1new, p1deck[0:p1card])
			copy(p2new, p2deck[0:p2card])

			winner, _ = playRecurse(p1new, p2new, false)
		} else {
			if p1card > p2card {
				winner = p1win
			} else {
				winner = p2win
			}
		}

		if winner == p1win {
			p1deck = append(p1deck, p1card, p2card)
		} else {
			p2deck = append(p2deck, p2card, p1card)
		}

		if len(p1deck) == 0 {
			return p2win, p2deck
		}

		if len(p2deck) == 0 {
			return p1win, p1deck
		}
	}
}

func parse(input string) ([]int, []int) {
	playerSlice := utils.ToStringSlice(input, "\n\n")
	players := make([][]int, len(playerSlice))

	for i, p := range playerSlice {
		cards, _ := utils.ToIntegerSlice(p[strings.IndexRune(p, '\n'):], "\n")
		players[i] = make([]int, len(cards))

		for j, c := range cards {
			players[i][j] = c
		}
	}

	return players[0], players[1]
}

func cacheKey(p1 []int, p2 []int) string {
	return fmt.Sprint(p1, p2)
}
