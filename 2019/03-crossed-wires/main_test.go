package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: "R8,U5,L5,D3\nU7,R6,D4,L4",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "6",
			},
			{
				Solver: SolvePart2,
				Output: "30",
			},
		},
	},
	{
		Input: "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "159",
			},
			{
				Solver: SolvePart2,
				Output: "610",
			},
		},
	},
	{
		Input: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "135",
			},
			{
				Solver: SolvePart2,
				Output: "410",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "1084",
			},
			{
				Solver: SolvePart2,
				Output: "9240",
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
