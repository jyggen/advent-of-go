package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
)

const target = 19690720

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	instructions, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	result, err := runMachine(instructions, 12, 2)

	return strconv.Itoa(result), err
}

func SolvePart2(input string) (string, error) {
	instructions, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	noun := 0
	verb := 0

	for result := 0; result <= target; noun++ {
		result, err = runMachine(instructions, noun, verb)

		if err != nil {
			return "", err
		}
	}

	noun -= 2
	for result := 0; result <= target; verb++ {
		result, err = runMachine(instructions, noun, verb)

		if err != nil {
			return "", err
		}
	}

	verb -= 2

	return strconv.Itoa(100*noun + verb), nil
}

func runMachine(input []int, noun int, verb int) (int, error) {
	position := 0
	instructions := make([]int, len(input))

	copy(instructions, input)

	instructions[1] = noun
	instructions[2] = verb

	for {
		switch instructions[position] {
		case 1:
			instructions[instructions[position+3]] = instructions[instructions[position+1]] + instructions[instructions[position+2]]
		case 2:
			instructions[instructions[position+3]] = instructions[instructions[position+1]] * instructions[instructions[position+2]]
		case 99:
			return instructions[0], nil
		default:
			return 0, fmt.Errorf("unknown opcode \"%d\"", instructions[position])
		}

		position += 4
	}
}
