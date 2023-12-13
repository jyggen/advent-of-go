package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func diffHorizontal(a []byte, b []byte) int {
	count := 0

	for i := range a {
		if a[i] != b[i] {
			count++
		}

		if count > 1 {
			return 2
		}
	}

	return count
}

func diffVertical(p [][]byte, a int, b int) int {
	count := 0

	for i := 0; i < len(p); i++ {
		if p[i][a] != p[i][b] {
			count++
		}

		if count > 1 {
			return 2
		}
	}

	return count
}

func findMirrors(g [][]byte, accepted int) int {
HorizontalLoop:
	for i := 1; i < len(g); i++ {
		limit := min(i, len(g)-i)
		sum := 0

		for j := 0; j < limit; j++ {
			sum += diffHorizontal(g[i-1-j], g[i+j])

			if sum > accepted {
				continue HorizontalLoop
			}
		}

		if sum != accepted {
			continue HorizontalLoop
		}

		return i * 100
	}

VerticalLoop:
	for i := 1; i < len(g[0]); i++ {
		limit := min(i, len(g[0])-i)
		sum := 0

		for j := 0; j < limit; j++ {
			sum += diffVertical(g, i-1-j, i+j)

			if sum > accepted {
				continue VerticalLoop
			}
		}

		if sum != accepted {
			continue VerticalLoop
		}

		return i
	}

	return -1
}

func SolvePart1(input string) (string, error) {
	patterns := utils.ToStringSlice(input, "\n\n")
	sum := 0

	for _, p := range patterns {
		g := utils.ToByteSlice(p, '\n')
		mirrorAt := findMirrors(g, 0)

		if mirrorAt == -1 {
			return "", errors.New("unable to find mirrors")
		}

		sum += mirrorAt
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	patterns := utils.ToStringSlice(input, "\n\n")
	sum := 0

	for _, p := range patterns {
		g := utils.ToByteSlice(p, '\n')
		mirrorAt := findMirrors(g, 1)

		if mirrorAt == -1 {
			return "", errors.New("unable to find mirrors")
		}

		sum += mirrorAt
	}

	return strconv.Itoa(sum), nil
}
