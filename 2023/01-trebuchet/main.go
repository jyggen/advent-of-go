package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

var letters = map[string]rune{
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

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToRuneSlice(input, "\n")
	sum := 0

	for _, r := range rows {
		first := rune(-1)
		last := rune(0)

		for _, c := range r {
			if c >= '1' && c <= '9' {
				if first == -1 {
					first = c
				}

				last = c
			}
		}

		number, _ := strconv.Atoi(string(first) + string(last))
		sum += number
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToRuneSlice(input, "\n")
	sum := 0

	for _, r := range rows {
		first := rune(-1)
		last := rune(0)
		segment := strings.Builder{}

		for _, c := range r {
			if segment.Len() != 0 {
				maybe := segment.String()
				for letter, k := range letters {
					if strings.HasSuffix(maybe, letter) {
						if first == -1 {
							first = k
						}

						last = k
					}
				}
			}

			if c >= '1' && c <= '9' {
				segment = strings.Builder{}

				if first == -1 {
					first = c
				}

				last = c
			} else {
				segment.WriteRune(c)
			}
		}

		if segment.Len() != 0 {
			maybe := segment.String()
			for letter, k := range letters {
				if strings.HasSuffix(maybe, letter) {
					if first == -1 {
						first = k
					}

					last = k
				}
			}
		}

		number, _ := strconv.Atoi(string(first) + string(last))
		sum += number
	}

	return strconv.Itoa(sum), nil
}
