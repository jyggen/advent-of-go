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
	visited := make(map[int]bool)

	for {
		if _, ok := visited[gb.Lookahead()]; ok {
			break
		}

		gb.Step()
		visited[gb.Position()] = true
	}

	return strconv.Itoa(gb.Accumulator()), nil
}

func SolvePart2(input string) (string, error) {
	gb := gameboy.New(input)
	opcodes := gb.Opcodes()
	opLen := len(opcodes)
	i := 0

BruteLoop:
	for {
		i++
		visited := make(map[int]bool)
		op := opcodes[i]

		if op.Kind() == gameboy.Jmp {
			op.SetKind(gameboy.Nop)
		} else if op.Kind() == gameboy.Nop {
			op.SetKind(gameboy.Jmp)
		}

		gb.Reset()

		for {
			if gb.Lookahead() >= opLen {
				break BruteLoop
			}

			if _, ok := visited[gb.Lookahead()]; ok {
				break
			}

			gb.Step()

			visited[gb.Position()] = true
		}

		if op.Kind() == gameboy.Jmp {
			op.SetKind(gameboy.Nop)
		} else if op.Kind() == gameboy.Nop {
			op.SetKind(gameboy.Jmp)
		}
	}

	return strconv.Itoa(gb.Accumulator()), nil
}
