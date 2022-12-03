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

func SolvePart1(input string) (string, error) {
	rucksacks := utils.ToRuneSlice(input, "\n")
	priorities := 0

	for _, rucksack := range rucksacks {
		first := rucksack[:len(rucksack)/2]
		second := rucksack[len(rucksack)/2:]

	Loop:
		for _, f := range first {
			for _, s := range second {
				if f == s {
					priorities += int(f)

					if f >= 'a' {
						priorities -= 96
					} else {
						priorities -= 38
					}
					break Loop
				}
			}
		}
	}

	return strconv.Itoa(priorities), nil
}

func SolvePart2(input string) (string, error) {
	rucksacks := utils.ToRuneSlice(input, "\n")
	priorities := 0

	for i := 0; i < len(rucksacks); i += 3 {
	Loop:
		for _, f := range rucksacks[i] {
			for _, s := range rucksacks[i+1] {
				for _, t := range rucksacks[i+2] {
					if f == s && f == t {
						priorities += int(f)

						if f >= 'a' {
							priorities -= 96
						} else {
							priorities -= 38
						}
						break Loop
					}
				}
			}
		}
	}

	return strconv.Itoa(priorities), nil
}
