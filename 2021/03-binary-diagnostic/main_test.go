package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "198",
			},
			{
				Solver: SolvePart2,
				Output: "230",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1025636",
			},
			{
				Solver: SolvePart2,
				Output: "793873",
			},
		},
	},
}

func BenchmarkSolvers(b *testing.B) {
	for i, testCase := range testCases {
		b.Run(fmt.Sprint(i), func(subtest *testing.B) {
			testCase.Benchmark(subtest)
		})
	}
}

func TestSolvers(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprint(i), func(subtest *testing.T) {
			testCase.Test(subtest)
		})
	}
}
