package main

import (
	"fmt"
	"testing"

	"github.com/jyggen/advent-of-go/internal/solver"
)

var testCases = []*solver.TestCase{
	{
		Input: "D2FE28",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "6",
			},
		},
	},
	{
		Input: "38006F45291200",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "9",
			},
		},
	},
	{
		Input: "EE00D40C823060",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "14",
			},
		},
	},
	{
		Input: "8A004A801A8002F478",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "16",
			},
		},
	},
	{
		Input: "620080001611562C8802118E34",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "12",
			},
		},
	},
	{
		Input: "C0015000016115A2E0802F182340",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "23",
			},
		},
	},
	{
		Input: "A0016C880162017C3686B18A3D4780",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "31",
			},
		},
	},
	{
		Input: "C200B40A82",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "3",
			},
		},
	},
	{
		Input: "04005AC33890",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "54",
			},
		},
	},
	{
		Input: "880086C3E88112",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "7",
			},
		},
	},
	{
		Input: "CE00C43D881120",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "9",
			},
		},
	},
	{
		Input: "D8005AC2A8F0",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "1",
			},
		},
	},
	{
		Input: "F600BC2D8F",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "0",
			},
		},
	},
	{
		Input: "9C005AC2F8F0",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "0",
			},
		},
	},
	{
		Input: "9C0141080250320F1802104A08",
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart2,
				Output: "1",
			},
		},
	},
	{
		Short: true,
		Solvers: []*solver.TestCaseSolver{
			{
				Solver: SolvePart1,
				Output: "938",
			},
			{
				Solver: SolvePart2,
				Output: "1495959086337",
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
