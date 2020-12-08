package intcode

import (
	"fmt"
	"strconv"
)

type Computer struct {
	initial  []int
	length   int
	position int
	running  bool
	output   []int
	state    []int
}

type opcode func(input int, p1mode int, p2mode int, p3mode int) error

const (
	positionMode  = iota
	immediateMode = iota
)

func NewComputer(input []int) *Computer {
	inputLen := len(input)
	pc := &Computer{
		initial: make([]int, inputLen),
		length:  inputLen,
		output:  make([]int, 0),
		state:   make([]int, inputLen),
	}

	copy(pc.initial, input)

	pc.Reset()

	return pc
}

func (pc *Computer) Execute(input int) error {
	pc.position = 0
	pc.output = make([]int, 0)
	pc.running = true

	for pc.running {
		id, p1mode, p2mode, p3mode := pc.parse()

		var op opcode

		switch id {
		case 1:
			op = pc.opAdd
		case 2:
			op = pc.opMultiply
		case 3:
			op = pc.opInput
		case 4:
			op = pc.opOutput
		case 5:
			op = pc.opJumpIfTrue
		case 6:
			op = pc.opJumpIfFalse
		case 7:
			op = pc.opLessThan
		case 8:
			op = pc.opEquals
		case 99:
			op = pc.opHalt
		default:
			return fmt.Errorf("unknown opcode: %d", id)
		}

		if err := op(input, p1mode, p2mode, p3mode); err != nil {
			return err
		}
	}

	return nil
}

func (pc *Computer) Output() []int {
	return pc.output
}

func (pc *Computer) Reset() {
	copy(pc.state, pc.initial)
}

func (pc *Computer) SetValue(position int, value int) {
	pc.state[position] = value
}

func (pc *Computer) Value(position int) int {
	return pc.state[position]
}

func (pc *Computer) parse() (int, int, int, int) {
	opcode := pc.state[pc.position]
	strOpcode := strconv.Itoa(opcode)
	opcodeLen := len(strOpcode)

	if opcode > 99 {
		opcode, _ = strconv.Atoi(strOpcode[opcodeLen-2:])
	}

	var p1mode, p2mode, p3mode int

	if opcodeLen > 2 {
		p1mode, _ = strconv.Atoi(strOpcode[opcodeLen-3 : opcodeLen-2])

		if opcodeLen > 3 {
			p2mode, _ = strconv.Atoi(strOpcode[opcodeLen-4 : opcodeLen-3])

			if opcodeLen > 4 {
				p3mode, _ = strconv.Atoi(strOpcode[opcodeLen-5 : opcodeLen-4])
			}
		}
	}

	return opcode, p1mode, p2mode, p3mode
}

func (pc *Computer) opAdd(input int, p1mode int, p2mode int, p3mode int) error {
	p1 := pc.state[pc.position+1]
	p2 := pc.state[pc.position+2]
	p3 := pc.state[pc.position+3]

	if p1mode == positionMode {
		p1 = pc.state[p1]
	}

	if p2mode == positionMode {
		p2 = pc.state[p2]
	}

	pc.state[p3] = p1 + p2
	pc.position += 4

	return nil
}

func (pc *Computer) opEquals(input int, p1mode int, p2mode int, p3mode int) error {
	p1 := pc.state[pc.position+1]
	p2 := pc.state[pc.position+2]
	p3 := pc.state[pc.position+3]

	if p1mode == positionMode {
		p1 = pc.state[p1]
	}

	if p2mode == positionMode {
		p2 = pc.state[p2]
	}

	if p1 == p2 {
		pc.state[p3] = 1
	} else {
		pc.state[p3] = 0
	}

	pc.position += 4

	return nil
}

func (pc *Computer) opHalt(input int, p1mode int, p2mode int, p3mode int) error {
	pc.running = false

	return nil
}

func (pc *Computer) opInput(input int, p1mode int, p2mode int, p3mode int) error {
	p1 := pc.state[pc.position+1]

	pc.state[p1] = input
	pc.position += 2

	return nil
}

func (pc *Computer) opJumpIfFalse(input int, p1mode int, p2mode int, p3mode int) error {
	p1 := pc.state[pc.position+1]
	p2 := pc.state[pc.position+2]

	if p1mode == positionMode {
		p1 = pc.state[p1]
	}

	if p2mode == positionMode {
		p2 = pc.state[p2]
	}

	if p1 == 0 {
		pc.position = p2
	} else {
		pc.position += 3
	}

	return nil
}

func (pc *Computer) opJumpIfTrue(input int, p1mode int, p2mode int, p3mode int) error {
	p1 := pc.state[pc.position+1]
	p2 := pc.state[pc.position+2]

	if p1mode == positionMode {
		p1 = pc.state[p1]
	}

	if p2mode == positionMode {
		p2 = pc.state[p2]
	}

	if p1 != 0 {
		pc.position = p2
	} else {
		pc.position += 3
	}

	return nil
}

func (pc *Computer) opLessThan(input int, p1mode int, p2mode int, p3mode int) error {
	p1 := pc.state[pc.position+1]
	p2 := pc.state[pc.position+2]
	p3 := pc.state[pc.position+3]

	if p1mode == positionMode {
		p1 = pc.state[p1]
	}

	if p2mode == positionMode {
		p2 = pc.state[p2]
	}

	if p1 < p2 {
		pc.state[p3] = 1
	} else {
		pc.state[p3] = 0
	}

	pc.position += 4

	return nil
}

func (pc *Computer) opMultiply(input int, p1mode int, p2mode int, p3mode int) error {
	p1 := pc.state[pc.position+1]
	p2 := pc.state[pc.position+2]
	p3 := pc.state[pc.position+3]

	if p1mode == positionMode {
		p1 = pc.state[p1]
	}

	if p2mode == positionMode {
		p2 = pc.state[p2]
	}

	pc.state[p3] = p1 * p2
	pc.position += 4

	return nil
}

func (pc *Computer) opOutput(input int, p1mode int, p2mode int, p3mode int) error {
	p1 := pc.state[pc.position+1]

	if p1mode == positionMode {
		p1 = pc.state[p1]
	}

	pc.output = append(pc.output, p1)
	pc.position += 2

	return nil
}
