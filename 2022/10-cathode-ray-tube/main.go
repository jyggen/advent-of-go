package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
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

func runClockCircuit(input string, maxCycles int) []int {
	instructions := utils.ToStringSlice(input, "\n")
	cycleLookup := make([]int, maxCycles+1)
	cycle, register := 0, 1

	for _, instruction := range instructions {
		cycle++
		cycleLookup[cycle] = register

		if instruction[0:4] == "addx" {
			cycle++
			cycleLookup[cycle] = register
			value, _ := strconv.Atoi(instruction[5:])
			register += value
		}

		if cycle == maxCycles {
			return cycleLookup
		}
	}

	return cycleLookup
}

func SolvePart1(input string) (string, error) {
	cycleLookup := runClockCircuit(input, 220)
	sum := 0

	for i := 20; i <= 220; i += 40 {
		sum += i * cycleLookup[i]
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	cycleLookup := runClockCircuit(input, 240)

	var b strings.Builder

	for i := 1; i <= 240; i++ {
		x := (i - 1) % 40

		if x >= cycleLookup[i]-1 && x <= cycleLookup[i]+1 {
			b.WriteRune('#')
		} else {
			b.WriteRune('.')
		}

		if x == 39 {
			b.WriteString("\n")
		}
	}

	return b.String(), nil
}
