package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	"testing"
)

var testCases = []*solver2.TestCase{
	{
		Input: "9",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "5158916779",
			},
		},
	},
	{
		Input: "5",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "0124515891",
			},
		},
	},
	{
		Input: "18",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "9251071085",
			},
		},
	},
	{
		Input: "2018",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "5941429882",
			},
		},
	},
	{
		Input: "51589",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "9",
			},
		},
	},
	{
		Input: "01245",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "5",
			},
		},
	},
	{
		Input: "92510",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "18",
			},
		},
	},
	{
		Input: "59414",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "2018",
			},
		},
	},
	{
		Input: solver2.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1464411010",
			},
			{
				Solver: SolvePart2,
				Output: "20288091",
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
