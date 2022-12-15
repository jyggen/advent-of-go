package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
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
	input = strings.TrimSpace(input)
	numOfRecipes, err := strconv.Atoi(input)
	if err != nil {
		return "", err
	}

	elves := []int{0, 1}
	stop := numOfRecipes + 10
	recipes := make([]int, 0, stop)
	recipes = append(recipes, 3, 7)
	solutionOne := ""

	for {
		sum := 0

		for j, e := range elves {
			sum += recipes[e]
			elves[j] += recipes[e] + 1
		}

		for _, d := range strconv.Itoa(sum) {
			d, _ := strconv.Atoi(string(d))
			recipes = append(recipes, d)

			if len(recipes) == stop {
				for _, i := range recipes[numOfRecipes:] {
					solutionOne += strconv.Itoa(i)
				}

				return solutionOne, nil
			}
		}

		for j, e := range elves {
			elves[j] = e % len(recipes)
		}
	}
}

func SolvePart2(input string) (string, error) {
	input = strings.TrimSpace(input)
	recipes := []int{3, 7}
	elves := [2]int{0, 1}
	scoreHistory := strings.Repeat("0", len(input))

	for {
		sum := 0

		for j, e := range elves {
			sum += recipes[e]
			elves[j] += recipes[e] + 1
		}

		for _, d := range strconv.Itoa(sum) {
			scoreHistory = scoreHistory[1:] + string(d)

			if scoreHistory == input {
				return strconv.Itoa(len(recipes) - len(input) + 1), nil
			}

			d, _ := strconv.Atoi(string(d))
			recipes = append(recipes, d)
		}

		for j, e := range elves {
			elves[j] = e % len(recipes)
		}
	}
}
