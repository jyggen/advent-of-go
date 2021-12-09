package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gitchander/permutation"
	"github.com/jyggen/advent-of-go/internal/intcode"
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
	instructions, err := utils.ToIntegerSlice(input, ",")
	if err != nil {
		return "", err
	}

	best := 0
	sequence := []int{0, 1, 2, 3, 4}
	p := permutation.New(permutation.IntSlice(sequence))

	for p.Next() {
		in := 0

		for i := 0; i < 5; i++ {
			pc := intcode.NewComputer(instructions)
			pcInput := make(chan int, 1)
			pcOutput := make(chan int, 1)

			go pc.Execute(pcInput, pcOutput)

			pcInput <- sequence[i]
			pcInput <- in
			close(pcInput)
			in = <-pcOutput
		}

		if in > best {
			best = in
		}
	}

	return strconv.Itoa(best), nil
}

type metadata struct {
	pc     *intcode.Computer
	input  chan int
	output chan int
}

func SolvePart2(input string) (string, error) {
	instructions, err := utils.ToIntegerSlice(input, ",")
	if err != nil {
		return "", err
	}

	best := 0
	sequence := []int{5, 6, 7, 8, 9}
	p := permutation.New(permutation.IntSlice(sequence))

	for p.Next() {
		pcs := make([]*metadata, 5)

		for i := 0; i < 5; i++ {
			pcs[i] = &metadata{
				intcode.NewComputer(instructions),
				make(chan int, 1),
				make(chan int, 1),
			}
		}

		for i := 0; i < 5; i++ {
			pcs[i].input = pcs[(5+i-1)%5].output
			pcs[i].output = pcs[(i+1)%5].input
		}

		pcs[4].output = make(chan int, 1)

		for i, pc := range pcs {
			go pc.pc.Execute(pc.input, pc.output)

			pc.input <- sequence[i]
		}

		pcs[0].input <- 0

		for output := range pcs[4].output {
			pcs[0].input <- output

			if output > best {
				best = output
			}
		}
	}

	return strconv.Itoa(best), nil
}
