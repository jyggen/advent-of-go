package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "165",
			},
		},
	},
	{
		Input: "mask = 000000000000000000000000000000X1001X\nmem[42] = 100\nmask = 00000000000000000000000000000000X0XX\nmem[26] = 1",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "208",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "9615006043476",
			},
			{
				Solver: SolvePart2,
				Output: "4275496544925",
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
