package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: "12",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "2",
			},
		},
	},
	{
		Input: "14",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "2",
			},
			{
				Solver: SolvePart2,
				Output: "2",
			},
		},
	},
	{
		Input: "1969",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "654",
			},
			{
				Solver: SolvePart2,
				Output: "966",
			},
		},
	},
	{
		Input: "100756",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "33583",
			},
			{
				Solver: SolvePart2,
				Output: "50346",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "3337604",
			},
			{
				Solver: SolvePart2,
				Output: "5003530",
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
