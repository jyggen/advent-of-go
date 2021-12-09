package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	"testing"
)

var testCases = []*solver2.TestCase{
	{
		Input: "abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "12",
			},
		},
	},
	{
		Input: "abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz",
		Solvers: []*solver2.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "fgij",
			},
		},
	},
	{
		Input: solver2.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver2.TestCaseSolver{
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
