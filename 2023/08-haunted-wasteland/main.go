package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type parsed struct {
	leftRight *ring.Ring
	network   map[[3]uint8][2][3]uint8
}

func parse(input string) parsed {
	lines := utils.ToStringSlice(input, "\n")
	r := ring.New(len(lines[0]))
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = lines[0][i]
		r = r.Next()
	}

	network := make(map[[3]uint8][2][3]uint8, len(lines[2:]))

	for _, l := range lines[2:] {
		network[[3]uint8{l[0], l[1], l[2]}] = [2][3]uint8{
			{l[7], l[8], l[9]},
			{l[12], l[13], l[14]},
		}
	}

	return parsed{
		leftRight: r,
		network:   network,
	}
}

func solve(p parsed, current [3]uint8, isReached func([3]uint8) bool) int {
	steps := 1

	var val uint8

	for {
		val, _ = p.leftRight.Value.(uint8)

		if val == 'L' {
			current = p.network[current][0]
		} else {
			current = p.network[current][1]
		}

		if isReached(current) {
			return steps
		}

		p.leftRight = p.leftRight.Next()
		steps++
	}
}

func SolvePart1(input string) (string, error) {
	p := parse(input)

	return strconv.Itoa(solve(p, [3]uint8{'A', 'A', 'A'}, func(current [3]uint8) bool {
		return current == [3]uint8{'Z', 'Z', 'Z'}
	})), nil
}

func SolvePart2(input string) (string, error) {
	p := parse(input)
	steps := 1

	for k := range p.network {
		if k[2] == 'A' {
			steps = utils.LeastCommonMultiple(steps, solve(p, k, func(current [3]uint8) bool {
				return current[2] == 'Z'
			}))
		}
	}

	return strconv.Itoa(steps), nil
}
