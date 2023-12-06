package main

import (
	"strconv"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "aA",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0",
			},
		},
	},
	{
		Input: "abBA",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0",
			},
		},
	},
	{
		Input: "abAB",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "4",
			},
		},
	},
	{
		Input: "aabAAB",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "6",
			},
		},
	},
	{
		Input: "dabAcCaCBAcCcaDA",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "10",
			},
			{
				Solver: SolvePart2,
				Output: "4",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "10368",
			},
			{
				Solver: SolvePart2,
				Output: "4122",
			},
		},
	},
}

func BenchmarkSolvers(b *testing.B) {
	for i, testCase := range testCases {
		b.Run(strconv.Itoa(i), func(subtest *testing.B) {
			testCase.Benchmark(subtest)
		})
	}
}

func TestSolvers(t *testing.T) {
	t.Parallel()

	for i, testCase := range testCases {
		testCase := testCase

		t.Run(strconv.Itoa(i), func(subtest *testing.T) {
			subtest.Parallel()
			testCase.Test(subtest)
		})
	}
}
