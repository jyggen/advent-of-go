package main

import (
	"strconv"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "12",
			},
		},
	},
	{
		Input: "abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "fgij",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "7350",
			},
			{
				Solver: SolvePart2,
				Output: "wmlnjevbfodamyiqpucrhsukg",
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
