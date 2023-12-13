package main

import (
	"bytes"
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

func rotate(g [][]byte) [][]byte {
	g2 := make([][]byte, len(g[0]))

	for i := range g2 {
		g2[i] = make([]byte, len(g))
	}

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			g2[x][len(g)-y-1] = g[y][x]
		}
	}

	return g2
}

func findMirrors(g [][]byte) int {
HorizontalLoop:
	for i := 1; i < len(g); i++ {
		if !bytes.Equal(g[i-1], g[i]) {
			continue HorizontalLoop
		}

		limit := min(i, len(g)-i)

		for j := 0; j < limit; j++ {
			if !bytes.Equal(g[i-1-j], g[i+j]) {
				continue HorizontalLoop
			}
		}

		return i
	}

	return -1
}

func diff(a []byte, b []byte) int {
	if bytes.Equal(a, b) {
		return 0
	}

	count := 0

	for i := range a {
		if a[i] != b[i] {
			count++
		}

		if count > 1 {
			return 2
		}
	}

	return 1
}

func findSmudgeMirrors(g [][]byte) int {
HorizontalLoop:
	for i := 1; i < len(g); i++ {
		if diff(g[i-1], g[i]) == 2 {
			continue HorizontalLoop
		}

		limit := min(i, len(g)-i)
		smudge := false

		for j := 0; j < limit; j++ {
			switch diff(g[i-1-j], g[i+j]) {
			case 1:
				if smudge {
					continue HorizontalLoop
				}

				smudge = true
			case 2:
				continue HorizontalLoop
			}
		}

		if !smudge {
			continue HorizontalLoop
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
		mirrorAt := findMirrors(g)

		if mirrorAt != -1 {
			sum += mirrorAt * 100

			continue
		}

		mirrorAt = findMirrors(rotate(g))

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
		mirrorAt := findSmudgeMirrors(g)

		if mirrorAt != -1 {
			sum += mirrorAt * 100

			continue
		}

		mirrorAt = findSmudgeMirrors(rotate(g))

		if mirrorAt == -1 {
			return "", errors.New("unable to find mirrors")
		}

		sum += mirrorAt
	}

	return strconv.Itoa(sum), nil
}
