package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	ingredientsCount, ingredients, _ := resolve(input)

	sum := 0

	for i, c := range ingredientsCount {
		if _, ok := ingredients[i]; !ok {
			sum += c
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	_, _, allergens := resolve(input)
	keys := make([]string, 0, len(allergens))

	for k := range allergens {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	output := make([]string, 0, len(allergens))

	for _, k := range keys {
		output = append(output, allergens[k])
	}

	return strings.Join(output, ","), nil
}

func resolve(input string) (map[string]int, map[string]string, map[string]string) {
	assessments := make(map[string]map[string]int, 0)
	ingredientsCount := make(map[string]int, 0)

	for _, r := range utils.ToStringSlice(input, "\n") {
		disclaimerAt := strings.IndexRune(r, '(')
		ingredients := r[0:disclaimerAt]
		disclaimer := r[disclaimerAt+10 : len(r)-1]

		for _, i := range strings.Fields(ingredients) {
			if _, ok := ingredientsCount[i]; !ok {
				ingredientsCount[i] = 0
			}

			ingredientsCount[i]++
		}

		for _, d := range utils.ToStringSlice(disclaimer, ", ") {
			if _, ok := assessments[d]; !ok {
				assessments[d] = make(map[string]int, 0)
			}

			for _, i := range strings.Fields(ingredients) {
				if _, ok := assessments[d][i]; !ok {
					assessments[d][i] = 0
				}

				assessments[d][i]++
			}
		}
	}

	allergens := make(map[string]string, len(assessments))
	ingredients := make(map[string]string, len(assessments))

	shouldRestart := true
	endless := false

	for shouldRestart && !endless {
		shouldRestart = false
		endless = true

		for d, a := range assessments {
			bestName := ""
			bestCount := 0
			isTied := false

			for i, c := range a {
				if _, ok := ingredients[i]; ok {
					continue
				}

				if c > bestCount {
					bestName = i
					bestCount = c
					isTied = false
				} else if c == bestCount {
					isTied = true
				}
			}

			if isTied {
				shouldRestart = true
			} else {
				endless = false
				ingredients[bestName] = d
				allergens[d] = bestName

				delete(assessments, d)
			}
		}
	}

	return ingredientsCount, ingredients, allergens

}
