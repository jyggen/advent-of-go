package main

import (
	"strconv"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1320",
			},
			{
				Solver: SolvePart2,
				Output: "145",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "504449",
			},
			{
				Solver: SolvePart2,
				Output: "262044",
			},
		},
	},
}

func BenchmarkSolvers(b *testing.B) {
	for i, testCase := range testCases {
		b.Run(strconv.Itoa(i), func(subtest *testing.B) {
			testCase.Benchmark(subtest)
		})
	}
}

func TestSolvers(t *testing.T) {
	t.Parallel()

	for i, testCase := range testCases {
		testCase := testCase

		t.Run(strconv.Itoa(i), func(subtest *testing.T) {
			subtest.Parallel()
			testCase.Test(subtest)
		})
	}
}
