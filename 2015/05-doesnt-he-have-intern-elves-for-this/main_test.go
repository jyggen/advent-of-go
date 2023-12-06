package main

import (
	"strconv"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "ugknbfddgicrmopn",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1",
			},
		},
	},
	{
		Input: "jchzalrnumimnmhp",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0",
			},
		},
	},
	{
		Input: "haegwjzuvuyypxyu",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0",
			},
		},
	},
	{
		Input: "dvszwmarrgswjxmb",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0",
			},
		},
	},
	{
		Input: "qjhvhtzxzqqjkmpb",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "1",
			},
		},
	},
	{
		Input: "xxyxx",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "1",
			},
		},
	},
	{
		Input: "uurcxstgmygtbstg",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "0",
			},
		},
	},
	{
		Input: "ieodomkazucvgmuy",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "0",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "236",
			},
			{
				Solver: SolvePart2,
				Output: "51",
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
