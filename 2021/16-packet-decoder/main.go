package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
)

const (
	sum int64 = iota
	product
	minimum
	maximum
	literal
	greaterThan
	lessThan
	equalTo
)

type packet struct {
	version int64
	kind    int64
	value   interface{}
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func hexToBinary(hex string) string {
	packet := big.NewInt(0)
	packet.SetString(hex, 16)
	binary := fmt.Sprintf("%b", packet)

	for _, r := range hex {
		if r != '0' {
			break
		}

		binary = "0000" + binary
	}

	for len(binary)%4 != 0 {
		binary = "0" + binary
	}

	return binary
}

func parseLiteralValue(binary string) (int64, string) {
	bits := ""
	offset := 0

	for {
		keepReading := binary[offset]
		bits += binary[offset+1 : offset+5]
		offset += 5

		if keepReading == '0' {
			break
		}

	}

	value, _ := strconv.ParseInt(bits, 2, 64)

	return value, binary[offset:]
}

func parsePacket(binary string) (packet, string) {
	version, _ := strconv.ParseInt(binary[0:3], 2, 64)
	typeId, _ := strconv.ParseInt(binary[3:6], 2, 64)
	p := packet{
		version: version,
		kind:    typeId,
	}

	switch p.kind {
	case literal:
		p.value, binary = parseLiteralValue(binary[6:])
	default:
		p.value, binary = parseOperator(binary[6:])
	}

	return p, binary
}

func parseOperator(binary string) ([]packet, string) {
	var p packet

	lengthType := binary[0]
	subPackets := make([]packet, 0)

	switch lengthType {
	case '0':
		totalLength, _ := strconv.ParseInt(binary[1:16], 2, 64)
		binary = binary[16:]
		start := len(binary)

		for start-len(binary) < int(totalLength) {
			p, binary = parsePacket(binary)
			subPackets = append(subPackets, p)
		}

	case '1':
		totalPackets, _ := strconv.ParseInt(binary[1:12], 2, 64)
		binary = binary[12:]

		for i := int64(0); i < totalPackets; i++ {
			p, binary = parsePacket(binary)
			subPackets = append(subPackets, p)
		}
	}

	return subPackets, binary
}

func resolvePacket(p packet) int64 {
	if p.kind == literal {
		return p.value.(int64)
	}

	value := int64(0)
	subPackets := p.value.([]packet)

	switch p.kind {
	case sum:
		for _, sp := range subPackets {
			value += resolvePacket(sp)
		}
	case product:
		value = 1

		for _, sp := range subPackets {
			value *= resolvePacket(sp)
		}
	case minimum:
		value = math.MaxInt64

		for _, sp := range subPackets {
			v := resolvePacket(sp)

			if v < value {
				value = v
			}
		}
	case maximum:
		for _, sp := range subPackets {
			v := resolvePacket(sp)

			if v > value {
				value = v
			}
		}
	case greaterThan:
		if resolvePacket(subPackets[0]) > resolvePacket(subPackets[1]) {
			value = 1
		}
	case lessThan:
		if resolvePacket(subPackets[0]) < resolvePacket(subPackets[1]) {
			value = 1
		}
	case equalTo:
		if resolvePacket(subPackets[0]) == resolvePacket(subPackets[1]) {
			value = 1
		}
	}

	return value
}

func sumVersions(p packet) int64 {
	sum := int64(0)
	sum += p.version

	switch p.kind {
	case literal:
		break
	default:
		subPackets := p.value.([]packet)

		for _, sp := range subPackets {
			sum += sumVersions(sp)
		}
	}

	return sum
}

func SolvePart1(input string) (string, error) {
	binary := hexToBinary(input)
	p, _ := parsePacket(binary)

	return strconv.Itoa(int(sumVersions(p))), nil
}

func SolvePart2(input string) (string, error) {
	binary := hexToBinary(input)
	p, _ := parsePacket(binary)

	return strconv.Itoa(int(resolvePacket(p))), nil
}
