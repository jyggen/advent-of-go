package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "10",
			},
			{
				Solver: SolvePart2,
				Output: "36",
			},
		},
	},
	{
		Input: "dc-end\nHN-start\nstart-kj\ndc-start\ndc-HN\nLN-dc\nHN-end\nkj-sa\nkj-HN\nkj-dc",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "19",
			},
			{
				Solver: SolvePart2,
				Output: "103",
			},
		},
	},
	{
		Input: "fs-end\nhe-DX\nfs-he\nstart-DX\npj-DX\nend-zg\nzg-sl\nzg-pj\npj-he\nRW-he\nfs-DX\npj-RW\nzg-RW\nstart-pj\nhe-WI\nzg-he\npj-fs\nstart-RW",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "226",
			},
			{
				Solver: SolvePart2,
				Output: "3509",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "4378",
			},
			{
				Solver: SolvePart2,
				Output: "133621",
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
