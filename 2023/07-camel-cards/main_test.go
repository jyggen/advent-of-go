package main

import (
	"strconv"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "6440",
			},
			{
				Solver: SolvePart2,
				Output: "5905",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "248812215",
			},
			{
				Solver: SolvePart2,
				Output: "250057090",
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
