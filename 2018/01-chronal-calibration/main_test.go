package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "+1\n-2\n+3\n+1",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "3",
			},
			{
				Solver: SolvePart2,
				Output: "2",
			},
		},
	},
	{
		Input: "+1\n+1\n+1",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "3",
			},
		},
	},
	{
		Input: "+1\n+1\n-2",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0",
			},
		},
	},
	{
		Input: "-1\n-2\n-3",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "-6",
			},
		},
	},
	{
		Input: "+1\n-1",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "0",
			},
		},
	},
	{
		Input: "+3\n+3\n+4\n-2\n-4",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "10",
			},
		},
	},
	{
		Input: "-6\n+3\n+8\n+5\n-6",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "5",
			},
		},
	},
	{
		Input: "+7\n+7\n-2\n-7\n-4",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "14",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "505",
			},
			{
				Solver: SolvePart2,
				Output: "72330",
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
