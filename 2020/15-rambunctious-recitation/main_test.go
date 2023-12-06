package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "0,3,6",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "436",
			},
			{
				Solver: SolvePart2,
				Output: "175594",
			},
		},
	},
	{
		Input: "1,3,2",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1",
			},
			{
				Solver: SolvePart2,
				Output: "2578",
			},
		},
	},
	{
		Input: "2,1,3",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "10",
			},
			{
				Solver: SolvePart2,
				Output: "3544142",
			},
		},
	},
	{
		Input: "1,2,3",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "27",
			},
			{
				Solver: SolvePart2,
				Output: "261214",
			},
		},
	},
	{
		Input: "2,3,1",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "78",
			},
			{
				Solver: SolvePart2,
				Output: "6895259",
			},
		},
	},
	{
		Input: "3,2,1",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "438",
			},
			{
				Solver: SolvePart2,
				Output: "18",
			},
		},
	},
	{
		Input: "3,1,2",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1836",
			},
			{
				Solver: SolvePart2,
				Output: "362",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "706",
			},
			{
				Solver: SolvePart2,
				Output: "19331",
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
