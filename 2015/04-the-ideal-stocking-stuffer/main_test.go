package main

import (
	"strconv"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "abcdef",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "609043",
			},
		},
	},
	{
		Input: "pqrstuv",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1048970",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "346386",
			},
			{
				Solver: SolvePart2,
				Output: "9958218",
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
