package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
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

var lookup map[int]int

func growth(days int, cycle int) int {
	if days <= 0 {
		return 0
	}

	remaining := days - cycle + 7

	if v, ok := lookup[remaining]; ok {
		return v
	}

	count := remaining / 7
	sum := 1

	for i := count; i > 0; i-- {
		sum += growth(remaining-(7*i), 9)
	}

	lookup[remaining] = sum

	return sum
}

func SolvePart1(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	lookup = make(map[int]int)
	sum := 0

	for _, v := range intSlice {
		sum += growth(80, v)
	}

	return strconv.Itoa(sum), err
}

func SolvePart2(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	lookup = make(map[int]int)
	sum := 0

	for _, v := range intSlice {
		sum += growth(256, v)
	}

	fmt.Println()

	return strconv.Itoa(sum), err
}
