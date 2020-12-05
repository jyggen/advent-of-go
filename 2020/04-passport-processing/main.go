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

var byr, ecl, eyr, field, hcl, hgt, iyr, pid *regexp.Regexp
var required []string

func init() {
	byr = regexp.MustCompile(`\bbyr:(?:19[2-9]\d|200[0-2])\b`)
	ecl = regexp.MustCompile(`\becl:(?:amb|blu|brn|gry|grn|hzl|oth)\b`)
	eyr = regexp.MustCompile(`\beyr:20(?:2\d|30)\b`)
	field = regexp.MustCompile(`\b([a-z]{3}):[a-z\d#]+\b`)
	hcl = regexp.MustCompile(`\bhcl:#[\da-f]{6}\b`)
	hgt = regexp.MustCompile(`\bhgt:(?:1(?:[5-8]\d|9[0-3])cm|(?:59|6\d|7[0-6])in)\b`)
	iyr = regexp.MustCompile(`\biyr:20(?:1\d|20)\b`)
	pid = regexp.MustCompile(`\bpid:\d{9}\b`)
	required = []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
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
	passports := utils.ToStringSlice(input, "\n\n")
	valid := 0

PassportLoop:
	for _, p := range passports {
		matches := field.FindAllStringSubmatch(p, -1)

		if matches == nil {
			return "", errors.New("unable to parse input")
		}

		data := make(map[string]bool, len(matches)-1)

		for _, m := range matches {
			data[m[1]] = true
		}

		for _, r := range required {
			if _, ok := data[r]; !ok {
				continue PassportLoop
			}
		}

		valid++
	}

	return strconv.Itoa(valid), nil
}

func SolvePart2(input string) (string, error) {
	passports := utils.ToStringSlice(input, "\n\n")
	valid := 0

PassportLoop:
	for _, p := range passports {
		for _, r := range []*regexp.Regexp{
			byr, ecl, eyr, hcl, hgt, iyr, pid,
		} {
			if !r.MatchString(p) {
				continue PassportLoop
			}
		}

		valid++
	}

	return strconv.Itoa(valid), nil
}
