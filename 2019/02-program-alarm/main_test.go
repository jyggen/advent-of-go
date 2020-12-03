package main

import (
"fmt"
"github.com/jyggen/advent-of-go/solver"
"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: solver.InputFromFile("input.txt"),
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "9581917",
			},
			{
				Solver: SolvePart2,
				Output: "2505",
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
		t.Run(fmt.Sprint(i), func (subtest *testing.T) {
			testCase.Test(subtest)
		})
	}
}
