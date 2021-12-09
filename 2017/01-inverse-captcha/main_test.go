package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	"testing"
)

var testCases = []*solver2.TestCase{
	{
		Input: "1122",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "3",
			},
		},
	},
	{
		Input: "1111",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "4",
			},
		},
	},
	{
		Input: "1234",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0",
			},
		},
	},
	{
		Input: "91212129",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "9",
			},
		},
	},
	{
		Input: "1212",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "6",
			},
		},
	},
	{
		Input: "1221",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "0",
			},
		},
	},
	{
		Input: "123425",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "4",
			},
		},
	},
	{
		Input: "123123",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "12",
			},
		},
	},
	{
		Input: "12131415",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "4",
			},
		},
	},
	{
		Input: solver2.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1253",
			},
			{
				Solver: SolvePart2,
				Output: "1278",
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
