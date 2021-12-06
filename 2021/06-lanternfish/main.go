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

func growth(days int, init []int) int {
	state := [9]int{}

	for _, v := range init {
		state[v]++
	}

	for i := 0; i < days; i++ {
		state[8], state[7], state[6], state[5], state[4], state[3], state[2], state[1], state[0] = state[0], state[8], state[7]+state[0], state[6], state[5], state[4], state[3], state[2], state[1]
	}

	return state[0] + state[1] + state[2] + state[3] + state[4] + state[5] + state[6] + state[7] + state[8]
}

func SolvePart1(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	return strconv.Itoa(growth(80, intSlice)), nil
}

func SolvePart2(input string) (string, error) {
	intSlice, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	return strconv.Itoa(growth(256, intSlice)), nil
}
