package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
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

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(solve(input, calculate)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(solve(input, calculate2)), nil
}

type calculator func(input string) int

func solve(input string, c calculator) int {
	sum := 0

	for _, line := range utils.ToStringSlice(input, "\n") {
		line = "(" + line + ")"
		stack := make([]int, 0)

		for i := 0; i < len(line); i++ {
			if line[i] == '(' {
				stack = append(stack, i)
			} else if line[i] == ')' {
				last := len(stack) - 1
				j := stack[last]
				stack = stack[:last]

				line = line[:j] + strconv.Itoa(c(line[j+1:i])) + line[i+1:]
				i = j - 1
			}
		}

		result, _ := strconv.Atoi(line)
		sum += result

	}

	return sum
}

func calculate(input string) int {
	operator := "+"
	sum := 0

	for _, f := range strings.Fields(input) {
		if f == "+" {
			operator = "+"
		} else if f == "*" {
			operator = "*"
		} else {
			n, _ := strconv.Atoi(f)

			switch operator {
			case "+":
				sum += n
			case "*":
				sum *= n
			}
		}
	}

	return sum
}

func calculate2(input string) int {
	fields := strings.Fields(input)

	for i := 0; i < len(fields); i++ {
		if fields[i] == "+" {
			a, _ := strconv.Atoi(fields[i-1])
			b, _ := strconv.Atoi(fields[i+1])
			fields = append(fields[:i-1], append([]string{strconv.Itoa(a + b)}, fields[i+2:]...)...)
			i = i - 2
		}
	}

	for i := 0; i < len(fields); i++ {
		if fields[i] == "*" {
			a, _ := strconv.Atoi(fields[i-1])
			b, _ := strconv.Atoi(fields[i+1])
			fields = append(fields[:i-1], append([]string{strconv.Itoa(a * b)}, fields[i+2:]...)...)
			i = i - 2
		}
	}

	result, _ := strconv.Atoi(fields[0])

	return result
}
