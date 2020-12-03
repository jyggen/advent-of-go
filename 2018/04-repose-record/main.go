package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type guard struct {
	id      int
	total   int
	minutes []int
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	guards := makeGuards(utils.ToStringSlice(input, "\n"))

	var bestCount, bestId, bestMinute int

	for _, g := range guards {
		if bestId == 0 || guards[bestId].total < g.total {
			bestCount = 0
			bestId = g.id
			bestMinute = 0
		}

		for minute, count := range g.minutes {
			if bestId == g.id && count > bestCount {
				bestCount = count
				bestMinute = minute
			}
		}
	}

	return strconv.Itoa(bestId * bestMinute), nil
}

func SolvePart2(input string) (string, error) {
	guards := makeGuards(utils.ToStringSlice(input, "\n"))

	var bestCount, bestId, bestMinute int

	for _, g := range guards {
		for minute, count := range g.minutes {
			if count > bestCount {
				bestCount = count
				bestId = g.id
				bestMinute = minute
			}
		}
	}

	return strconv.Itoa(bestId * bestMinute), nil
}

func makeGuards(input []string) map[int]*guard {
	sort.Strings(input)

	var asleepAt time.Time
	var current *guard

	guards := make(map[int]*guard, 0)
	guardIdRegex := regexp.MustCompile("#\\d+")

	for _, line := range input {
		date, _ := time.Parse("2006-01-02 15:04", line[1:17])

		switch line[19:24] {
		case "Guard":
			guardId, _ := strconv.Atoi(guardIdRegex.FindString(line)[1:])

			if _, ok := guards[guardId]; !ok {
				guards[guardId] = &guard{
					id:      guardId,
					total:   0,
					minutes: make([]int, 59),
				}
			}

			current = guards[guardId]
		case "falls":
			asleepAt = date
		case "wakes":
			current.total += date.Minute() - asleepAt.Minute()

			for i := asleepAt.Minute(); i < date.Minute(); i++ {
				current.minutes[i]++
			}
		}
	}

	return guards
}
