package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"sort"
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

type rule struct {
	lowerX int
	upperX int
	Y      int
}

func compileRules(input string, max int) []rule {
	coordinates := utils.ToOptimisticIntSlice(input)
	rules := make([]rule, 0)

	//now := time.Now()

	for i := 0; i < len(coordinates); i += 4 {
		distance := utils.ManhattanDistance(coordinates[i]-coordinates[i+2], coordinates[i+1]-coordinates[i+3])

		for k, l := distance, 0; k >= 0; k, l = k-1, l+1 {
			rules = append(rules, rule{
				coordinates[i] - k,
				coordinates[i] + k,
				coordinates[i+1] - l,
			}, rule{
				coordinates[i] - k,
				coordinates[i] + k,
				coordinates[i+1] + l,
			})
		}
	}

	//fmt.Println("Create Rules", time.Now().Sub(now))
	//now = time.Now()

	sort.Slice(rules, func(i, j int) bool {
		if rules[i].Y == rules[j].Y {
			if rules[i].lowerX == rules[j].lowerX {
				return rules[i].upperX < rules[j].upperX
			}
			return rules[i].lowerX < rules[j].lowerX
		}

		return rules[i].Y < rules[j].Y
	})

	//fmt.Println("Sort Rules", time.Now().Sub(now))
	//now = time.Now()

	newRules := make([]rule, 0, len(rules))

	for i := 0; i < len(rules); i++ {
		if rules[i].Y < 0 || rules[i].Y > max {
			continue
		}

		if len(newRules) == 0 {
			newRules = append(newRules, rules[i])
			continue
		}

		lastRule := newRules[len(newRules)-1]

		if rules[i].Y != lastRule.Y || rules[i].lowerX > lastRule.upperX+1 {
			newRules = append(newRules, rules[i])
		} else {
			newRules[len(newRules)-1].lowerX = utils.MinInt(lastRule.lowerX, rules[i].lowerX)
			newRules[len(newRules)-1].upperX = utils.MaxInt(lastRule.upperX, rules[i].upperX)
		}
	}

	//fmt.Println("Reduce Rules", time.Now().Sub(now))

	return newRules
}

func SolvePart1(input string) (string, error) {
	rules := compileRules(input, 2000000)
	sum := 0

	for _, r := range rules {
		if r.Y < 2000000 {
			continue
		}

		if r.Y > 2000000 {
			break
		}

		sum += r.upperX - r.lowerX
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	rules := compileRules(input, 4000000)

	for _, r := range rules {
		if r.Y < 0 || r.Y > 4000000 || (r.lowerX <= 0 && r.upperX >= 4000000) {
			continue
		}

		return strconv.Itoa((r.upperX+1)*4000000 + r.Y), nil
	}

	return "-1", nil
}
