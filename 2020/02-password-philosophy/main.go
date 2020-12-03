package main

import (
	"errors"
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"regexp"
	"strconv"
)

var inputRegex *regexp.Regexp

func init() {
	inputRegex = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)
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
	rules := utils.ToStringSlice(input, "\n")
	valid := 0

	for _, rule := range rules {
		password, letter, lowerLimit, upperLimit, err := parseRule(rule)

		if err != nil {
			return "", err
		}

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
		password, letter, i1, i2, err := parseRule(rule)

		if err != nil {
			return "", err
		}

		i1--
		i2--

		if password[i1] != password[i2] && (password[i1] == letter || password[i2] == letter) {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}

func parseRule(rule string) ([]rune, rune, int, int, error) {
	match := inputRegex.FindStringSubmatch(rule)

	if match == nil {
		return nil, 0, 0, 0, errors.New("unable to parse input")
	}

	lowerLimit, err := strconv.Atoi(match[1])

	if err != nil {
		return nil, 0, 0, 0, err
	}

	upperLimit, err := strconv.Atoi(match[2])

	if err != nil {
		return nil, 0, 0, 0, err
	}

	letter := []rune(match[3])[0]
	password := []rune(match[4])

	return password, letter, lowerLimit, upperLimit, nil
}
