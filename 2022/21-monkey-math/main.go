package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
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
	dependencies []*monkey
	operation    func() int
	number       int
}

func (m *monkey) add() int {
	return m.dependencies[0].operation() + m.dependencies[1].operation()
}

func (m *monkey) subtract() int {
	return m.dependencies[0].operation() - m.dependencies[1].operation()
}

func (m *monkey) multiply() int {
	return m.dependencies[0].operation() * m.dependencies[1].operation()
}

func (m *monkey) divide() int {
	return m.dependencies[0].operation() / m.dependencies[1].operation()
}

func (m *monkey) say() int {
	return m.number
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	lookup := make(map[string]*monkey, len(rows))
	monkeys := make([]*monkey, len(rows))

	for i, r := range rows {
		name := r[:4]
		m := &monkey{
			dependencies: make([]*monkey, 0),
		}
		monkeys[i] = m
		lookup[name] = m
	}

	for i, r := range rows {
		r = r[6:]
		m := monkeys[i]
		switch {
		case strings.ContainsRune(r, '+'):
			m.dependencies = append(m.dependencies, lookup[r[0:4]], lookup[r[7:]])
			m.operation = m.add
		case strings.ContainsRune(r, '/'):
			m.dependencies = append(m.dependencies, lookup[r[0:4]], lookup[r[7:]])
			m.operation = m.divide
		case strings.ContainsRune(r, '*'):
			m.dependencies = append(m.dependencies, lookup[r[0:4]], lookup[r[7:]])
			m.operation = m.multiply
		case strings.ContainsRune(r, '-'):
			m.dependencies = append(m.dependencies, lookup[r[0:4]], lookup[r[7:]])
			m.operation = m.subtract
		default:
			m.number, _ = strconv.Atoi(r)
			m.operation = m.say
		}
	}

	return strconv.Itoa(lookup["root"].operation()), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	lookup := make(map[string]*monkey, len(rows))
	monkeys := make([]*monkey, len(rows))

	for i, r := range rows {
		name := r[:4]
		m := &monkey{
			dependencies: make([]*monkey, 0),
		}
		monkeys[i] = m
		lookup[name] = m
	}

	for i, r := range rows {
		r = r[6:]
		m := monkeys[i]
		switch {
		case strings.ContainsRune(r, '+'):
			m.dependencies = append(m.dependencies, lookup[r[0:4]], lookup[r[7:]])
			m.operation = m.add
		case strings.ContainsRune(r, '/'):
			m.dependencies = append(m.dependencies, lookup[r[0:4]], lookup[r[7:]])
			m.operation = m.divide
		case strings.ContainsRune(r, '*'):
			m.dependencies = append(m.dependencies, lookup[r[0:4]], lookup[r[7:]])
			m.operation = m.multiply
		case strings.ContainsRune(r, '-'):
			m.dependencies = append(m.dependencies, lookup[r[0:4]], lookup[r[7:]])
			m.operation = m.subtract
		default:
			m.number, _ = strconv.Atoi(r)
			m.operation = m.say
		}
	}

	i := 1
	diff := -1
	increase := true
	static := lookup["root"].dependencies[1].operation()
	amount := static

	for {
		lookup["humn"].number = i
		a := lookup["root"].dependencies[0].operation()

		if a == static {
			break
		}

		newDiff := utils.AbsInt(a - static)

		if diff != -1 && newDiff > diff {
			if increase {
				increase = false
			} else {
				increase = true
			}

			amount /= 2
		}

		diff = newDiff

		if increase {
			i += amount
		} else {
			i -= amount
		}
	}

	for lookup["root"].dependencies[0].operation() == static {
		lookup["humn"].number--
	}

	return strconv.Itoa(lookup["humn"].number + 1), nil
}
