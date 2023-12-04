package solver

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
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

var enableProfiling = flag.Bool("profile", false, "enable cpu profiling")

func (tc *TestCase) Benchmark(b *testing.B) {
	if testing.Short() && tc.Short == false {
		b.SkipNow()
	}

	for j, solver := range tc.Solvers {
		b.Run(fmt.Sprint(j), func(subtest *testing.B) {
			subtest.ResetTimer()

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

	for j, solver := range tc.Solvers {
		solver := solver
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
	flag.Parse()

	input, err := ioutil.ReadAll(f)
	if err != nil {
		return "", "", err
	}

	inputStr := strings.Replace(string(input), "\r", "", -1)

	part1, err := solvePart(inputStr, s1, "part1.prof")

	if err != nil {
		return part1, "", err
	}

	part2, err := solvePart(inputStr, s2, "part2.prof")

	return part1, part2, err
}

func solvePart(input string, solver Solver, profName string) (string, error) {
	if *enableProfiling {
		f, err := os.Create(profName)
		if err != nil {
			return "", err
		}

		defer f.Close()

		err = pprof.StartCPUProfile(f)
		if err != nil {
			return "", err
		}

		defer pprof.StopCPUProfile()
	}

	return solver(input)
}
