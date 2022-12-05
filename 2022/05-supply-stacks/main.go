package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
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

func parseInput(input string) ([][]uint8, []int) {
	sections := strings.Split(input, "\n\n")
	rows := strings.Split(sections[0], "\n")
	columnCount := (len(rows[0]) - 2) / 3
	columns := make([][]uint8, columnCount)

	for i := range columns {
		columns[i] = make([]uint8, 0)
	}

	for _, row := range rows[:len(rows)-1] {
		for i, column := 1, 0; i < len(row); i, column = i+4, column+1 {
			if row[i] != ' ' {
				columns[column] = append(columns[column], row[i])
			}
		}
	}

	return columns, utils.ToOptimisticIntSlice(sections[1])
}

func arrangeCrates(columns [][]uint8, commands []int, stacked bool) string {
	for i := 0; i < len(commands); i += 3 {
		count, from, to := commands[i], commands[i+1]-1, commands[i+2]-1
		movedCrates := make([]uint8, len(columns[from][:count]))

		copy(movedCrates, columns[from][:count])

		if !stacked {
			for j, k := 0, len(movedCrates)-1; j < k; j, k = j+1, k-1 {
				movedCrates[j], movedCrates[k] = movedCrates[k], movedCrates[j]
			}
		}

		columns[from] = columns[from][count:]
		columns[to] = append(movedCrates, columns[to]...)
	}

	result := make([]uint8, 0, len(columns))

	for _, column := range columns {
		if len(column) != 0 {
			result = append(result, column[0])
		}
	}

	return string(result)
}

func SolvePart1(input string) (string, error) {
	columns, commands := parseInput(input)

	return arrangeCrates(columns, commands, false), nil
}

func SolvePart2(input string) (string, error) {
	columns, commands := parseInput(input)

	return arrangeCrates(columns, commands, true), nil
}
