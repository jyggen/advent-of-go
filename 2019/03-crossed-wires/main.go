package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

const x = 0
const y = 1

var up, right, down, left rune

func init() {
	up = []rune("U")[0]
	right = []rune("R")[0]
	down = []rune("D")[0]
	left = []rune("L")[0]
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
	wires := utils2.ToStringSlice(input, "\n")

	wireOne, err := toPoints(wires[0])

	if err != nil {
		return "", nil
	}

	wireTwo, err := toPoints(wires[1])

	if err != nil {
		return "", nil
	}

	closest := -1

	for i, _ := range wireOne {
		if i == 0 {
			continue
		}

		a1 := wireOne[i-1]
		b1 := wireOne[i]

		for j, _ := range wireTwo {
			if j == 0 {
				continue
			}

			a2 := wireTwo[j-1]
			b2 := wireTwo[j]

			for _, pairs := range [][][][]int{
				{{a1, b1}, {a2, b2}},
				{{a2, b2}, {a1, b1}},
			} {
				if isOverlapping(pairs[0], pairs[1]) {
					distance := utils2.ManhattanDistance(pairs[0][0][x], pairs[1][0][y])

					if closest == -1 || distance < closest {
						closest = distance
					}
				}
			}
		}
	}
	return strconv.Itoa(closest), nil
}

func SolvePart2(input string) (string, error) {
	wires := utils2.ToStringSlice(input, "\n")

	wireOne, err := toPoints(wires[0])

	if err != nil {
		return "", nil
	}

	wireTwo, err := toPoints(wires[1])

	if err != nil {
		return "", nil
	}

	lowest := -1
	wireOneSteps := 0

	for i, _ := range wireOne {
		if i == 0 {
			continue
		}

		a1 := wireOne[i-1]
		b1 := wireOne[i]

		if a1[x] == b1[x] {
			wireOneSteps += utils2.AbsInt(a1[y] - b1[y])
		} else {
			wireOneSteps += utils2.AbsInt(a1[x] - b1[x])
		}

		wireTwoSteps := 0

		for j, _ := range wireTwo {
			if j == 0 {
				continue
			}

			a2 := wireTwo[j-1]
			b2 := wireTwo[j]

			if a2[x] == b2[x] {
				wireTwoSteps += utils2.AbsInt(a2[y] - b2[y])
			} else {
				wireTwoSteps += utils2.AbsInt(a2[x] - b2[x])
			}

			for _, pairs := range [][][][]int{
				{{a1, b1}, {a2, b2}},
				{{a2, b2}, {a1, b1}},
			} {
				if isOverlapping(pairs[0], pairs[1]) {
					score := wireTwoSteps + wireOneSteps
					score -= utils2.AbsInt(pairs[0][1][y] - pairs[1][0][y])
					score -= utils2.AbsInt(pairs[1][1][x] - pairs[0][0][x])

					if lowest == -1 || score < lowest {
						lowest = score
					}
				}
			}
		}
	}
	return strconv.Itoa(int(lowest)), nil
}

func isOverlapping(a [][]int, b [][]int) bool {
	if a[0][x] == a[1][x] && b[0][y] == b[1][y] {
		if (a[0][y] > b[0][y] && a[1][y] < b[0][y]) || (a[0][y] < b[0][y] && a[1][y] > b[0][y]) {
			if (a[0][x] > b[0][x] && a[0][x] < b[1][x]) || (a[0][x] < b[0][x] && a[0][x] > b[1][x]) {
				return true
			}
		}
	}

	return false
}

func toPoints(wire string) ([][]int, error) {
	moves := utils2.ToRuneSlice(wire, ",")
	wired := make([][]int, len(moves)+1)
	x, y := 0, 0
	wired[0] = []int{x, y}

	for i, move := range moves {
		moveInt, err := strconv.Atoi(string(move[1:]))

		if err != nil {
			return wired, err
		}

		switch move[0] {
		case up:
			y -= moveInt
		case right:
			x += moveInt
		case down:
			y += moveInt
		case left:
			x -= moveInt
		}

		wired[i+1] = []int{x, y}
	}

	return wired, nil
}
