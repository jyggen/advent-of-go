package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"
)

var cubeLimits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	valid := 0

	for _, r := range rows {
		parts := strings.SplitN(r, ": ", 2)
		sets := strings.Split(parts[1], "; ")
		ok := true

	GameLoop:
		for _, s := range sets {
			cubes := strings.Split(s, ", ")

			for _, c := range cubes {
				details := strings.SplitN(c, " ", 2)
				count, err := strconv.Atoi(details[0])

				if err != nil {
					return "", err
				}

				if cubeLimits[details[1]] < count {
					ok = false
					break GameLoop
				}
			}
		}

		if ok {
			gameNo, err := strconv.Atoi(parts[0][5:])

			if err != nil {
				return "", err
			}

			valid += gameNo
		}
	}

	return strconv.Itoa(valid), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	sum := 0

	for _, r := range rows {
		parts := strings.SplitN(r, ": ", 2)
		sets := strings.Split(parts[1], "; ")
		lowest := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}

		for _, s := range sets {
			cubes := strings.Split(s, ", ")

			for _, c := range cubes {
				details := strings.SplitN(c, " ", 2)
				count, err := strconv.Atoi(details[0])

				if err != nil {
					return "", err
				}

				current := lowest[details[1]]

				if current == 0 || current < count {
					lowest[details[1]] = count
				}
			}
		}

		sum += lowest["blue"] * lowest["green"] * lowest["red"]
	}

	return strconv.Itoa(sum), nil
}
