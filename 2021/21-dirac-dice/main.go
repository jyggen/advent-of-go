package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type player struct {
	position int
	score    int
	next     *player
}

func SolvePart1(input string) (string, error) {
	integers := utils.ToOptimisticIntSlice(input, true)
	player1, player2 := &player{position: integers[1] - 1, score: 0}, &player{position: integers[3] - 1, score: 0}
	player1.next = player2
	player2.next = player1
	current := player1
	dice := 0
	rolls := 0

	for player1.score < 1000 && player2.score < 1000 {
		for i := 0; i < 3; i++ {
			current.position = (current.position + dice + 1) % 10
			dice = (dice + 1) % 100
			rolls += 1
		}

		current.score += current.position + 1
		current = current.next
	}

	return strconv.Itoa(utils.MinInt(player1.score, player2.score) * rolls), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(0), nil
}
