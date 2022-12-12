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
	operation   operation
	test        int
	ifTrue      int
	ifFalse     int
}

type operation func(i int) int

type constantOp struct {
	constant int
}

func newConstantOp(op string) constantOp {
	constant, _ := strconv.Atoi(op)

	return constantOp{constant}
}

func (c constantOp) addItemWithConstant(i int) int {
	return i + c.constant
}

func (c constantOp) multiplyItemWithConstant(i int) int {
	return i * c.constant
}

func addItemWithItem(i int) int {
	return i + i
}

func multiplyItemWithItem(i int) int {
	return i * i
}

func parseMonkeys(input string) []*monkey {
	sections := utils.ToStringSlice(input, "\n\n")
	monkeys := make([]*monkey, len(sections))

	for i, section := range sections {
		rows := utils.ToStringSlice(section, "\n")
		items, _ := utils.ToIntegerSlice(rows[1][18:], ", ")
		op := utils.ToStringSlice(rows[2][19:], " ")
		test, _ := strconv.Atoi(rows[3][21:])
		ifTrue, _ := strconv.Atoi(rows[4][29:])
		ifFalse, _ := strconv.Atoi(rows[5][30:])

		var opFunc operation

		switch {
		case op[0] == "old" && op[2] == "old" && op[1] == "+":
			opFunc = addItemWithItem
		case op[0] == "old" && op[2] == "old" && op[1] == "*":
			opFunc = multiplyItemWithItem
		case op[0] == "old" && op[2] != "old" && op[1] == "+":
			opFunc = newConstantOp(op[2]).addItemWithConstant
		case op[0] == "old" && op[2] != "old" && op[1] == "*":
			opFunc = newConstantOp(op[2]).multiplyItemWithConstant
		case op[0] != "old" && op[2] == "old" && op[1] == "+":
			opFunc = newConstantOp(op[0]).addItemWithConstant
		case op[0] != "old" && op[2] == "old" && op[1] == "*":
			opFunc = newConstantOp(op[0]).multiplyItemWithConstant
		default:
			panic(fmt.Sprintf("unable to parse operation: %s", rows[2][19:]))
		}

		monkeys[i] = &monkey{
			items:       items,
			inspections: 0,
			operation:   opFunc,
			test:        test,
			ifTrue:      ifTrue,
			ifFalse:     ifFalse,
		}
	}

	return monkeys
}

func SolvePart1(input string) (string, error) {
	monkeys := parseMonkeys(input)

	var item int

	for i := 1; i <= 20; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				item, m.items = m.items[0], m.items[1:]
				item = m.operation(item) / 3

				if item%m.test == 0 {
					monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, item)
				} else {
					monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, item)
				}

				m.inspections++
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return strconv.Itoa(monkeys[0].inspections * monkeys[1].inspections), nil
}

func SolvePart2(input string) (string, error) {
	monkeys := parseMonkeys(input)
	magicNumber := 1

	for _, m := range monkeys {
		magicNumber *= m.test
	}

	var item int

	for i := 1; i <= 10000; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				item, m.items = m.items[0], m.items[1:]
				item = m.operation(item) % magicNumber

				if item%m.test == 0 {
					monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, item)
				} else {
					monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, item)
				}

				m.inspections++
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return strconv.Itoa(monkeys[0].inspections * monkeys[1].inspections), nil
}
