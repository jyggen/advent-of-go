package main

import (
	"container/list"
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
	"strings"
)

type item struct {
	key      string
	distance int
}

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	rows := utils2.ToStringSlice(strings.TrimSpace(input), "\n")
	orbits := make(map[string]string, len(rows))

	for _, row := range rows {
		parts := utils2.ToStringSlice(row, ")")
		orbits[parts[1]] = parts[0]
	}

	checksum := 0

	for _, orbit := range orbits {
		checksum += getChecksum(orbits, orbit, 1)
	}

	return strconv.Itoa(checksum - 1), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils2.ToStringSlice(strings.TrimSpace(input), "\n")
	orbits := make(map[string][]string, len(rows))

	for _, row := range rows {
		parts := utils2.ToStringSlice(row, ")")

		if _, ok := orbits[parts[1]]; !ok {
			orbits[parts[1]] = make([]string, 0)
		}

		orbits[parts[1]] = append(orbits[parts[1]], parts[0])

		if _, ok := orbits[parts[0]]; !ok {
			orbits[parts[0]] = make([]string, 0)
		}

		orbits[parts[0]] = append(orbits[parts[0]], parts[1])
	}

	target := orbits["SAN"][0]
	queue := list.New()
	queue.PushBack(item{orbits["YOU"][0], 0})
	visited := make(map[string]struct{})

	for queue.Len() > 0 {
		e := queue.Front()
		asItem := e.Value.(item)
		visited[asItem.key] = struct{}{}

		if asItem.key == target {
			return strconv.Itoa(asItem.distance), nil
		}

		for _, dest := range orbits[asItem.key] {
			if _, ok := visited[dest]; !ok {
				queue.PushBack(item{dest, asItem.distance + 1})
			}
		}

		queue.Remove(e)
	}

	return "", nil
}

func getChecksum(orbits map[string]string, orbit string, depth int) int {
	if _, ok := orbits[orbits[orbit]]; ok {
		return getChecksum(orbits, orbits[orbit], depth+1)
	}

	return depth + 1
}
