package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	rows := utils2.ToStringSlice(input, "\n")
	instructions := make([][]int, len(rows))

	var err error

	for i, row := range rows {
		instructions[i], err = utils2.ToIntegerSlice(row, "")

		if err != nil {
			return "", err
		}
	}

	epsilon := make([]rune, len(instructions[0]))
	gamma := make([]rune, len(instructions[0]))

	for i := 0; i < len(instructions[0]); i++ {
		ones := 0
		zeros := 0

		for j := 0; j < len(instructions); j++ {
			if instructions[j][i] == 1 {
				ones++
			} else {
				zeros++
			}
		}

		if ones > zeros {
			epsilon[i] = '0'
			gamma[i] = '1'
		} else {
			epsilon[i] = '1'
			gamma[i] = '0'
		}
	}

	epsilonInt, err := strconv.ParseInt(string(epsilon), 2, 64)

	if err != nil {
		return "", fmt.Errorf("unable to parse epsilon: %w", err)
	}

	gammaInt, err := strconv.ParseInt(string(gamma), 2, 64)

	if err != nil {
		return "", fmt.Errorf("unable to parse gamma: %w", err)
	}

	return strconv.Itoa(int(epsilonInt * gammaInt)), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils2.ToStringSlice(input, "\n")
	instructions := make([][]int, len(rows))

	var err error

	for i, row := range rows {
		instructions[i], err = utils2.ToIntegerSlice(row, "")

		if err != nil {
			return "", err
		}
	}

	oxygen := make([][]int, len(instructions))
	co2 := make([][]int, len(instructions))

	copy(oxygen, instructions)
	copy(co2, instructions)

	higher := func(a int, b int) bool { return a >= b }
	lower := func(a int, b int) bool { return a < b }
	parts := []*struct {
		instructions [][]int
		compare      func(int, int) bool
	}{
		{oxygen, higher},
		{co2, lower},
	}

	for _, s := range parts {
		for i := 0; i < len(s.instructions[0]); i++ {
			if len(s.instructions) == 1 {
				break
			}

			ones := 0
			zeros := 0

			for j := 0; j < len(s.instructions); j++ {
				if s.instructions[j][i] == 1 {
					ones++
				} else {
					zeros++
				}
			}

			newInstructions := make([][]int, 0, len(s.instructions))

			for j := 0; j < len(s.instructions); j++ {
				if s.compare(ones, zeros) {
					if s.instructions[j][i] == 1 {
						newInstructions = append(newInstructions, s.instructions[j])
					}
				} else {
					if s.instructions[j][i] == 0 {
						newInstructions = append(newInstructions, s.instructions[j])
					}
				}
			}

			s.instructions = newInstructions
		}
	}

	oxygenInt, err := strconv.ParseInt(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(parts[0].instructions[0])), ""), "[]"), 2, 64)

	if err != nil {
		return "", fmt.Errorf("unable to parse oxygen: %w", err)
	}

	co2Int, err := strconv.ParseInt(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(parts[1].instructions[0])), ""), "[]"), 2, 64)

	if err != nil {
		return "", fmt.Errorf("unable to parse co2: %w", err)
	}

	return strconv.Itoa(int(oxygenInt * co2Int)), nil
}
