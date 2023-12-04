package main

import (
	"errors"
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
	ids := utils.ToRuneSlice(input, "\n")

	for i, id := range ids {
	Outer:
		for _, id2 := range ids[i+1:] {
			hasDiff := false
			index := 0

			for k, v := range id {
				if v != id2[k] {
					if hasDiff {
						continue Outer
					}

					hasDiff = true
					index = k
				}
			}

			return string(id[0:index]) + string(id[index+1:]), nil
		}
	}

	return "", errors.New("unable to solve with the provided input")
}
