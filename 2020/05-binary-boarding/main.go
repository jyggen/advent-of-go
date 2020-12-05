package main

import (
	"errors"
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"sort"
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
	tickets := utils.ToStringSlice(input, "\n")
	ids, err := decodeTickets(tickets)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(ids[len(ids)-1]), nil
}

func SolvePart2(input string) (string, error) {
	tickets := utils.ToStringSlice(input, "\n")
	ids, err := decodeTickets(tickets)

	if err != nil {
		return "", err
	}

	numIds := len(ids)

	for i := 1; i < numIds; i++ {
		if ids[i-1] != ids[i]-1 {
			return strconv.Itoa(ids[i] - 1), nil
		}
	}

	return "", errors.New("unable to find ticket number")
}

func decodeTickets(tickets []string) ([]int, error) {
	ids := make([]int, len(tickets))
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	for i, ticket := range tickets {
		id, err := strconv.ParseInt(replacer.Replace(ticket), 2, 64)

		if err != nil {
			return ids, err
		}

		ids[i] = int(id)
	}

	sort.Ints(ids)

	return ids, nil
}
