package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"
)

type round struct {
	blue  int
	green int
	red   int
}

type game struct {
	number int
	rounds []round
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
		parts := strings.SplitN(row, ": ", 2)
		sets := strings.Split(parts[1], "; ")
		gameNo, _ := strconv.Atoi(parts[0][5:])
		g := game{number: gameNo, rounds: make([]round, 0, len(sets))}

		for _, s := range sets {
			cubes := strings.Split(s, ", ")
			r := round{}

			for _, c := range cubes {
				details := strings.SplitN(c, " ", 2)
				count, _ := strconv.Atoi(details[0])

				switch details[1] {
				case "blue":
					r.blue = count
				case "green":
					r.green = count
				case "red":
					r.red = count
				}
			}

			g.rounds = append(g.rounds, r)
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
		ok := true
		for _, r := range g.rounds {
			if r.blue > 14 || r.green > 13 || r.red > 12 {
				ok = false
				break
			}
		}

		if ok {
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
		var highBlue, highGreen, highRed int

		for _, r := range g.rounds {
			if highBlue < r.blue {
				highBlue = r.blue
			}

			if highGreen < r.green {
				highGreen = r.green
			}

			if highRed < r.red {
				highRed = r.red
			}
		}

		sum += highBlue * highGreen * highRed
	}

	return strconv.Itoa(sum), nil
}
