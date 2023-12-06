package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func solve(time int, distance int) int {
	from := sort.Search(time, func(i int) bool {
		return i*(time-i) > distance
	})

	to := sort.Search(time, func(i int) bool {
		return i*(time-i) <= distance
	})

	return to - from
}

func SolvePart1(input string) (string, error) {
	numbers := utils.ToOptimisticIntSlice(input, false)
	half := len(numbers) / 2
	sum := 1

	for i := 0; i < half; i++ {
		sum *= solve(numbers[i], numbers[half+i])
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	parts := utils.ToStringSlice(input, "\n")

	var numbers [2]int

	for i, p := range parts {
		var numberString strings.Builder

		for _, c := range p {
			if unicode.IsDigit(c) {
				numberString.WriteRune(c)
			}
		}

		number, _ := strconv.Atoi(numberString.String())
		numbers[i] = number
	}

	return strconv.Itoa(solve(numbers[0], numbers[1])), nil
}
