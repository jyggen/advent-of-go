package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"sort"
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

type monkey struct {
	items       []int
	inspections int
	operation   func(int) int
	test        func(int) int
}

func simulateMonkeys(input string, rounds int, manageWorry func(int) int) int {
	sections := utils.ToStringSlice(input, "\n\n")
	monkeys := make([]*monkey, len(sections))

	for i, section := range sections {
		rows := utils.ToStringSlice(section, "\n")
		items, _ := utils.ToIntegerSlice(rows[1][18:], ", ")
		operation := utils.ToStringSlice(rows[2][19:], " ")
		test, _ := strconv.Atoi(rows[3][21:])
		ifTrue, _ := strconv.Atoi(rows[4][29:])
		ifFalse, _ := strconv.Atoi(rows[5][30:])
		monkeys[i] = &monkey{
			items:       items,
			inspections: 0,
			operation: func(i int) int {
				a, b := 0, 0

				if operation[0] == "old" {
					a = i
				} else {
					a, _ = strconv.Atoi(operation[0])
				}

				if operation[2] == "old" {
					b = i
				} else {
					b, _ = strconv.Atoi(operation[2])
				}

				if operation[1] == "+" {
					return a + b
				} else {
					return a * b
				}
			},
			test: func(i int) int {
				if i%test == 0 {
					return ifTrue
				} else {
					return ifFalse
				}
			},
		}
	}

	var item int

	for i := 1; i <= rounds; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				item, m.items = m.items[0], m.items[1:]
				item = m.operation(item)
				item = manageWorry(item)
				to := m.test(item)

				monkeys[to].items = append(monkeys[to].items, item)
				m.inspections++
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return monkeys[0].inspections * monkeys[1].inspections
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(simulateMonkeys(input, 20, func(i int) int {
		return i / 3
	})), nil
}

func SolvePart2(input string) (string, error) {
	sections := utils.ToStringSlice(input, "\n\n")
	magicNumber := 1

	for _, section := range sections {
		rows := utils.ToStringSlice(section, "\n")
		test, _ := strconv.Atoi(rows[3][21:])
		magicNumber *= test

	}

	return strconv.Itoa(simulateMonkeys(input, 10000, func(i int) int {
		return i % magicNumber
	})), nil
}
