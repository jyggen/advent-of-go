package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

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

	recipes := []int{3, 7}
	elves := []int{0, 1}
	recipesLen := 2
	stop := numOfRecipes + 10
	solutionOne := ""

	for {
		sum := 0

		for j, e := range elves {
			sum += recipes[e]
			elves[j] += recipes[e] + 1
		}

		digits := strings.Split(strconv.Itoa(sum), "")

		for _, d := range digits {
			d, _ := strconv.Atoi(d)
			recipes = append(recipes, d)
			recipesLen++

			if recipesLen == stop {
				for _, i := range recipes[numOfRecipes:] {
					solutionOne += strconv.Itoa(i)
				}

				return solutionOne, nil
			}
		}

		for j, e := range elves {
			elves[j] = e % recipesLen
		}
	}
}

func SolvePart2(input string) (string, error) {
	input = strings.TrimSpace(input)
	recipes := []int{3, 7}
	elves := []int{0, 1}
	recipesLen := 2
	scoreHistory := strings.Repeat("0", len(input))
	solutionTwo := 0

	for {
		sum := 0

		for j, e := range elves {
			sum += recipes[e]
			elves[j] += recipes[e] + 1
		}

		digits := strings.Split(strconv.Itoa(sum), "")

		for _, d := range digits {
			scoreHistory = scoreHistory[1:] + d

			if solutionTwo == 0 && scoreHistory == input {
				return strconv.Itoa(len(recipes) - len(input) + 1), nil
			}

			d, _ := strconv.Atoi(d)
			recipes = append(recipes, d)
			recipesLen++
		}

		for j, e := range elves {
			elves[j] = e % recipesLen
		}
	}
}
