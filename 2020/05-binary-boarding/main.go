package main

import (
	"errors"
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
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

func SolvePart1(input string) (string, error) {
	tickets := utils.ToStringSlice(input, "\n")
	ids := decodeTickets(tickets)

	return strconv.Itoa(ids[len(ids)-1]), nil
}

func SolvePart2(input string) (string, error) {
	tickets := utils.ToStringSlice(input, "\n")
	ids := decodeTickets(tickets)
	numIds := len(ids)

	for i := 1; i < numIds; i++ {
		if ids[i-1] != ids[i]-1 {
			return strconv.Itoa(ids[i] - 1), nil
		}
	}

	return "", errors.New("unable to find ticket number")
}

func decodeTickets(tickets []string) []int {
	numTickets := len(tickets)
	ids := make([]int, numTickets)

	for i, ticket := range tickets {
		rows := utils.IntRange(0, 128)
		cols := utils.IntRange(0, 7)

		for _, v := range ticket[0:7] {
			switch v {
			case 'F':
				rows = rows[0 : len(rows)/2]
			case 'B':
				rows = rows[len(rows)/2:]
			}
		}

		for _, v := range ticket[7:] {
			switch v {
			case 'L':
				cols = cols[0 : len(cols)/2]
			case 'R':
				cols = cols[len(cols)/2:]
			}
		}

		ids[i] = (rows[0] * 8) + cols[0]
	}

	sort.Ints(ids)

	return ids
}
