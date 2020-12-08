package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/intcode"
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

func SolvePart1(input string) (string, error) {
	instructions, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	pc := intcode.NewComputer(instructions)

	err = pc.Execute(1)
	output := pc.Output()

	return strconv.Itoa(output[len(output)-1]), err
}

func SolvePart2(input string) (string, error) {
	instructions, err := utils.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	pc := intcode.NewComputer(instructions)

	err = pc.Execute(5)
	output := pc.Output()

	return strconv.Itoa(output[len(output)-1]), err
}
