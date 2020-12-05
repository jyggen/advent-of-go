package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

type opcode func(registers []int, a int, b int, c int)

var opcodeList = []opcode{
	addi,
	addr,
	bani,
	banr,
	bori,
	borr,
	eqir,
	eqri,
	eqrr,
	gtir,
	gtri,
	gtrr,
	muli,
	mulr,
	seti,
	setr,
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	data := utils.ToStringSlice(input, "\n\n\n")
	samples := data[0]
	partOne := 0
	before := make([]int, 4)
	instructions := make([]int, 4)
	after := make([]int, 4)
	i := 0

	for _, s := range strings.Split(samples, "\n") {
		if s == "" {
			continue
		}

		switch i {
		case 0:
			var a, b, c, d int

			_, err := fmt.Sscanf(s, "Before: [%d, %d, %d, %d]", &a, &b, &c, &d)

			if err != nil {
				panic(err)
			}

			before[0] = a
			before[1] = b
			before[2] = c
			before[3] = d
		case 1:
			var instruction, a, b, c int

			_, err := fmt.Sscanf(s, "%d %d %d %d", &instruction, &a, &b, &c)

			if err != nil {
				panic(err)
			}

			instructions[0] = instruction
			instructions[1] = a
			instructions[2] = b
			instructions[3] = c
		case 2:
			var a, b, c, d int

			_, err := fmt.Sscanf(s, "After:  [%d, %d, %d, %d]", &a, &b, &c, &d)

			if err != nil {
				panic(err)
			}

			after[0] = a
			after[1] = b
			after[2] = c
			after[3] = d
		}

		if i < 2 {
			i++
			continue
		}

		matches := 0

		// Test against all opcodes for part one.
		for _, o := range opcodeList {
			registers := []int{before[0], before[1], before[2], before[3]}

			o(registers, instructions[1], instructions[2], instructions[3])

			if registers[0] == after[0] && registers[1] == after[1] && registers[2] == after[2] && registers[3] == after[3] {
				matches++

				if matches == 3 {
					break
				}
			}
		}

		if matches == 3 {
			partOne++
		}

		i = 0
	}

	return strconv.Itoa(partOne), nil
}

func SolvePart2(input string) (string, error) {
	data := utils.ToStringSlice(input, "\n\n\n")
	samples := data[0]
	before := make([]int, 4)
	instructions := make([]int, 4)
	after := make([]int, 4)
	i := 0
	testProgram := strings.TrimSpace(data[1])
	possibilities := make([][]opcode, len(opcodeList))

	for i := range possibilities {
		possibilities[i] = make([]opcode, len(opcodeList))
		copy(possibilities[i], opcodeList)
	}

	for _, s := range strings.Split(samples, "\n") {
		if s == "" {
			continue
		}

		switch i {
		case 0:
			var a, b, c, d int

			_, err := fmt.Sscanf(s, "Before: [%d, %d, %d, %d]", &a, &b, &c, &d)

			if err != nil {
				panic(err)
			}

			before[0] = a
			before[1] = b
			before[2] = c
			before[3] = d
		case 1:
			var instruction, a, b, c int

			_, err := fmt.Sscanf(s, "%d %d %d %d", &instruction, &a, &b, &c)

			if err != nil {
				panic(err)
			}

			instructions[0] = instruction
			instructions[1] = a
			instructions[2] = b
			instructions[3] = c
		case 2:
			var a, b, c, d int

			_, err := fmt.Sscanf(s, "After:  [%d, %d, %d, %d]", &a, &b, &c, &d)

			if err != nil {
				panic(err)
			}

			after[0] = a
			after[1] = b
			after[2] = c
			after[3] = d
		}

		if i < 2 {
			i++
			continue
		}

		// Narrow down this specific opcode number.
		for j := 0; j < len(possibilities[instructions[0]]); j++ {
			registers := []int{before[0], before[1], before[2], before[3]}

			possibilities[instructions[0]][j](registers, instructions[1], instructions[2], instructions[3])

			if registers[0] != after[0] || registers[1] != after[1] || registers[2] != after[2] || registers[3] != after[3] {
				possibilities[instructions[0]] = append(possibilities[instructions[0]][:j], possibilities[instructions[0]][j+1:]...)
				j--
			}
		}

		i = 0
	}

	assignments := make([]opcode, 16)

	for i := 0; i < len(possibilities); i++ {
		if len(possibilities[i]) != 1 {
			continue
		}

		assignments[i] = possibilities[i][0]
		opcodeName := getFunctionName(assignments[i])
		possibilities[i] = make([]opcode, 0)

		for j, available := range possibilities {
			for k := 0; k < len(available); k++ {
				if getFunctionName(available[k]) == opcodeName {
					possibilities[j] = append(available[:k], available[k+1:]...)
					break
				}
			}
		}

		i = -1
	}

	registers := []int{0, 0, 0, 0}

	for _, t := range strings.Split(testProgram, "\n") {
		if t == "" {
			continue
		}

		var instruction, a, b, c int

		_, err := fmt.Sscanf(t, "%d %d %d %d", &instruction, &a, &b, &c)

		if err != nil {
			panic(err)
		}

		assignments[instruction](registers, a, b, c)
	}

	return strconv.Itoa(registers[0]), nil
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func addi(registers []int, a int, b int, c int) {
	registers[c] = registers[a] + b
}

func addr(registers []int, a int, b int, c int) {
	registers[c] = registers[a] + registers[b]
}

func bani(registers []int, a int, b int, c int) {
	registers[c] = registers[a] & b
}

func banr(registers []int, a int, b int, c int) {
	registers[c] = registers[a] & registers[b]
}

func bori(registers []int, a int, b int, c int) {
	registers[c] = registers[a] | b
}

func borr(registers []int, a int, b int, c int) {
	registers[c] = registers[a] | registers[b]
}

func eqir(registers []int, a int, b int, c int) {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(registers []int, a int, b int, c int) {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(registers []int, a int, b int, c int) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtir(registers []int, a int, b int, c int) {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(registers []int, a int, b int, c int) {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(registers []int, a int, b int, c int) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func muli(registers []int, a int, b int, c int) {
	registers[c] = registers[a] * b
}

func mulr(registers []int, a int, b int, c int) {
	registers[c] = registers[a] * registers[b]
}

func seti(registers []int, a int, b int, c int) {
	registers[c] = a
}

func setr(registers []int, a int, b int, c int) {
	registers[c] = registers[a]
}
