package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/intcode"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
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

	pc := intcode.NewComputer(instructions)

	pc.SetValue(1, 12)
	pc.SetValue(2, 2)

	pcInput := make(chan int, 1)
	pcOutput := make(chan int, 1)

	go pc.Execute(pcInput, pcOutput)

	pcInput <- 0
	close(pcInput)
	<-pcOutput

	return strconv.Itoa(pc.Value(0)), nil
}

func SolvePart2(input string) (string, error) {
	instructions, err := utils.ToIntegerSlice(input, ",")
	if err != nil {
		return "", err
	}

	pc := intcode.NewComputer(instructions)
	noun := 0
	verb := 0

	for result := 0; result <= target; noun++ {
		pc.Reset()
		pc.SetValue(1, noun)
		pc.SetValue(2, verb)

		pcInput := make(chan int, 1)
		pcOutput := make(chan int, 1)

		go pc.Execute(pcInput, pcOutput)

		pcInput <- 0
		close(pcInput)
		<-pcOutput

		result = pc.Value(0)
	}

	noun -= 2

	for result := 0; result <= target; verb++ {
		pc.Reset()
		pc.SetValue(1, noun)
		pc.SetValue(2, verb)

		pcInput := make(chan int, 1)
		pcOutput := make(chan int, 1)

		go pc.Execute(pcInput, pcOutput)

		pcInput <- 0
		close(pcInput)
		<-pcOutput

		result = pc.Value(0)
	}

	verb -= 2

	return strconv.Itoa(100*noun + verb), nil
}
