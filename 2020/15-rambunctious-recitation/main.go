package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

const p1rounds = 2020
const p2rounds = 30000000

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(solve(input, p1rounds)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(solve(input, p2rounds)), nil
}

func solve(input string, rounds int) int {
	numbers, _ := utils2.ToIntegerSlice(input, ",")
	memory := make(map[int]int, rounds)
	round := len(numbers)

	for k, v := range numbers {
		memory[v] = k
	}

	lastSpoken := numbers[round-1]

	for i := round; i < rounds; i++ {
		var newNumber int

		if n, ok := memory[lastSpoken]; ok {
			newNumber = i - 1 - n
		} else {
			newNumber = 0
		}

		memory[lastSpoken] = i - 1
		lastSpoken = newNumber
	}

	return lastSpoken
}
