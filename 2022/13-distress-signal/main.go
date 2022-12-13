package main

import (
	"fmt"
	"os"
	"sort"
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

type parseCache map[string][]string

func parse(packet string, cache parseCache) []string {
	if v, ok := cache[packet]; ok {
		return v
	}

	output := make([]string, 0)
	depth := -1
	buffer := make([]rune, 0)

	for _, r := range packet {
		switch r {
		case '[':
			depth++

			if depth > 0 {
				buffer = append(buffer, r)
			}
		case ']':
			if depth > 0 {
				buffer = append(buffer, r)
			}

			depth--
		case ',':
			if depth > 0 {
				buffer = append(buffer, r)
			} else {
				output = append(output, string(buffer))
				buffer = make([]rune, 0)
			}
		default:
			buffer = append(buffer, r)
		}
	}

	if len(buffer) > 0 {
		output = append(output, string(buffer))
	}

	cache[packet] = output

	return output
}

func compare(left []string, right []string, cache parseCache) int {
	max := utils.MaxInt(len(left), len(right))

	for i := 0; i < max; i++ {
		if i == len(left) {
			return 1
		}

		if i == len(right) {
			return -1
		}

		leftIsList := false
		rightIsList := false

		if left[i][0] == '[' {
			leftIsList = true
		}

		if right[i][0] == '[' {
			rightIsList = true
		}

		if !leftIsList && !rightIsList {
			leftInt, _ := strconv.Atoi(left[i])
			rightInt, _ := strconv.Atoi(right[i])

			if leftInt < rightInt {
				return 1
			} else if leftInt > rightInt {
				return -1
			}

			continue
		} else if leftIsList && !rightIsList {
			result := compare(parse(left[i], cache), parse("["+right[i]+"]", cache), cache)

			if result != 0 {
				return result
			}
		} else if !leftIsList && rightIsList {
			result := compare(parse("["+left[i]+"]", cache), parse(right[i], cache), cache)

			if result != 0 {
				return result
			}
		} else if leftIsList && rightIsList {
			result := compare(parse(left[i], cache), parse(right[i], cache), cache)

			if result != 0 {
				return result
			}
		}
	}

	return 0
}

func SolvePart1(input string) (string, error) {
	pairs := utils.ToStringSlice(input, "\n\n")
	cache := make(parseCache, 0)
	sorted := 0

	for i, p := range pairs {
		values := utils.ToStringSlice(p, "\n")

		if compare(parse(values[0], cache), parse(values[1], cache), cache) != -1 {
			sorted += i + 1
		}
	}

	return strconv.Itoa(sorted), nil
}

func SolvePart2(input string) (string, error) {
	pairs := utils.ToStringSlice(input, "\n\n")
	cache := make(parseCache, 0)
	packets := [][]string{parse("[[2]]", cache), parse("[[6]]", cache)}
	result := 1

	for _, p := range pairs {
		values := utils.ToStringSlice(p, "\n")
		packets = append(packets, parse(values[0], cache), parse(values[1], cache))
	}

	sort.Slice(packets, func(i, j int) bool {
		result := compare(packets[i], packets[j], cache)

		if result == 1 {
			return true
		}

		return false
	})

	for i, p := range packets {
		if len(p) == 1 && (p[0] == "[2]" || p[0] == "[6]") {
			result *= i + 1
		}
	}

	return strconv.Itoa(result), nil
}
