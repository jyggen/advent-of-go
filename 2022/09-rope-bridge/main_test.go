package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "13",
			},
			{
				Solver: SolvePart2,
				Output: "1",
			},
		},
	},
	{
		Input: "R 5\nU 8\nL 8\nD 3\nR 17\nD 10\nL 25\nU 20",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "36",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "6181",
			},
			{
				Solver: SolvePart2,
				Output: "2386",
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
