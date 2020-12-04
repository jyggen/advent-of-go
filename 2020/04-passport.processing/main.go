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

var ecl, field, hcl, hgt, pid *regexp.Regexp
var required []string

func init() {
	ecl = regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	field = regexp.MustCompile(`\b([a-z]{3}):([a-z\d#]+)\b`)
	hcl = regexp.MustCompile(`^#[\da-f]{6}$`)
	hgt = regexp.MustCompile(`^(\d+)(cm|in)$`)
	pid = regexp.MustCompile(`^\d{9}$`)
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

		data := make(map[string]string, len(matches)-1)

		for _, m := range matches {
			data[m[1]] = m[2]
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
		matches := field.FindAllStringSubmatch(p, -1)

		if matches == nil {
			return "", errors.New("unable to parse input")
		}

		data := make(map[string]string, len(matches)-1)

		for _, m := range matches {
			data[m[1]] = m[2]
		}

		for _, r := range required {
			if _, ok := data[r]; !ok {
				continue PassportLoop
			}
		}

		if !hcl.MatchString(data["hcl"]) || !ecl.MatchString(data["ecl"]) || !pid.MatchString(data["pid"]) {
			continue
		}

		if byr, err := strconv.Atoi(data["byr"]); err != nil || byr < 1920 || byr > 2002 {
			continue
		}

		if byr, err := strconv.Atoi(data["iyr"]); err != nil || byr < 2010 || byr > 2020 {
			continue
		}

		if byr, err := strconv.Atoi(data["eyr"]); err != nil || byr < 2020 || byr > 2030 {
			continue
		}

		match := hgt.FindStringSubmatch(data["hgt"])

		if match == nil {
			continue
		}

		if height, err := strconv.Atoi(match[1]); err != nil || (match[2] == "cm" && (height < 150 || height > 193)) || (match[2] == "in" && (height < 59 || height > 76)) {
			continue
		}

		valid++
	}

	return strconv.Itoa(valid), nil
}
