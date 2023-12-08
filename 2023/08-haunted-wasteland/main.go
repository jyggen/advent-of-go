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
	network   [][2]int
}

func hashCode(a uint8, b uint8, c uint8) int {
	return (int(c)*255+int(b))*255 + int(a)
}

func parse(input string) parsed {
	lines := utils.ToStringSlice(input, "\n")
	r := ring.New(len(lines[0]))
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = lines[0][i]
		r = r.Next()
	}

	network := make([][2]int, hashCode(255, 255, 255)+1)

	for _, l := range lines[2:] {
		network[hashCode(l[0], l[1], l[2])] = [2]int{
			hashCode(l[7], l[8], l[9]),
			hashCode(l[12], l[13], l[14]),
		}
	}

	return parsed{
		leftRight: r,
		network:   network,
	}
}

func solve(p parsed, current, endMin, endMax int) int {
	steps := 1

	var val uint8

	for {
		val, _ = p.leftRight.Value.(uint8)

		if val == 'L' {
			current = p.network[current][0]
		} else {
			current = p.network[current][1]
		}

		if current >= endMin && current <= endMax {
			return steps
		}

		p.leftRight = p.leftRight.Next()
		steps++
	}
}

func SolvePart1(input string) (string, error) {
	p := parse(input)

	return strconv.Itoa(solve(p, hashCode('A', 'A', 'A'), hashCode('Z', 'Z', 'Z'), hashCode('Z', 'Z', 'Z'))), nil
}

func SolvePart2(input string) (string, error) {
	p := parse(input)
	endMin := hashCode(0, 0, 'Z')
	endMax := hashCode(255, 255, 'Z')
	startMin := hashCode(0, 0, 'A')
	startMax := hashCode(255, 255, 'A')
	steps := 1

	for k, v := range p.network {
		if k >= startMin && k <= startMax && v[0] != 0 {
			steps = utils.LeastCommonMultiple(steps, solve(p, k, endMin, endMax))
		}
	}

	return strconv.Itoa(steps), nil
}
