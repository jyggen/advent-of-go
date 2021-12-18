package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"math"
	"os"
	"strconv"
)

type snailNumber struct {
	left   interface{}
	parent *snailNumber
	right  interface{}
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
	rows := utils.ToRuneSlice(input, "\n")
	numbers := make([]*snailNumber, len(rows))

	for i, r := range rows {
		numbers[i], _ = parseInput(r, 1)
	}

	for len(numbers) > 1 {
		numbers = append([]*snailNumber{add(numbers[0], numbers[1])}, numbers[2:]...)
		numbers[0] = reduce(numbers[0])
	}

	return strconv.Itoa(magnitude(numbers[0])), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToRuneSlice(input, "\n")
	highest := 0

	for i := range rows {
		for j := range rows {
			if i == j {
				continue
			}

			n1, _ := parseInput(rows[i], 1)
			n2, _ := parseInput(rows[j], 1)
			result := magnitude(reduce(add(n1, n2)))

			if result > highest {
				highest = result
			}
		}
	}

	return strconv.Itoa(highest), nil
}

func add(a *snailNumber, b *snailNumber) *snailNumber {
	n := &snailNumber{left: a, right: b}
	a.parent = n
	b.parent = n

	return n
}

func parseInput(input []rune, offset int) (*snailNumber, int) {
	number := &snailNumber{}
	numOffset := 0
	numbers := [2]interface{}{}

	var value *snailNumber

Loop:
	for ; offset < len(input); offset++ {
		switch input[offset] {
		case '[':
			value, offset = parseInput(input, offset+1)
			value.parent = number
			numbers[numOffset] = value
		case ',':
			numOffset++
		case ']':
			break Loop
		default:
			numbers[numOffset], _ = strconv.Atoi(string(input[offset]))
		}
	}

	number.left = numbers[0]
	number.right = numbers[1]

	return number, offset
}

func magnitude(number *snailNumber) int {
	left, right := 0, 0

	switch v := number.left.(type) {
	case int:
		left = v * 3
	case *snailNumber:
		left = magnitude(v) * 3
	}

	switch v := number.right.(type) {
	case int:
		right = v * 2
	case *snailNumber:
		right = magnitude(v) * 2
	}

	return left + right
}

func reduce(number *snailNumber) *snailNumber {
	target := findExplodeTarget(number, 0)

	if target != nil {
		explode(target)

		return reduce(number)
	}

	target = findSplitTarget(number)

	if target != nil {
		split(target)

		return reduce(number)
	}

	return number
}

func split(number *snailNumber) {
	v, ok := number.left.(int)

	if ok && v > 9 {
		number.left = &snailNumber{
			left:   int(math.Floor(float64(v) / 2)),
			right:  int(math.Ceil(float64(v) / 2)),
			parent: number,
		}

		return
	}

	v, _ = number.right.(int)
	number.right = &snailNumber{
		left:   int(math.Floor(float64(v) / 2)),
		right:  int(math.Ceil(float64(v) / 2)),
		parent: number,
	}
}

func explode(number *snailNumber) {
	prevParent := number
	parent := number.parent

	for parent != nil && parent.left == prevParent {
		prevParent = parent
		parent = parent.parent
	}

	if parent != nil {
		switch v := parent.left.(type) {
		case int:
			parent.left = v + number.left.(int)
		case *snailNumber:
			parent = v
		LeftChildLoop:
			for {
				switch v2 := parent.right.(type) {
				case int:
					parent.right = v2 + number.left.(int)
					break LeftChildLoop
				case *snailNumber:
					parent = v2
				}
			}
		}
	}

	prevParent = number
	parent = number.parent

	for parent != nil && parent.right == prevParent {
		prevParent = parent
		parent = parent.parent
	}

	if parent != nil {
		switch v := parent.right.(type) {
		case int:
			parent.right = v + number.right.(int)
		case *snailNumber:
			parent = v
		RightChildLoop:
			for {
				switch v2 := parent.left.(type) {
				case int:
					parent.left = v2 + number.right.(int)
					break RightChildLoop
				case *snailNumber:
					parent = v2
				}
			}
		}
	}

	if number.parent.left == number {
		number.parent.left = 0
	} else {
		number.parent.right = 0
	}
}

func findExplodeTarget(number *snailNumber, depth int) *snailNumber {
	var target *snailNumber

	if depth == 4 {
		return number
	}

	switch v := number.left.(type) {
	case *snailNumber:
		target = findExplodeTarget(v, depth+1)
	}

	if target != nil {
		return target
	}

	switch v := number.right.(type) {
	case *snailNumber:
		target = findExplodeTarget(v, depth+1)
	}

	return target
}

func findSplitTarget(number *snailNumber) *snailNumber {
	var target *snailNumber

	switch v := number.left.(type) {
	case int:
		if v > 9 {
			return number
		}
	case *snailNumber:
		target = findSplitTarget(v)
	}

	if target != nil {
		return target
	}

	switch v := number.right.(type) {
	case int:
		if v > 9 {
			return number
		}
	case *snailNumber:
		target = findSplitTarget(v)
	}

	return target
}
