package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

var digitStrings = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

var firstReplacer *strings.Replacer
var lastReplacer *strings.Replacer

var digitStringsReverse map[string]rune

func init() {
	digitStringsReverse = make(map[string]rune, len(digitStrings))

	for k, v := range digitStrings {
		reversed := []rune(k)

		for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
			reversed[i], reversed[j] = reversed[j], reversed[i]
		}

		digitStringsReverse[string(reversed)] = v
	}

	firstReplacer = strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	lastReplacer = strings.NewReplacer("eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9")
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func containsDigitString(builder strings.Builder, digitStringsMap map[string]rune) (bool, rune) {
	maybeDigitString := builder.String()

	for letter, r := range digitStringsMap {
		if strings.HasSuffix(maybeDigitString, letter) {
			return true, r
		}
	}

	return false, 0
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToRuneSlice(input, "\n")
	sum := 0

	for _, r := range rows {
		var first, last rune

		for _, c := range r {
			if c >= '1' && c <= '9' {
				first = c
				break
			}
		}

		for i := len(r) - 1; i >= 0; i-- {
			if r[i] >= '1' && r[i] <= '9' {
				last = r[i]
				break
			}
		}

		number, _ := strconv.Atoi(string(first) + string(last))
		sum += number
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	sum := 0

	for _, r := range rows {
		var first, last rune

		for _, c := range firstReplacer.Replace(r) {
			if c >= '1' && c <= '9' {
				first = c
				break
			}
		}

		reversed := []rune(r)

		for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
			reversed[i], reversed[j] = reversed[j], reversed[i]
		}

		for _, c := range lastReplacer.Replace(string(reversed)) {
			if c >= '1' && c <= '9' {
				last = c
				break
			}
		}

		number, _ := strconv.Atoi(string(first) + string(last))
		sum += number
	}

	return strconv.Itoa(sum), nil
}
