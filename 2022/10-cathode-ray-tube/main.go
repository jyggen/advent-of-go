package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"math"
	"os"
	"strconv"
	"strings"
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
	instructions := utils.ToStringSlice(input, "\n")
	cycle := 0
	register := 1
	sum := 0
	idx := 0

	for {
		var skip int
		var value int

		switch instructions[idx][0:4] {
		case "noop":
			value = math.MaxInt
			skip = 1
		case "addx":
			value, _ = strconv.Atoi(instructions[idx][5:])
			skip = 2
		}

		for ; skip > 0; skip-- {
			cycle++

			if cycle%40 == 20 {
				sum += cycle * register
			}

			if cycle == 220 {
				return strconv.Itoa(sum), nil
			}
		}

		if value != math.MaxInt {
			register += value
		}

		idx = (idx + 1) % len(instructions)
	}
}

func SolvePart2(input string) (string, error) {
	instructions := utils.ToStringSlice(input, "\n")
	cycle := 0
	register := 1
	idx := 0

	var display [6][40]rune

	for {
		var skip int
		var value int

		switch instructions[idx][0:4] {
		case "noop":
			value = math.MaxInt
			skip = 1
		case "addx":
			value, _ = strconv.Atoi(instructions[idx][5:])
			skip = 2
		}

		for ; skip > 0; skip-- {
			cycle++
			x := (cycle - 1) % 40
			y := (cycle - 1) / 40

			if x >= register-1 && x <= register+1 {
				display[y][x] = '#'
			} else {
				display[y][x] = '.'
			}

			if cycle == 240 {
				var b strings.Builder

				for _, rows := range display {
					for _, column := range rows {
						b.WriteRune(column)
					}

					b.WriteString("\n")
				}

				return b.String(), nil
			}
		}

		if value != math.MaxInt {
			register += value
		}

		idx = (idx + 1) % len(instructions)
	}
}
