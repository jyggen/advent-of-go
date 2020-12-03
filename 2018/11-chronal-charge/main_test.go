package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: "18",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "33,45",
			},
			{
				Solver: SolvePart2,
				Output: "90,269,16",
			},
		},
	},
	{
		Input: "42",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "21,61",
			},
			{
				Solver: SolvePart2,
				Output: "232,251,12",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "21,41",
			},
			{
				Solver: SolvePart2,
				Output: "227,199,19",
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
