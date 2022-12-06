package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "7",
			},
			{
				Solver: SolvePart2,
				Output: "19",
			},
		},
	},
	{
		Input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "5",
			},
			{
				Solver: SolvePart2,
				Output: "23",
			},
		},
	},
	{
		Input: "nppdvjthqldpwncqszvftbrmjlhg",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "6",
			},
			{
				Solver: SolvePart2,
				Output: "23",
			},
		},
	},
	{
		Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "10",
			},
			{
				Solver: SolvePart2,
				Output: "29",
			},
		},
	},
	{
		Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "11",
			},
			{
				Solver: SolvePart2,
				Output: "26",
			},
		},
	},
	/*{
		Input: solver.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "TQRFCBSJJ",
			},
			{
				Solver: SolvePart2,
				Output: "RMHFJNVFP",
			},
		},
	},*/
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
