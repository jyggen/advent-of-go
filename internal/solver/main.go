package solver

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Solver func(input string) (string, error)

type TestCaseSolver struct {
	Solver Solver
	Output string
}

type TestCase struct {
	Short   bool
	Input   string
	Solvers []*TestCaseSolver
}

func (tc *TestCase) Benchmark(b *testing.B) {
	if testing.Short() && tc.Short == false {
		b.SkipNow()
	}

	for j, solver := range tc.Solvers {
		b.Run(fmt.Sprint(j), func(subtest *testing.B) {
			for i := 0; i < subtest.N; i++ {
				_, err := solver.Solver(tc.Input)
				if err != nil {
					subtest.Error(err)
				}
			}
		})
	}
}

func (tc *TestCase) Test(t *testing.T) {
	if testing.Short() && tc.Short == false {
		t.SkipNow()
	}

	t.Parallel()

	for j, solver := range tc.Solvers {
		t.Run(fmt.Sprint(j), func(subtest *testing.T) {
			subtest.Parallel()

			actualOutput, err := solver.Solver(tc.Input)

			assert.NoError(subtest, err)
			assert.Equal(subtest, solver.Output, actualOutput)
		})
	}
}

func InputFromFile(name string) string {
	input, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Replace(string(input), "\r", "", -1)
}

func SolveFromFile(f *os.File, s1 Solver, s2 Solver) (string, string, error) {
	input, err := ioutil.ReadAll(f)
	if err != nil {
		return "", "", err
	}

	inputStr := strings.Replace(string(input), "\r", "", -1)

	part1, err := s1(inputStr)
	if err != nil {
		return part1, "", err
	}

	part2, err := s2(inputStr)

	return part1, part2, err
}
