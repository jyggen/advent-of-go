package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

func SolvePart1(input string) (string, error) {
	instructions := utils.ToStringSlice(input, "\n")

	var position [2]int

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		value, err := strconv.Atoi(parts[1])

		if err != nil {
			return "", err
		}

		switch parts[0] {
		case "forward":
			position[1] += value
		case "up":
			position[0] -= value
		case "down":
			position[0] += value
		}
	}

	return strconv.Itoa(position[0] * position[1]), nil
}

func SolvePart2(input string) (string, error) {
	instructions := utils.ToStringSlice(input, "\n")

	var position [3]int

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		value, err := strconv.Atoi(parts[1])

		if err != nil {
			return "", err
		}

		switch parts[0] {
		case "forward":
			position[0] += value
			position[1] += position[2] * value
		case "up":
			position[2] -= value
		case "down":
			position[2] += value
		}
	}

	return strconv.Itoa(position[0] * position[1]), nil
}
