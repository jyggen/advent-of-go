package main

import (
	"fmt"
	"os"
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

type board struct {
	numbers []int
}

func (b *board) Mark(number int) {
	for i, n := range b.numbers {
		if n == number {
			b.numbers[i] = -1

			break
		}
	}
}

func (b *board) Score() int {
	sum := 0

	for _, n := range b.numbers {
		if n != -1 {
			sum += n
		}
	}

	return sum
}

func (b *board) IsSolved() bool {
	for i := 0; i < 5; i++ {
		solved := true

		for j := i * 5; j < (i+1)*5; j++ {
			if b.numbers[j] != -1 {
				solved = false

				break
			}
		}

		if solved {
			return true
		}

		solved = true

		for j := i; j < len(b.numbers); j += 5 {
			if b.numbers[j] != -1 {
				solved = false

				break
			}
		}

		if solved {
			return true
		}
	}

	return false
}

func SolvePart1(input string) (string, error) {
	numbers, boards, err := parseInput(input)
	if err != nil {
		return "", err
	}

	for i, n := range numbers {
		for _, b := range boards {
			b.Mark(n)

			if i <= 4 {
				continue
			}

			if b.IsSolved() {
				return strconv.Itoa(b.Score() * n), nil
			}
		}
	}

	return "", fmt.Errorf("no bingo achieved")
}

func SolvePart2(input string) (string, error) {
	numbers, boards, err := parseInput(input)
	if err != nil {
		return "", err
	}

	for i, n := range numbers {
		for j := 0; j < len(boards); j++ {
			boards[j].Mark(n)

			if i <= 4 {
				continue
			}

			if boards[j].IsSolved() {
				if len(boards) == 1 {
					return strconv.Itoa(boards[j].Score() * n), nil
				}

				boards[j] = boards[len(boards)-1]
				boards = boards[:len(boards)-1]
				j--
			}
		}
	}

	return "", fmt.Errorf("no bingos achieved")
}

func parseInput(input string) ([]int, []*board, error) {
	rows := utils.ToStringSlice(input, "\n")
	numbers, err := utils.ToIntegerSlice(rows[0], ",")

	if err != nil {
		return numbers, make([]*board, 0), err
	}

	rawBoards := utils.ToStringSlice(strings.Join(rows[2:], "\n"), "\n\n")
	boards := make([]*board, len(rawBoards))

	for i, rawBoard := range rawBoards {
		fields := strings.Fields(rawBoard)
		boards[i] = &board{
			numbers: make([]int, len(fields)),
		}

		for j, number := range fields {
			boards[i].numbers[j], err = strconv.Atoi(number)

			if err != nil {
				return numbers, boards, err
			}
		}
	}

	return numbers, boards, nil
}
