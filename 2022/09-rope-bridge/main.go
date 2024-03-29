package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func applyRopePhysics(head [2]int, tail [2]int) [2]int {
	distanceY := head[0] - tail[0]
	distanceX := head[1] - tail[1]

	if distanceY >= 1 && distanceX > 1 || distanceY > 1 && distanceX >= 1 {
		tail[0]++
		tail[1]++
	} else if distanceY <= -1 && distanceX > 1 || distanceY < -1 && distanceX >= 1 {
		tail[0]--
		tail[1]++
	} else if distanceY >= 1 && distanceX < -1 || distanceY > 1 && distanceX <= -1 {
		tail[0]++
		tail[1]--
	} else if distanceY <= -1 && distanceX < -1 || distanceY < -1 && distanceX <= -1 {
		tail[0]--
		tail[1]--
	} else if distanceY > 1 && distanceX == 0 {
		tail[0]++
	} else if distanceY < -1 && distanceX == 0 {
		tail[0]--
	} else if distanceY == 0 && distanceX > 1 {
		tail[1]++
	} else if distanceY == 0 && distanceX < -1 {
		tail[1]--
	}

	return tail
}

func simulate(input string, knotsCount int) int {
	instructions := utils.ToStringSlice(input, "\n")
	visited := [][2]int{{0, 0}}
	knots := make([][2]int, knotsCount)

	for k := range knots {
		knots[k] = [2]int{0, 0}
	}

	for _, instruction := range instructions {
		times, _ := strconv.Atoi(instruction[2:])
	Loop:
		for i := 0; i < times; i++ {
			switch instruction[0] {
			case 'U':
				knots[0][0]--
			case 'R':
				knots[0][1]++
			case 'D':
				knots[0][0]++
			case 'L':
				knots[0][1]--
			}

			for j := 1; j < knotsCount; j++ {
				newCoords := applyRopePhysics(knots[j-1], knots[j])

				if knots[j] == newCoords {
					continue Loop
				}

				knots[j] = newCoords
			}

			visited = append(visited, knots[knotsCount-1])
		}
	}

	visitedMap := make(map[[2]int]struct{}, len(visited))

	for _, v := range visited {
		if _, ok := visitedMap[v]; !ok {
			visitedMap[v] = struct{}{}
		}
	}

	return len(visitedMap)
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(simulate(input, 2)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(simulate(input, 10)), nil
}
