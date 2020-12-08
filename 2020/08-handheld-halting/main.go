package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/gameboy"
	"github.com/jyggen/advent-of-go/solver"
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
	gb := gameboy.New(input)

	for {
		gb.Step()

		if gb.Lookahead().Visits() == 1 {
			return strconv.Itoa(gb.Accumulator()), nil
		}
	}
}

func SolvePart2(input string) (string, error) {
	gb := gameboy.New(input)
	opcodes := gb.Opcodes()
	opLen := len(opcodes)
	i := 0

BruteLoop:
	for {
		i++
		op := opcodes[i]

		if op.Kind() == gameboy.Jmp {
			op.SetKind(gameboy.Nop)
		} else if op.Kind() == gameboy.Nop {
			op.SetKind(gameboy.Jmp)
		}

		gb.Reset()

		for {
			gb.Step()

			if gb.Position() + 1 >= opLen {
				break BruteLoop
			}

			if gb.Lookahead().Visits() == i {
				break
			}
		}

		if op.Kind() == gameboy.Jmp {
			op.SetKind(gameboy.Nop)
		} else if op.Kind() == gameboy.Nop {
			op.SetKind(gameboy.Jmp)
		}
	}

	return strconv.Itoa(gb.Accumulator()), nil
}
