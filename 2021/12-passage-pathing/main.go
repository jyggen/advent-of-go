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

type cave struct {
	name  string
	small bool
}

type item struct {
	name            string
	duplicatedSmall bool
	available       []string
}

func parseInput(input string) map[string][]cave {
	remaining := utils.ToStringSlice(input, "\n")
	connections := make(map[string][]cave, len(remaining))

	for _, r := range remaining {
		parts := utils.ToStringSlice(r, "-")

		for _, v := range [2][2]string{
			{parts[0], parts[1]},
			{parts[1], parts[0]},
		} {
			p1 := cave{name: v[1], small: false}

			if v[1] != "start" && v[1] != "end" && !utils.IsUpper(v[1]) {
				p1.small = true
			}

			if _, ok := connections[v[0]]; !ok {
				connections[v[0]] = make([]cave, 0, len(connections))
			}

			connections[v[0]] = append(connections[v[0]], p1)
		}
	}

	return connections
}

func traverse(connections map[string][]cave, i item, allowDuplicates bool) int {
	if i.name == "end" {
		return 1
	}

	paths := 0

	for _, o := range connections[i.name] {
		if o.name == "start" {
			continue
		}

		duplicated := i.duplicatedSmall
		a := make([]string, len(i.available))

		copy(a, i.available)

		if o.small {
			visited := true

			for k, v := range a {
				if v == o.name {
					a = append(a[:k], a[k+1:]...)
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

		paths += traverse(connections, item{
			name:            o.name,
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
		if k != "start" && k != "end" && !utils.IsUpper(k) {
			available = append(available, k)
		}
	}

	paths := traverse(connections, item{
		name:            "start",
		duplicatedSmall: false,
		available:       available,
	}, false)

	return strconv.Itoa(paths), nil
}

func SolvePart2(input string) (string, error) {
	connections := parseInput(input)
	available := make([]string, 0, len(connections))

	for k := range connections {
		if k != "start" && k != "end" && !utils.IsUpper(k) {
			available = append(available, k)
		}
	}

	paths := traverse(connections, item{
		name:            "start",
		duplicatedSmall: false,
		available:       available,
	}, true)

	return strconv.Itoa(paths), nil
}
