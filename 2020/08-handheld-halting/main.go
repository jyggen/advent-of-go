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
		if gb.Lookahead().Visits() == 1 {
			break
		}

		gb.Step()
	}

	return strconv.Itoa(gb.Accumulator()), nil
}

func SolvePart2(input string) (string, error) {
	gb := gameboy.New(input)
	opcodes := gb.Opcodes()
	opLen := len(opcodes)
	i := opLen

	for {
		i--
		op := opcodes[i]

		if op.Kind() == gameboy.Jmp && op.Value() < 0 {
			break
		}
	}

BruteLoop:
	for {
		op := opcodes[i]

		if op.Kind() == gameboy.Jmp {
			op.SetKind(gameboy.Nop)
		} else if op.Kind() == gameboy.Nop {
			op.SetKind(gameboy.Jmp)
		}

		gb.Reset()
		for {
			if gb.Lookahead() == nil {
				break BruteLoop
			}

			if gb.Lookahead().Visits() == 1 {
				break
			}

			gb.Step()
		}

		if op.Kind() == gameboy.Jmp {
			op.SetKind(gameboy.Nop)
		} else if op.Kind() == gameboy.Nop {
			op.SetKind(gameboy.Jmp)
		}

		i--
	}

	return strconv.Itoa(gb.Accumulator()), nil
}
