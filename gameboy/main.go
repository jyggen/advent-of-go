package gameboy

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	Acc = iota
	Jmp = iota
	Nop = iota
)

var inputParser *regexp.Regexp

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

func init() {
	inputParser = regexp.MustCompile(`(?m)^([a-z]{3}) ([+-]\d+)$`)
}

func New(input string) *Gameboy {
	input = strings.TrimSpace(input)
	opcodes := make([]*opcode, strings.Count(input, "\n")+1)

	for i, m := range inputParser.FindAllStringSubmatch(input, -1) {
		value, _ := strconv.Atoi(m[2])
		kind := Nop

		switch m[1] {
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
	return gb.opcodes[gb.next]
}

func (gb *Gameboy) Reset() {
	gb.accumulator = 0
	gb.current = -1
	gb.next = 0
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
