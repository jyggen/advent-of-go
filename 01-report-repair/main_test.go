package main

import (
	"github.com/jyggen/advent-of-go-utils/solver"
	"io/ioutil"
	"testing"
)

func TestSolvers(t *testing.T) {
	t.Run("example #1", func(subtest *testing.T) {
		solver.TestSolver(subtest, SolvePart1, "1721\n979\n366\n299\n675\n1456", "514579")
		solver.TestSolver(subtest, SolvePart2, "1721\n979\n366\n299\n675\n1456", "241861950")
	})
	t.Run("actual input", func(subtest *testing.T) {
		input, err := ioutil.ReadFile("input.txt")

		if err != nil {
			subtest.Error(err)
		}

		solver.TestSolver(subtest, SolvePart1, string(input), "926464")
		solver.TestSolver(subtest, SolvePart2, string(input), "65656536")
	})
}
