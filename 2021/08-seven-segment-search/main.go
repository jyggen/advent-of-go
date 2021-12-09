package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"math/bits"
	"os"
	"strconv"
)

const (
	a uint8 = 1 << iota
	b
	c
	d
	e
	f
	g
)

const (
	zero  = a | b | c | e | f | g
	one   = c | f
	two   = a | c | d | e | g
	three = a | c | d | f | g
	four  = b | c | d | f
	five  = a | b | d | f | g
	six   = a | b | d | e | f | g
	seven = a | c | f
	eight = a | b | c | d | e | f | g
	nine  = a | b | c | d | f | g
)

var numLookup = map[uint8]int{
	zero:  0,
	one:   1,
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
}

func toByteSlice(input string, separator string) []uint8 {
	stringSlice := utils2.ToStringSlice(input, separator)
	byteSlice := make([]uint8, 0, len(stringSlice))

	for _, runes := range stringSlice {
		val := uint8(0)

		for _, r := range runes {
			switch r {
			case 'a':
				val |= a
			case 'b':
				val |= b
			case 'c':
				val |= c
			case 'd':
				val |= d
			case 'e':
				val |= e
			case 'f':
				val |= f
			case 'g':
				val |= g
			}
		}

		byteSlice = append(byteSlice, val)
	}

	return byteSlice
}

type lookup struct {
	segments map[uint8]uint8
}

func (l *lookup) lockSegment(segment uint8) {
	for k := range l.segments {
		if k != segment {
			l.clearSegment(k, l.segments[segment])
		}
	}
}

func (l *lookup) clearNumber(number uint8, value uint8) {
	for k := range l.segments {
		if number&k != 0 {
			l.clearSegment(k, value)
		}
	}
}

func (l *lookup) clearSegment(segment uint8, value uint8) {
	l.segments[segment] &^= value
}

func (l *lookup) andNumber(number uint8, value uint8) {
	for k := range l.segments {
		if number&k != 0 {
			l.segments[k] &= value
		}
	}
}

func (l *lookup) clearSegmentByNumber(segment uint8, number uint8) {
	for k, v := range l.segments {
		if number&k != 0 {
			l.clearSegment(segment, v)
		}
	}
}

func byBitCount(bytes []uint8) map[int][]uint8 {
	countMap := make(map[int][]uint8, 0)

	for _, v := range bytes {
		count := bits.OnesCount8(v)

		if _, ok := countMap[count]; !ok {
			countMap[count] = make([]uint8, 0)
		}

		countMap[count] = append(countMap[count], v)
	}

	return countMap
}

func (l *lookup) hasAllBitsSet(number uint8, bits uint8) bool {
	for k, v := range l.segments {
		if bits&k != 0 && v&^number != 0 {
			return false
		}
	}

	return true
}

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	stringSlice := utils2.ToStringSlice(input, "\n")
	sum := 0

	for _, s := range stringSlice {
		output := utils2.ToStringSlice(s, "|")[1]
		bytes := toByteSlice(output, " ")

		for _, v := range bytes {
			switch bits.OnesCount8(v) {
			case 2, 3, 4, 7:
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	stringSlice := utils2.ToStringSlice(input, "\n")
	sum := 0

	for _, s := range stringSlice {
		parts := utils2.ToStringSlice(s, "|")
		patterns, output := toByteSlice(parts[0], " "), toByteSlice(parts[1], " ")
		countMap := byBitCount(append(patterns, output...))
		segments := &lookup{
			segments: map[uint8]uint8{
				a: eight,
				b: eight,
				c: eight,
				d: eight,
				e: eight,
				f: eight,
				g: eight,
			},
		}

		if v, ok := countMap[2]; ok {
			segments.andNumber(one, v[0])
			segments.clearSegmentByNumber(a, one)
		}

		if v, ok := countMap[4]; ok {
			segments.andNumber(four, v[0])
			segments.clearSegmentByNumber(a, four)
		}

		if v, ok := countMap[3]; ok {
			segments.andNumber(seven, v[0])
			segments.lockSegment(a)
		}

		if v, ok := countMap[6]; ok {
			for _, n := range v {
				if segments.hasAllBitsSet(n, four) {
					segments.andNumber(nine, n)
				} else if segments.hasAllBitsSet(n, seven) {
					segments.andNumber(zero, n)
				} else {
					segments.andNumber(six, n)
				}
			}

			segments.lockSegment(f)
			segments.lockSegment(c)
			segments.lockSegment(b)
			segments.lockSegment(g)
		}

		for i, bit := range output {
			bits := uint8(0)
			for k, v := range segments.segments {
				if bit&v != 0 {
					bits |= k
				}
			}

			value := numLookup[bits]

			for j := len(output) - 1; j > i; j-- {
				value *= 10
			}

			sum += value
		}
	}

	return strconv.Itoa(sum), nil
}
