package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "3068",
			},
			{
				Solver: SolvePart2,
				Output: "1514285714288",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "3206",
			},
			{
				Solver: SolvePart2,
				Output: "1602881844347",
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
