package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"testing"
)

var testCases = []*solver.TestCase{
	{
		Input: "#######\n#.G...#\n#...EG#\n#.#.#G#\n#..G#E#\n#.....#\n#######",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "27730",
			},
			{
				Solver: SolvePart2,
				Output: "4988",
			},
		},
	},
	{
		Input: "#######\n#G..#E#\n#E#E.E#\n#G.##.#\n#...#E#\n#...E.#\n#######",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "36334",
			},
		},
	},
	{
		Input: "#######\n#E..EG#\n#.#G.E#\n#E.##E#\n#G..#.#\n#..E#.#\n#######",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "39514",
			},
			{
				Solver: SolvePart2,
				Output: "31284",
			},
		},
	},
	{
		Input: "#######\n#E.G#.#\n#.#G..#\n#G.#.G#\n#G..#.#\n#...E.#\n#######",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "27755",
			},
			{
				Solver: SolvePart2,
				Output: "3478",
			},
		},
	},
	{
		Input: "#######\n#.E...#\n#.#..G#\n#.###.#\n#E#G#G#\n#...#G#\n#######",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "7,3",
			},
			{
				Solver: SolvePart2,
				Output: "7,3",
			},
		},
	},
	{
		Input: "#########\n#G......#\n#.E.#...#\n#..##..G#\n#...##..#\n#...#...#\n#.G...G.#\n#.....G.#\n#########",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "28944",
			},
			{
				Solver: SolvePart2,
				Output: "6474",
			},
		},
	},
	{
		Input: solver.InputFromFile("input.txt"),
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "250594",
			},
			{
				Solver: SolvePart2,
				Output: "52133",
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
