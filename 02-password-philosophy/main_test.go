package main

import (
	"github.com/jyggen/advent-of-go-utils/solver"
	"io/ioutil"
	"testing"
)

func TestSolvers(t *testing.T) {
	t.Run("examples", func(subtest *testing.T) {
		solver.TestSolver(subtest, SolvePart1, "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc", "2")
		solver.TestSolver(subtest, SolvePart2, "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc", "1")
	})
	t.Run("input", func(subtest *testing.T) {
		input, err := ioutil.ReadFile("input.txt")

		if err != nil {
			subtest.Error(err)
		}

		solver.TestSolver(subtest, SolvePart1, string(input), "572")
		solver.TestSolver(subtest, SolvePart2, string(input), "306")
	})
}
