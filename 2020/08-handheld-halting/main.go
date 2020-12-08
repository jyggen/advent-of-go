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

// non-bruteforce solution based on /u/smmalis37.
func SolvePart2(input string) (string, error) {
	gb := gameboy.New(input)

	for {
		gb.Step()

		if gb.Lookahead().Visits() == 1 {
			break
		}
	}

	opcodes := gb.Opcodes()
	potential := make(map[int]bool, 0)
	i := len(opcodes)

	for {
		potential[i] = true
		i--

		if opcodes[i].Kind() == gameboy.Jmp && opcodes[i].Value() < 0 {
			break
		}
	}

	start := i

	var swap int

	if opcodes[i].Visits() > 0 {
		swap = i
	} else {
		for {
			i -= 1

			if potential[i] {
				continue
			}

			if opcodes[i].Kind() == gameboy.Nop && opcodes[i].Visits() > 0 {
				if _, ok := potential[i+opcodes[i].Value()]; ok {
					swap = i
					break
				}
			} else if opcodes[i].Kind() == gameboy.Jmp && opcodes[i].Visits() == 0 {
				if _, ok := potential[i+opcodes[i].Value()]; !ok {
					continue
				}

				if _, ok := potential[i]; ok {
					continue
				}

				j := i - 1

				for {
					if opcodes[j].Kind() == gameboy.Jmp {
						break
					}

					j--
				}

				if opcodes[j].Visits() > 0 {
					swap = j
					break
				}

				for j = j + 1; j <= i; j++ {
					potential[j] = true
				}

				i = start
			}
		}
	}

	op := opcodes[swap]

	if op.Kind() == gameboy.Nop {
		op.SetKind(gameboy.Jmp)
	} else {
		op.SetKind(gameboy.Nop)
	}

	gb.Reset()
	gb.Run()

	return strconv.Itoa(gb.Accumulator()), nil
}
