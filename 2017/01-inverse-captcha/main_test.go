package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: "1122",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "3",
			},
		},
	},
	{
		Input: "1111",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "4",
			},
		},
	},
	{
		Input: "1234",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0",
			},
		},
	},
	{
		Input: "91212129",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "9",
			},
		},
	},
	{
		Input: "1212",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "6",
			},
		},
	},
	{
		Input: "1221",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "0",
			},
		},
	},
	{
		Input: "123425",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "4",
			},
		},
	},
	{
		Input: "123123",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "12",
			},
		},
	},
	{
		Input: "12131415",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "4",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver.TestCaseSolver{
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
