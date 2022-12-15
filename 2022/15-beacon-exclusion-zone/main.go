package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"math"
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

type rule struct {
	lowerX int
	upperX int
	Y      int
}

func compileRules(input string, max int) ([]rule, [][2]int) {
	coordinates := utils.ToOptimisticIntSlice(input, true)
	rules := make([]rule, 0)
	bounds := make([][2]int, max+1)

	for i := 0; i < len(coordinates); i += 4 {
		distance := utils.ManhattanDistance(coordinates[i]-coordinates[i+2], coordinates[i+1]-coordinates[i+3])

		for k, l := distance, 0; k >= 0; k, l = k-1, l+1 {
			if coordinates[i+1]-l >= 0 && coordinates[i+1]-l <= max {
				rules = append(rules, rule{
					coordinates[i] - k,
					coordinates[i] + k,
					coordinates[i+1] - l,
				})
			}

			if coordinates[i+1]+l >= 0 && coordinates[i+1]+l <= max {
				rules = append(rules, rule{
					coordinates[i] - k,
					coordinates[i] + k,
					coordinates[i+1] + l,
				})
			}
		}
	}

	for {
		queue := make([]rule, 0, len(rules))

		for _, r := range rules {
			if bounds[r.Y][0] == 0 && bounds[r.Y][1] == 0 {
				bounds[r.Y][0] = r.lowerX
				bounds[r.Y][1] = r.upperX
			} else {
				if bounds[r.Y][1] >= r.upperX && bounds[r.Y][0] <= r.lowerX {
					bounds[r.Y][0] = utils.MinInt(r.lowerX, bounds[r.Y][0])
					bounds[r.Y][1] = utils.MaxInt(r.upperX, bounds[r.Y][1])
				} else if bounds[r.Y][1] <= r.upperX && bounds[r.Y][0] >= r.lowerX {
					bounds[r.Y][0] = utils.MinInt(r.lowerX, bounds[r.Y][0])
					bounds[r.Y][1] = utils.MaxInt(r.upperX, bounds[r.Y][1])
				} else if bounds[r.Y][1]+1 >= r.lowerX && bounds[r.Y][1] < r.upperX {
					bounds[r.Y][0] = utils.MinInt(r.lowerX, bounds[r.Y][0])
					bounds[r.Y][1] = utils.MaxInt(r.upperX, bounds[r.Y][1])
				} else if bounds[r.Y][0] <= r.upperX+1 && bounds[r.Y][0] > r.lowerX {
					bounds[r.Y][0] = utils.MinInt(r.lowerX, bounds[r.Y][0])
					bounds[r.Y][1] = utils.MaxInt(r.upperX, bounds[r.Y][1])
				} else {
					queue = append(queue, r)
				}
			}
		}

		if len(rules) == len(queue) {
			break
		}

		rules = queue
	}

	return rules, bounds
}

func SolvePart1(input string) (string, error) {
	_, bounds := compileRules(input, 2000000)

	return strconv.Itoa(bounds[2000000][1] - bounds[2000000][0]), nil
}

func SolvePart2(input string) (string, error) {
	rules, _ := compileRules(input, 4000000)

	min := math.MaxInt

	for _, r := range rules {
		min = utils.MinInt(min, r.lowerX)
	}

	return strconv.Itoa((min-1)*4000000 + rules[0].Y), nil
}
