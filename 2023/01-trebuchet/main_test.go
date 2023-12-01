package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "142",
			},
		},
	},
	{
		Input: "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "281",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "55017",
			},
			{
				Solver: SolvePart2,
				Output: "53539",
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
