package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
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
	nice := 0
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	nasty := []string{"ab", "cd", "pq", "xy"}

	for _, s := range utils.ToStringSlice(input, "\n") {
		numVowels := 0
		hasDouble := false
		hasNasty := false

		for i := 0; i < len(s); i++ {
			if numVowels < 3 && slices.Contains(vowels, rune(s[i])) {
				numVowels++
			}

			if i == 0 {
				continue
			}

			if !hasDouble && s[i] == s[i-1] {
				hasDouble = true
			}

			if slices.Contains(nasty, s[i-1:i+1]) {
				hasNasty = true

				break
			}
		}

		if numVowels == 3 && hasDouble && !hasNasty {
			nice++
		}
	}

	return strconv.Itoa(nice), nil
}

func SolvePart2(input string) (string, error) {
	nice := 0

	for _, s := range utils.ToStringSlice(input, "\n") {
		hasDoublePair := false
		hasRepeat := false

		for i, r := range []rune(s) {
			if i+2 >= len(s) {
				continue
			}

			if !hasDoublePair && strings.Contains(s[i+2:], s[i:i+2]) {
				hasDoublePair = true
			}

			if !hasRepeat && rune(s[i+2]) == r {
				hasRepeat = true
			}

			if hasDoublePair && hasRepeat {
				nice++

				break
			}
		}
	}

	return strconv.Itoa(nice), nil
}
