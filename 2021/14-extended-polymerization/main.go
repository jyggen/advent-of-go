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

func parseInput(input string) (map[[2]byte]byte, map[[2]byte]int, map[byte]int) {
	sections := utils.ToStringSlice(input, "\n\n")
	word := []byte(sections[0])
	rules := make(map[[2]byte]byte)

	for _, rule := range utils.ToStringSlice(sections[1], "\n") {
		from := []byte(rule[0:2])
		to := rule[6]
		rules[[2]byte{from[0], from[1]}] = to
	}

	stateMap := make(map[[2]byte]int)
	letterCount := make(map[byte]int)

	for r := byte('A'); r < byte('Z'); r++ {
		letterCount[r] = 0
	}

	letterCount[word[0]]++

	for i, j := 0, 1; j < len(word); i, j = i+1, j+1 {
		key := [2]byte{word[i], word[j]}

		if _, ok := stateMap[key]; !ok {
			stateMap[key] = 0
		}

		letterCount[word[j]]++
		stateMap[key]++
	}

	return rules, stateMap, letterCount
}

func score(letterCount map[byte]int) int {
	low := math.MaxInt64
	high := 0

	for _, v := range letterCount {
		if v > high {
			high = v
		}

		if v > 0 && v < low {
			low = v
		}
	}

	return high - low
}

func transform(rules map[[2]byte]byte, stateMap map[[2]byte]int, letterCount map[byte]int) (map[[2]byte]int, map[byte]int) {
	newStateMap := make(map[[2]byte]int, len(stateMap))

	for k, v := range stateMap {
		if rule, ok := rules[k]; ok {
			letterCount[rule] += v

			for _, newKey := range [2][2]byte{{k[0], rule}, {rule, k[1]}} {
				if _, ok = newStateMap[newKey]; !ok {
					newStateMap[newKey] = 0
				}

				newStateMap[newKey] += v
			}
		} else {
			if _, ok = newStateMap[k]; !ok {
				newStateMap[k] = 0
			}

			newStateMap[k] += v
		}
	}

	return newStateMap, letterCount
}

func SolvePart1(input string) (string, error) {
	rules, stateMap, letterCount := parseInput(input)

	for round := 0; round < 10; round++ {
		stateMap, letterCount = transform(rules, stateMap, letterCount)
	}

	return strconv.Itoa(score(letterCount)), nil
}

func SolvePart2(input string) (string, error) {
	rules, stateMap, letterCount := parseInput(input)

	for round := 0; round < 40; round++ {
		stateMap, letterCount = transform(rules, stateMap, letterCount)
	}

	return strconv.Itoa(score(letterCount)), nil
}
