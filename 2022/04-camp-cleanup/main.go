package main

import (
	"bytes"
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func toByteArray(from int, to int) []byte {
	output := make([]byte, 0, to-from+1)

	for i := from; i <= to; i++ {
		output = append(output, byte(i))
	}

	return output
}

func SolvePart1(input string) (string, error) {
	integers := utils.ToOptimisticIntSlice(input)
	overlaps := 0

	for i := 0; i < len(integers); i += 4 {
		first, second := toByteArray(integers[i], integers[i+1]), toByteArray(integers[i+2], integers[i+3])

		if bytes.Contains(first, second) || bytes.Contains(second, first) {
			overlaps++
		}
	}

	return strconv.Itoa(overlaps), nil
}

func SolvePart2(input string) (string, error) {
	integers := utils.ToOptimisticIntSlice(input)
	overlaps := 0

	for i := 0; i < len(integers); i += 4 {
		first, second := toByteArray(integers[i], integers[i+1]), toByteArray(integers[i+2], integers[i+3])

		if bytes.ContainsAny(first, string(second)) || bytes.ContainsAny(second, string(first)) {
			overlaps++
		}
	}

	return strconv.Itoa(overlaps), nil
}
