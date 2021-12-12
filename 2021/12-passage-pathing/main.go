package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

type item struct {
	pos             string
	duplicatedSmall bool
	history         []string
}

func SolvePart1(input string) (string, error) {
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

	queue := make([]item, 0, len(connections))
	queue = append(queue, item{
		pos:     "start",
		history: make([]string, 0, len(connections)),
	})

	paths := 0

	var i item

	for len(queue) > 0 {
		i, queue = queue[len(queue)-1], queue[:len(queue)-1]
		i.history = append(i.history, i.pos)

		if i.pos == "end" {
			paths++
			continue
		}

	Outer:
		for _, o := range connections[i.pos] {
			h := make([]string, len(i.history))
			copy(h, i.history)

			if o != strings.ToUpper(o) {
				for _, p := range h {
					if p == o {
						continue Outer
					}
				}
			}

			queue = append(queue, item{
				pos:     o,
				history: h,
			})
		}
	}

	return strconv.Itoa(paths), nil
}

func SolvePart2(input string) (string, error) {
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

	queue := make([]item, 0, len(connections))
	queue = append(queue, item{
		pos:             "start",
		duplicatedSmall: false,
		history:         make([]string, 0, len(connections)),
	})

	paths := 0

	var i item

	for len(queue) > 0 {
		i, queue = queue[len(queue)-1], queue[:len(queue)-1]
		i.history = append(i.history, i.pos)

		if i.pos == "end" {
			paths++
			continue
		}

	Outer:
		for _, o := range connections[i.pos] {
			if o == "start" {
				continue
			}

			h := make([]string, len(i.history))
			copy(h, i.history)

			duplicated := i.duplicatedSmall

			if o != strings.ToUpper(o) {
				for _, p := range h {
					if p == o {
						if !duplicated {
							duplicated = true
							break
						}
						continue Outer
					}
				}
			}

			queue = append(queue, item{
				pos:             o,
				duplicatedSmall: duplicated,
				history:         h,
			})
		}
	}

	return strconv.Itoa(paths), nil
}
