package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: ">",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "2",
			},
		},
	},
	{
		Input: "^>v<",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "4",
			},
			{
				Solver: SolvePart2,
				Output: "3",
			},
		},
	},
	{
		Input: "^v^v^v^v^v",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "2",
			},
			{
				Solver: SolvePart2,
				Output: "11",
			},
		},
	},
	{
		Input: "^v",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "3",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "2572",
			},
			{
				Solver: SolvePart2,
				Output: "2631",
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
