package intcode

import (
	"fmt"
	"strconv"
)

type Computer struct {
	initial  []int
	length   int
	position int
	relative int
	running  bool
	state    []int
}

const (
	positionMode  = iota
	immediateMode = iota
	relativeMode  = iota
)

func NewComputer(input []int) *Computer {
	inputLen := len(input)
	pc := &Computer{
		initial: make([]int, inputLen),
		length:  inputLen,
		state:   make([]int, inputLen*10),
	}

	copy(pc.initial, input)

	pc.Reset()

	return pc
}

func (pc *Computer) Execute(input chan int, output chan int) {
	pc.position = 0
	pc.running = true

	for pc.running {
		id, p1mode, p2mode, p3mode := pc.parse()

		switch id {
		case 1:
			pc.opAdd(p1mode, p2mode, p3mode)
		case 2:
			pc.opMultiply(p1mode, p2mode, p3mode)
		case 3:
			pc.opInput(input, p1mode)
		case 4:
			pc.opOutput(output, p1mode)
		case 5:
			pc.opJumpIfTrue(p1mode, p2mode)
		case 6:
			pc.opJumpIfFalse(p1mode, p2mode)
		case 7:
			pc.opLessThan(p1mode, p2mode, p3mode)
		case 8:
			pc.opEquals(p1mode, p2mode, p3mode)
		case 9:
			pc.opRelativeBase(p1mode)
		case 99:
			pc.opHalt(output)
		default:
			panic(fmt.Errorf("unknown opcode: %d", id))
		}
	}
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

func (pc *Computer) opAdd(p1mode int, p2mode int, p3mode int) {
	pc.set(pc.get(pc.position+1, p1mode)+pc.get(pc.position+2, p2mode), pc.position+3, p3mode)
	pc.position += 4
}

func (pc *Computer) opEquals(p1mode int, p2mode int, p3mode int) {
	p1 := pc.get(pc.position+1, p1mode)
	p2 := pc.get(pc.position+2, p2mode)

	if p1 == p2 {
		pc.set(1, pc.position+3, p3mode)
	} else {
		pc.set(0, pc.position+3, p3mode)
	}

	pc.position += 4
}

func (pc *Computer) opHalt(output chan int) {
	close(output)
	pc.running = false
}

func (pc *Computer) opInput(input chan int, p1mode int) {
	pc.set(<-input, pc.position+1, p1mode)
	pc.position += 2
}

func (pc *Computer) opJump(p1mode int, p2mode int, compare func(a int, b int) bool) {
	p1 := pc.get(pc.position+1, p1mode)
	p2 := pc.get(pc.position+2, p2mode)

	if compare(p1, 0) {
		pc.position = p2
	} else {
		pc.position += 3
	}
}

func (pc *Computer) opJumpIfFalse(p1mode int, p2mode int) {
	pc.opJump(p1mode, p2mode, func(a int, b int) bool {
		return a == b
	})
}

func (pc *Computer) opJumpIfTrue(p1mode int, p2mode int) {
	pc.opJump(p1mode, p2mode, func(a int, b int) bool {
		return a != b
	})
}

func (pc *Computer) opLessThan(p1mode int, p2mode int, p3mode int) {
	p1 := pc.get(pc.position+1, p1mode)
	p2 := pc.get(pc.position+2, p2mode)

	if p1 < p2 {
		pc.set(1, pc.position+3, p3mode)
	} else {
		pc.set(0, pc.position+3, p3mode)
	}

	pc.position += 4
}

func (pc *Computer) opMultiply(p1mode int, p2mode int, p3mode int) {
	pc.set(pc.get(pc.position+1, p1mode)*pc.get(pc.position+2, p2mode), pc.position+3, p3mode)
	pc.position += 4
}

func (pc *Computer) opOutput(output chan int, p1mode int) {
	output <- pc.get(pc.position+1, p1mode)

	pc.position += 2
}

func (pc *Computer) opRelativeBase(p1mode int) {
	pc.relative += pc.get(pc.position+1, p1mode)
	pc.position += 2
}

func (pc *Computer) get(offset int, mode int) int {
	value := pc.state[offset]

	switch mode {
	case positionMode:
		return pc.state[value]
	case immediateMode:
		return value
	case relativeMode:
		return pc.state[value+pc.relative]
	}

	return 0
}

func (pc *Computer) set(value int, offset int, mode int) {
	offset = pc.state[offset]

	switch mode {
	case immediateMode:
	case positionMode:
		pc.state[offset] = value
	case relativeMode:
		pc.state[offset+pc.relative] = value
	}
}
