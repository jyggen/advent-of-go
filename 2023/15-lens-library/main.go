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

func hash(characters []rune) int {
	value := rune(0)

	for _, c := range characters {
		value += c
		value *= 17
		value %= 256
	}

	return int(value)
}

func SolvePart1(input string) (string, error) {
	sequence := utils.ToRuneSlice(input, ",")
	sum := 0

	for _, characters := range sequence {
		sum += hash(characters)
	}

	return strconv.Itoa(sum), nil
}

type lens struct {
	label       []rune
	focalLength int
}

func SolvePart2(input string) (string, error) {
	sequence := utils.ToRuneSlice(input, ",")
	boxes := make([][]*lens, 256)

	for i := range boxes {
		boxes[i] = make([]*lens, 0)
	}

	for _, characters := range sequence {
		if characters[len(characters)-1] == '-' {
			label := characters[:len(characters)-1]
			h := hash(label)
			lensIndex := slices.IndexFunc(boxes[h], func(l *lens) bool {
				return slices.Equal(l.label, label)
			})

			if lensIndex != -1 {
				boxes[h] = append(boxes[h][:lensIndex], boxes[h][lensIndex+1:]...)
			}
		} else {
			index := strings.IndexRune(string(characters), '=')
			label := characters[:index]
			h := hash(label)
			focalLength, _ := strconv.Atoi(string(characters[index+1:]))
			lensIndex := slices.IndexFunc(boxes[h], func(l *lens) bool {
				return slices.Equal(l.label, label)
			})

			if lensIndex != -1 {
				boxes[h][lensIndex].focalLength = focalLength
			} else {
				boxes[h] = append(boxes[h], &lens{label, focalLength})
			}
		}
	}

	focusingPower := 0

	for b, lenses := range boxes {
		for i, l := range lenses {
			focusingPower += (1 + b) * (i + 1) * l.focalLength
		}
	}

	return strconv.Itoa(focusingPower), nil
}
