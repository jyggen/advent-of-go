package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"
)

type game struct {
	highBlue  int
	highGreen int
	highRed   int
	number    int
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func parse(rows []string) []game {
	games := make([]game, 0, len(rows))

	for _, row := range rows {
		colonIndex := strings.Index(row, ": ")
		gameNo, _ := strconv.Atoi(row[5:colonIndex])
		g := game{number: gameNo}
		cubes := strings.Split(row[colonIndex+2:], " ")

		for i := 0; i < len(cubes); i += 2 {
			count, _ := strconv.Atoi(cubes[i])

			switch cubes[i+1] {
			case "blue", "blue,", "blue;":
				if count > g.highBlue {
					g.highBlue = count
				}
			case "green", "green,", "green;":
				if count > g.highGreen {
					g.highGreen = count
				}
			case "red", "red,", "red;":
				if count > g.highRed {
					g.highRed = count
				}
			}
		}

		games = append(games, g)
	}

	return games
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	games := parse(rows)
	valid := 0

	for _, g := range games {
		if g.highBlue <= 14 && g.highGreen <= 13 && g.highRed <= 12 {
			valid += g.number
		}
	}

	return strconv.Itoa(valid), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	games := parse(rows)
	sum := 0

	for _, g := range games {
		sum += g.highBlue * g.highGreen * g.highRed
	}

	return strconv.Itoa(sum), nil
}
