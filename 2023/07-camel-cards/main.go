package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

const (
	fiveOfAKind = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
)

const (
	ace = iota
	king
	queen
	jack
	ten
	nine
	eight
	seven
	six
	five
	four
	three
	two
	joker
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type hand struct {
	handType int
	cards    []int
	bid      int
}

func parseHand(input string, useJokers bool) hand {
	delimiterAt := strings.IndexRune(input, ' ')
	cardIdentifiers := []rune(input[:delimiterAt])
	bid, _ := strconv.Atoi(input[delimiterAt+1:])
	h := hand{
		handType: highCard,
		cards:    make([]int, 0, len(cardIdentifiers)),
		bid:      bid,
	}

	counts := make([]int, 13)
	numJokers := 0

	for i, c := range cardIdentifiers {
		switch c {
		case 'A':
			h.cards = append(h.cards, ace)
		case 'K':
			h.cards = append(h.cards, king)
		case 'Q':
			h.cards = append(h.cards, queen)
		case 'J':
			if useJokers {
				h.cards = append(h.cards, joker)
				numJokers++
			} else {
				h.cards = append(h.cards, jack)
			}
		case 'T':
			h.cards = append(h.cards, ten)
		case '9':
			h.cards = append(h.cards, nine)
		case '8':
			h.cards = append(h.cards, eight)
		case '7':
			h.cards = append(h.cards, seven)
		case '6':
			h.cards = append(h.cards, six)
		case '5':
			h.cards = append(h.cards, five)
		case '4':
			h.cards = append(h.cards, four)
		case '3':
			h.cards = append(h.cards, three)
		case '2':
			h.cards = append(h.cards, two)
		}

		if !useJokers || c != 'J' {
			counts[h.cards[i]]++
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	switch {
	case counts[0] == (5 - numJokers):
		h.handType = fiveOfAKind
	case counts[0] == (4 - numJokers):
		h.handType = fourOfAKind
	case counts[0] == (3-numJokers) && counts[1] == 2:
		h.handType = fullHouse
	case counts[0] == (3 - numJokers):
		h.handType = threeOfAKind
	case counts[0] == (2-numJokers) && counts[1] == 2:
		h.handType = twoPair
	case counts[0] == (2 - numJokers):
		h.handType = onePair
	default:
		h.handType = highCard
	}

	return h
}

func parseHands(input string, useJokers bool) []hand {
	rows := utils.ToStringSlice(input, "\n")
	hands := make([]hand, 0, len(rows))

	for _, r := range rows {
		hands = append(hands, parseHand(r, useJokers))
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for k := 0; k < len(hands[i].cards); k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}

				return hands[i].cards[k] > hands[j].cards[k]
			}
		}

		return hands[i].handType > hands[j].handType
	})

	return hands
}

func SolvePart1(input string) (string, error) {
	hands := parseHands(input, false)
	total := 0

	for i, h := range hands {
		total += (i + 1) * h.bid
	}

	return strconv.Itoa(total), nil
}

func SolvePart2(input string) (string, error) {
	hands := parseHands(input, true)
	total := 0

	for i, h := range hands {
		total += (i + 1) * h.bid
	}

	return strconv.Itoa(total), nil
}
