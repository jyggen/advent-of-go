package main

import (
	"bytes"
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func resolve(symbols []byte, ints []int, cache map[string]int) int {
	key := cacheKey([][]byte{symbols}, ints)

	if v, ok := cache[key]; ok {
		return v
	}

	combinations := utils.Combinations([]rune{'#', '?'}, ints[0])
	arrangements := 0

	for _, c := range combinations {
		mustInclude := bytes.Index(symbols, []byte{'#'})
		maxOffset := len(symbols) - len(c)

		if mustInclude != -1 {
			maxOffset = mustInclude
		}

		for offset := 0; offset <= maxOffset; offset++ {
			index := bytes.Index(symbols[offset:], []byte(string(c)))
			if index == -1 || (index+offset > maxOffset) {
				break
			}

			if (offset+index+len(c) < len(symbols) && symbols[offset+index+len(c)] == '#') || (offset+index != 0 && symbols[offset+index-1] == '#') {
				continue
			}

			if len(ints) == 1 {
				if bytes.IndexByte(symbols[index+offset+len(c):], '#') == -1 {
					arrangements++
				}
			} else if offset+index+len(c)+1 < len(symbols) {
				arrangements += resolve(symbols[offset+index+len(c)+1:], ints[1:], cache)
			}

			offset += index
		}
	}

	cache[key] = arrangements

	return arrangements
}

func cacheKey(clusters [][]byte, ints []int) string {
	var builder strings.Builder

	builder.Write(bytes.Join(clusters, []byte{0}))
	builder.WriteByte(0)

	for _, i := range ints {
		builder.Write([]byte{byte(i), 0})
	}

	return builder.String()
}

func SolvePart1(input string) (string, error) {
	lines := utils.ToByteSlice(input, '\n')
	arrangements := 0
	cache := make(map[string]int)

	for _, l := range lines {
		separatorIndex := bytes.IndexByte(l, ' ')
		symbols := l[:separatorIndex]
		ints, _ := utils.ToIntegerSlice(string(l[separatorIndex+1:]), ",")
		yo := resolve(symbols, ints, cache)
		arrangements += yo
	}

	return strconv.Itoa(arrangements), nil
}

func SolvePart2(input string) (string, error) {
	lines := utils.ToByteSlice(input, '\n')
	arrangements := 0
	cache := make(map[string]int)

	for _, l := range lines {
		separatorIndex := bytes.IndexByte(l, ' ')
		symbols := l[:separatorIndex]
		symbols = bytes.Repeat(append(symbols, '?'), 5)
		symbols = symbols[:len(symbols)-1]
		ints, _ := utils.ToIntegerSlice(string(l[separatorIndex+1:]), ",")
		newInts := make([]int, 0, len(ints)*5)
		for i := 0; i < 5; i++ {
			newInts = append(newInts, ints...)
		}
		yo := resolve(symbols, newInts, cache)

		arrangements += yo
	}

	return strconv.Itoa(arrangements), nil
}
