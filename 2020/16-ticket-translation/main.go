package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"
)

type field struct {
	name  string
	rules [2]rule
}

type rule struct {
	lower int
	upper int
}

type ticket struct {
	values []int
}

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	_, tickets, fields := parseInput(input)
	_, errorRate := validateTickets(tickets, fields)

	return strconv.Itoa(errorRate), nil
}

func SolvePart2(input string) (string, error) {
	myTicket, tickets, fields := parseInput(input)
	tickets, _ = validateTickets(tickets, fields)
	possibilities := make([][]int, len(myTicket.values))
	fieldLen := len(fields)

	for i := range possibilities {
		possibilities[i] = make([]int, fieldLen)

		for j := range fields {
			possibilities[i][j] = j
		}
	}

	for _, t := range append(tickets, myTicket) {
		for i, v := range t.values {
			for j := 0; j < len(possibilities[i]); {
				rules := fields[possibilities[i][j]].rules

				if (v >= rules[0].lower && v <= rules[0].upper) || (v >= rules[1].lower && v <= rules[1].upper) {
					j++
				} else {
					possibilities[i] = remove(possibilities[i], j)
				}
			}
		}
	}

	for i := 0; i < len(possibilities); {
		if len(possibilities[i]) == 1 {
			changes := 0

			for j := range possibilities {
				if j == i {
					continue
				}

				index := -1

				for k, v := range possibilities[j] {
					if v == possibilities[i][0] {
						index = k
						break
					}
				}

				if index != -1 {
					possibilities[j] = remove(possibilities[j], index)
					changes++
				}
			}

			if changes > 0 {
				i = 0
			} else {
				i++
			}
		} else {
			i++
		}
	}

	result := 1

	for k, v := range myTicket.values {
		f := fields[possibilities[k][0]]

		if !strings.HasPrefix(f.name, "departure") {
			continue
		}

		result = result * v
	}

	return strconv.Itoa(result), nil
}

func parseInput(input string) (ticket, []ticket, []field) {
	sections := utils2.ToStringSlice(input, "\n\n")
	fields := parseField(sections[0])
	myTicket := parseTicket(utils2.ToStringSlice(sections[1], "\n")[1])
	rows := utils2.ToStringSlice(sections[2], "\n")[1:]
	tickets := make([]ticket, len(rows))

	for i, r := range rows {
		tickets[i] = parseTicket(r)
	}

	return myTicket, tickets, fields
}

func parseField(input string) []field {
	rows := utils2.ToStringSlice(input, "\n")
	fields := make([]field, len(rows))

	for i, row := range rows {
		parts := utils2.ToStringSlice(row, ": ")
		rules := [2]rule{}

		for k, r := range utils2.ToStringSlice(parts[1], " or ") {
			bounds, _ := utils2.ToIntegerSlice(r, "-")
			rules[k] = rule{
				lower: bounds[0],
				upper: bounds[1],
			}
		}

		fields[i] = field{
			name:  parts[0],
			rules: rules,
		}
	}

	return fields
}

func parseTicket(input string) ticket {
	values, _ := utils2.ToIntegerSlice(input, ",")

	return ticket{
		values: values,
	}
}

func validateTickets(tickets []ticket, fields []field) ([]ticket, int) {
	validTickets := make([]ticket, 0)
	errorRate := 0

	for _, t := range tickets {
		isValidTicket := true

		for _, v := range t.values {
			isValidValue := false

		RulesLoop:
			for _, f := range fields {
				for _, r := range f.rules {
					if v >= r.lower && v <= r.upper {
						isValidValue = true
						break RulesLoop
					}
				}
			}

			if !isValidValue {
				isValidTicket = false
				errorRate += v
			}
		}

		if isValidTicket {
			validTickets = append(validTickets, t)
		}
	}

	return validTickets, errorRate
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]

	return s[:len(s)-1]
}
