package main

import (
	"errors"
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
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
	boxes := utils.ToStringSlice(input, "\n")
	ids := make([][]rune, len(boxes))
	twos, threes := 0, 0

	for i, box := range boxes {
		letters := make(map[rune]int)
		ids[i] = []rune(box)

		for _, letter := range ids[i] {
			letters[letter]++
		}

		for _, count := range letters {
			if count == 2 {
				twos++
				break
			}
		}

		for _, count := range letters {
			if count == 3 {
				threes++
				break
			}
		}
	}

	return strconv.Itoa(threes * twos), nil
}

func SolvePart2(input string) (string, error) {
	boxes := utils.ToStringSlice(input, "\n")
	ids := make([][]rune, len(boxes))

	for i, box := range boxes {
		ids[i] = []rune(box)
	}

	idLen := len(ids[0])
	seen := make([][]rune, 0)

	for _, id := range ids {
		for i := range id {
			pair := make([]rune, idLen)
			copy(pair, id)
			pair[i] = 0

		MatchLoop:
			for _, possibility := range seen {
				for j := range possibility {
					if pair[j] != possibility[j] {
						continue MatchLoop
					}
				}

				return string(pair[0:i]) + string(pair[i+1:]), nil
			}

			seen = append(seen, pair)
		}
	}

	return "", errors.New("unable to solve with the provided input")
}
