package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
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
	rules := utils.ToStringSlice(input, "\n")
	valid := 0

	for _, rule := range rules {
		password, letter, lowerLimit, upperLimit := parseRule(rule)

		found := 0

		for _, char := range password {
			if char == letter {
				found++
			}
		}

		if found >= lowerLimit && found <= upperLimit {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}

func SolvePart2(input string) (string, error) {
	rules := utils.ToStringSlice(input, "\n")
	valid := 0

	for _, rule := range rules {
		password, letter, i1, i2 := parseRule(rule)

		i1--
		i2--

		if password[i1] != password[i2] && (password[i1] == letter || password[i2] == letter) {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}

func parseRule(rule string) ([]rune, rune, int, int) {
	parts := strings.Split(rule, " ")
	limits := strings.Split(parts[0], "-")
	lowerLimit, _ := strconv.Atoi(limits[0])
	upperLimit, _ := strconv.Atoi(limits[1])
	letter := []rune(parts[1])[0]
	password := []rune(parts[2])

	return password, letter, lowerLimit, upperLimit
}
