package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "1 + 2 * 3 + 4 * 5 + 6",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "71",
			},
			{
				Solver: SolvePart2,
				Output: "231",
			},
		},
	},
	{
		Input: "1 + (2 * 3) + (4 * (5 + 6))",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "51",
			},
			{
				Solver: SolvePart2,
				Output: "51",
			},
		},
	},
	{
		Input: "2 * 3 + (4 * 5)",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "26",
			},
			{
				Solver: SolvePart2,
				Output: "46",
			},
		},
	},
	{
		Input: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "437",
			},
			{
				Solver: SolvePart2,
				Output: "1445",
			},
		},
	},
	{
		Input: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "12240",
			},
			{
				Solver: SolvePart2,
				Output: "669060",
			},
		},
	},
	{
		Input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "13632",
			},
			{
				Solver: SolvePart2,
				Output: "23340",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "5783053349377",
			},
			{
				Solver: SolvePart2,
				Output: "74821486966872",
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
