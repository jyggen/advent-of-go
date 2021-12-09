package gameboy

import (
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/utils"
)

const (
	Acc = iota
	Jmp = iota
	Nop = iota
)

type Gameboy struct {
	accumulator int
	opcodes     []*opcode
	length      int
	current     int
	next        int
}

type opcode struct {
	kind   int
	value  int
	visits int
}

func New(input string) *Gameboy {
	lines := utils.ToStringSlice(input, "\n")
	opcodes := make([]*opcode, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		value, _ := strconv.Atoi(parts[1])
		kind := Nop

		switch parts[0] {
		case "acc":
			kind = Acc
		case "jmp":
			kind = Jmp
		}

		opcodes[i] = &opcode{
			kind:  kind,
			value: value,
		}
	}

	return &Gameboy{
		current: -1,
		next:    0,
		length:  len(opcodes),
		opcodes: opcodes,
	}
}

func (gb *Gameboy) Accumulator() int {
	return gb.accumulator
}

func (gb *Gameboy) Opcodes() []*opcode {
	return gb.opcodes
}

func (gb *Gameboy) Lookahead() *opcode {
	if gb.next >= gb.length {
		return nil
	}

	return gb.opcodes[gb.next]
}

func (gb *Gameboy) Reset() {
	gb.accumulator = 0
	gb.current = -1
	gb.next = 0

	for _, op := range gb.opcodes {
		op.visits = 0
	}
}

func (gb *Gameboy) Run() {
	for gb.next < gb.length {
		gb.Step()
	}
}

func (gb *Gameboy) Step() {
	gb.current = gb.next

	switch gb.opcodes[gb.current].kind {
	case Acc:
		gb.accumulator += gb.opcodes[gb.current].value
		gb.next++
	case Jmp:
		gb.next += gb.opcodes[gb.current].value
	case Nop:
		gb.next++
	}

	gb.opcodes[gb.current].visits++
}

func (op *opcode) Kind() int {
	return op.kind
}

func (op *opcode) SetKind(kind int) {
	op.kind = kind
}

func (op *opcode) Value() int {
	return op.value
}

func (op *opcode) Visits() int {
	return op.visits
}
