package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "Player 1 starting position: 4\nPlayer 2 starting position: 8",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "739785",
			},
			/*{
				Solver: SolvePart2,
				Output: "5",
			},*/
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "989352",
			},
			/*{
				Solver: SolvePart2,
				Output: "1567",
			},*/
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
