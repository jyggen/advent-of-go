package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "939\n7,13,x,x,59,x,31,19",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "295",
			},
			{
				Solver: SolvePart2,
				Output: "1068781",
			},
		},
	},
	{
		Input: "939\n17,x,13,19",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "3417",
			},
		},
	},
	{
		Input: "939\n67,7,59,61",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "754018",
			},
		},
	},
	{
		Input: "939\n67,x,7,59,61",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "779210",
			},
		},
	},
	{
		Input: "939\n67,7,x,59,61",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "1261476",
			},
		},
	},
	{
		Input: "939\n1789,37,47,1889",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "1202161486",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "174",
			},
			{
				Solver: SolvePart2,
				Output: "780601154795940",
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
