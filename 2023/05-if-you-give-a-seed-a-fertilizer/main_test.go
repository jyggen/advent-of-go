package main

import (
	"strconv"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "35",
			},
			{
				Solver: SolvePart2,
				Output: "46",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "993500720",
			},
			{
				Solver: SolvePart2,
				Output: "4917124",
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
