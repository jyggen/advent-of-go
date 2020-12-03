package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "7",
			},
			{
				Solver: SolvePart2,
				Output: "336",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "198",
			},
			{
				Solver: SolvePart2,
				Output: "5140884672",
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
