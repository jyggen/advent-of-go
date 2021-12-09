package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

var binaryReplacer *strings.Replacer

func init() {
	binaryReplacer = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	ids, length, err := decodeTickets(input)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(ids[length-1]), nil
}

func SolvePart2(input string) (string, error) {
	ids, length, err := decodeTickets(input)
	if err != nil {
		return "", err
	}

	for i := 1; i < length; i++ {
		if ids[i-1] != ids[i]-1 {
			return strconv.Itoa(ids[i] - 1), nil
		}
	}

	return "", errors.New("unable to find ticket number")
}

func decodeTickets(input string) ([]int, int, error) {
	tickets := utils.ToStringSlice(binaryReplacer.Replace(input), "\n")
	idLen := len(tickets)
	ids := make([]int, idLen)

	for i, ticket := range tickets {
		id, err := strconv.ParseInt(ticket, 2, 64)
		if err != nil {
			return ids, idLen, err
		}

		ids[i] = int(id)
	}

	sort.Ints(ids)

	return ids, idLen, nil
}
