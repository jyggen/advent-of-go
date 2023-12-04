package main

import (
	"fmt"
	"os"
	"slices"
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

type blueprint struct {
	oreRobotCost      int
	clayRobotCost     int
	obsidianRobotCost [2]int
	geodeRobotCost    [2]int
}

type state struct {
	minute         int
	ore            int
	oreRobots      int
	clay           int
	clayRobots     int
	obsidian       int
	obsidianRobots int
	geode          int
	geodeRobots    int
	didNotBuild    []int
}

func (s *state) copy() *state {
	return &state{
		s.minute,
		s.ore,
		s.oreRobots,
		s.clay,
		s.clayRobots,
		s.obsidian,
		s.obsidianRobots,
		s.geode,
		s.geodeRobots,
		make([]int, 0),
	}
}

func (s *state) key() [7]int {
	return [7]int{
		s.minute,
		s.ore,
		s.clay,
		s.obsidian,
		s.oreRobots,
		s.clayRobots,
		s.obsidianRobots,
	}
}

const (
	ORE = iota
	CLAY
	OBSIDIAN
	GEODE
)

func bruteforce(b *blueprint, minutes int) int {
	queue := []*state{
		{
			minute:      0,
			oreRobots:   1,
			didNotBuild: make([]int, 0),
		},
	}

	var s *state

	cache := make(map[[7]int]struct{})
	highest := 0

	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]

		s.minute++
		s.ore += s.oreRobots
		s.clay += s.clayRobots
		s.obsidian += s.obsidianRobots
		s.geode += s.geodeRobots

		if s.minute == minutes {
			highest = utils.MaxInt(highest, s.geode)

			continue
		}

		cacheKey := s.key()

		if _, ok := cache[cacheKey]; ok {
			continue
		}

		cache[cacheKey] = struct{}{}
		couldNotBuildOne := false
		didNotBuild := make([]int, len(s.didNotBuild))

		copy(didNotBuild, s.didNotBuild)

		if len(didNotBuild) == 4 {
			continue
		}

		if !slices.Contains(didNotBuild, ORE) && (s.oreRobots < b.oreRobotCost || s.oreRobots < b.clayRobotCost || s.oreRobots < b.obsidianRobotCost[0] || s.oreRobots < b.geodeRobotCost[0]) {
			if (s.ore - s.oreRobots) >= b.oreRobotCost {
				s2 := s.copy()
				s2.ore -= b.oreRobotCost
				s2.oreRobots++
				queue = append(queue, s2)
				didNotBuild = append(didNotBuild, ORE) // nozero
			} else {
				couldNotBuildOne = true
			}
		}

		if !slices.Contains(didNotBuild, CLAY) && s.clayRobots < b.obsidianRobotCost[1] {
			if (s.ore - s.oreRobots) >= b.clayRobotCost {
				s2 := s.copy()
				s2.ore -= b.clayRobotCost
				s2.clayRobots++
				queue = append(queue, s2)
				didNotBuild = append(didNotBuild, CLAY) // nozero
			} else {
				couldNotBuildOne = true
			}
		}

		if !slices.Contains(didNotBuild, OBSIDIAN) && s.obsidianRobots < b.geodeRobotCost[1] {
			if (s.ore-s.oreRobots) >= b.obsidianRobotCost[0] && (s.clay-s.clayRobots) >= b.obsidianRobotCost[1] {
				s2 := s.copy()
				s2.ore -= b.obsidianRobotCost[0]
				s2.clay -= b.obsidianRobotCost[1]
				s2.obsidianRobots++
				queue = append(queue, s2)
				didNotBuild = append(didNotBuild, OBSIDIAN) // nozero
			} else {
				couldNotBuildOne = true
			}
		}

		if !slices.Contains(didNotBuild, GEODE) && (s.ore-s.oreRobots) >= b.geodeRobotCost[0] && (s.obsidian-s.obsidianRobots) >= b.geodeRobotCost[1] {
			s2 := s.copy()
			s2.ore -= b.geodeRobotCost[0]
			s2.obsidian -= b.geodeRobotCost[1]
			s2.geodeRobots++
			queue = append(queue, s2)
			didNotBuild = append(didNotBuild, GEODE) // nozero
		} else {
			couldNotBuildOne = true
		}

		if couldNotBuildOne {
			s2 := s.copy()
			s2.didNotBuild = didNotBuild
			queue = append(queue, s2)
		}
	}

	return highest
}

func SolvePart1(input string) (string, error) {
	numbers := utils.ToOptimisticIntSlice(input, false)
	blueprints := make([]*blueprint, 0)

	for i := 0; i < len(numbers); i += 7 {
		blueprints = append(blueprints, &blueprint{
			numbers[i+1],
			numbers[i+2],
			[2]int{numbers[i+3], numbers[i+4]},
			[2]int{numbers[i+5], numbers[i+6]},
		})
	}

	sum := 0

	for i, b := range blueprints {
		sum += (i + 1) * bruteforce(b, 24)
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	numbers := utils.ToOptimisticIntSlice(input, false)
	blueprints := make([]*blueprint, 0)

	for i := 0; i < len(numbers); i += 7 {
		blueprints = append(blueprints, &blueprint{
			numbers[i+1],
			numbers[i+2],
			[2]int{numbers[i+3], numbers[i+4]},
			[2]int{numbers[i+5], numbers[i+6]},
		})

		if len(blueprints) == 3 {
			break
		}
	}

	sum := 1

	for _, b := range blueprints {
		sum *= bruteforce(b, 32)
	}

	return strconv.Itoa(sum), nil
}
