package main

import (
	"fmt"
	intcode2 "github.com/jyggen/advent-of-go/internal/intcode"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	instructions, err := utils2.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	pc := intcode2.NewComputer(instructions)
	pcInput := make(chan int, 1)
	pcOutput := make(chan int, 1)

	go pc.Execute(pcInput, pcOutput)

	pcInput <- 1

	return strconv.Itoa(<-pcOutput), err
}

func SolvePart2(input string) (string, error) {
	instructions, err := utils2.ToIntegerSlice(input, ",")

	if err != nil {
		return "", err
	}

	pc := intcode2.NewComputer(instructions)
	pcInput := make(chan int, 1)
	pcOutput := make(chan int, 1)

	go pc.Execute(pcInput, pcOutput)

	pcInput <- 2

	return strconv.Itoa(<-pcOutput), err
}
