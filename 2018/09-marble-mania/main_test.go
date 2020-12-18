package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: "9 players; last marble is worth 23 points",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "32",
			},
		},
	},
	{
		Input: "10 players; last marble is worth 1618 points",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "8317",
			},
		},
	},
	{
		Input: "13 players; last marble is worth 7999 points",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "146373",
			},
		},
	},
	{
		Input: "17 players; last marble is worth 1104 points",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "2764",
			},
		},
	},
	{
		Input: "21 players; last marble is worth 6111 points",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "54718",
			},
		},
	},
	{
		Input: "30 players; last marble is worth 5807 points",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "37305",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "390592",
			},
			{
				Solver: SolvePart2,
				Output: "3277920293",
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
