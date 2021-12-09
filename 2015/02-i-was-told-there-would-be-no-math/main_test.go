package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "2x3x4",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "58",
			},
			{
				Solver: SolvePart2,
				Output: "34",
			},
		},
	},
	{
		Input: "1x1x10",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "43",
			},
			{
				Solver: SolvePart2,
				Output: "14",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1586300",
			},
			{
				Solver: SolvePart2,
				Output: "3737498",
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
