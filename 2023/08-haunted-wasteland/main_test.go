package main

import (
	"strconv"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "2",
			},
		},
	},
	{
		Input: "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "6",
			},
		},
	},
	{
		Input: "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "6",
			},
		},
	},
	{
		Short:   true,
		Solvers: []*solver.TestCaseSolver{},
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
