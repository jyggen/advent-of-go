package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/stretchr/testify/assert"
)

var testCases = []*solver.TestCase{
	{
		Input: "1=-0-2\n12111\n2=0=\n21\n2=01\n111\n20012\n112\n1=-1=\n1-12\n12\n1=\n122",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "2=-1=0",
			},
			/*{
				Solver: SolvePart2,
				Output: "1623178306",
			},*/
		},
	},
	/*{

		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "7225",
			},
			{
				Solver: SolvePart2,
				Output: "548634267428",
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

type snafuTestCase struct {
	decimal int
	snafu   string
}

var snafusTestCases = []snafuTestCase{
	{1, "1"},
	{2, "2"},
	{3, "1="},
	{4, "1-"},
	{5, "10"},
	{6, "11"},
	{7, "12"},
	{8, "2="},
	{9, "2-"},
	{10, "20"},
	{15, "1=0"},
	{20, "1-0"},
	{2022, "1=11-2"},
	{12345, "1-0---0"},
	{314159265, "1121-1110-1=0"},
}

func TestFromSnafu(t *testing.T) {
	for i, snafu := range snafusTestCases {
		t.Run(fmt.Sprint(i), func(subtest *testing.T) {
			assert.Equal(subtest, snafu.decimal, fromSnafu(snafu.snafu))
		})
	}
}

func TestToSnafu(t *testing.T) {
	for i, snafu := range snafusTestCases {
		t.Run(fmt.Sprint(i), func(subtest *testing.T) {
			assert.Equal(subtest, snafu.snafu, toSnafu(snafu.decimal))
		})
	}
}
