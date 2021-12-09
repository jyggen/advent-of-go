package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

var tree rune

func init() {
	tree = []rune("#")[0]
}

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	rows := parseInput(input)

	return strconv.Itoa(travel(rows, 1, 3)), nil
}

func SolvePart2(input string) (string, error) {
	rows := parseInput(input)
	total := 1

	for _, trees := range []int{
		travel(rows, 1, 1),
		travel(rows, 1, 3),
		travel(rows, 1, 5),
		travel(rows, 1, 7),
		travel(rows, 2, 1),
	} {
		total = total * trees
	}

	return strconv.Itoa(total), nil
}

func parseInput(input string) [][]rune {
	data := utils2.ToStringSlice(input, "\n")
	rows := make([][]rune, len(data))

	for i, row := range data {
		rows[i] = []rune(row)
	}

	return rows
}

func travel(rows [][]rune, down int, right int) int {
	rowLen := len(rows)
	colLen := len(rows[0])
	trees := 0
	x := 0

	for y := 0; y < rowLen; y += down {
		if rows[y][x] == tree {
			trees++
		}

		x = (x + right) % colLen
	}

	return trees
}
