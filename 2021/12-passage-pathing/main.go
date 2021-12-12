package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type item struct {
	pos             string
	duplicatedSmall bool
	available       []string
}

func parseInput(input string) map[string][]string {
	remaining := utils.ToStringSlice(input, "\n")
	connections := make(map[string][]string)

	for _, r := range remaining {
		parts := utils.ToStringSlice(r, "-")

		for _, v := range [2][2]string{
			{parts[0], parts[1]},
			{parts[1], parts[0]},
		} {
			if _, ok := connections[v[0]]; !ok {
				connections[v[0]] = make([]string, 0, len(connections))
			}

			connections[v[0]] = append(connections[v[0]], v[1])
		}
	}

	return connections
}

func traverse(connections map[string][]string, i item, allowDuplicates bool) int {
	if i.pos == "end" {
		return 1
	}

	paths := 0

	for _, o := range connections[i.pos] {
		if o == "start" {
			continue
		}

		a := make([]string, len(i.available))
		copy(a, i.available)

		duplicated := i.duplicatedSmall

		if !utils.IsUpper(o) {
			visited := true

			for _, p := range a {
				if p == o {
					visited = false
					break
				}
			}

			if visited {
				if duplicated {
					continue
				} else {
					duplicated = true
				}
			}
		}

		for k, v := range i.available {
			if v == o {
				a = append(a[:k], a[k+1:]...)
			}
		}

		paths += traverse(connections, item{
			pos:             o,
			duplicatedSmall: duplicated,
			available:       a,
		}, allowDuplicates)
	}

	return paths
}

func SolvePart1(input string) (string, error) {
	connections := parseInput(input)
	available := make([]string, 0, len(connections))

	for k := range connections {
		if k != "start" {
			available = append(available, k)
		}
	}

	paths := traverse(connections, item{
		pos:             "start",
		duplicatedSmall: false,
		available:       available,
	}, false)

	return strconv.Itoa(paths), nil
}

func SolvePart2(input string) (string, error) {
	connections := parseInput(input)
	available := make([]string, 0, len(connections))

	for k := range connections {
		if k != "start" {
			available = append(available, k)
		}
	}

	paths := traverse(connections, item{
		pos:             "start",
		duplicatedSmall: false,
		available:       available,
	}, true)

	return strconv.Itoa(paths), nil
}
