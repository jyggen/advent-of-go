package main

import (
	"container/ring"
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	var totalPlayers, endingPoints int

	_, err := fmt.Sscanf(strings.TrimSpace(input), "%d players; last marble is worth %d points", &totalPlayers, &endingPoints)

	if err != nil {
		panic(err)
	}

	return strconv.Itoa(play(totalPlayers, endingPoints)), nil
}

func SolvePart2(input string) (string, error) {
	var totalPlayers, endingPoints int

	_, err := fmt.Sscanf(strings.TrimSpace(input), "%d players; last marble is worth %d points", &totalPlayers, &endingPoints)

	if err != nil {
		panic(err)
	}

	return strconv.Itoa(play(totalPlayers, endingPoints*100)), nil
}

func play(totalPlayers int, endingPoints int) int {
	circle := &ring.Ring{Value: 0}
	scores := make([]int, totalPlayers)

	for i := 1; i <= endingPoints; i++ {
		if i%23 == 0 {
			circle = circle.Move(-6)
			removed := circle.Move(-2).Link(circle)
			scores[i%totalPlayers] += i + removed.Value.(int)
		} else {
			circle = circle.Next()
			circle.Link(&ring.Ring{Value: i})
			circle = circle.Next()
		}
	}

	best := 0

	for _, score := range scores {
		if score > best {
			best = score
		}
	}

	return best
}
