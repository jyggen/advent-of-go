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

type elf struct {
	y int
	x int
}

type direction struct {
	consider  [3]int
	direction int
}

type order struct {
	elf *elf
	to  [2]int
}

const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func simulate(input string, maxRounds int) (int, int) {
	rows := utils.ToRuneSlice(input, "\n")
	rowLength := len(rows)
	colLength := len(rows[0])
	elves := make([]*elf, 0)
	lowY, highY, lowX, highX := 0, rowLength-1, 0, colLength-1
	positions := make(map[[2]int]struct{}, len(elves))
	directions := []*direction{
		{[3]int{N, NE, NW}, N},
		{[3]int{S, SE, SW}, S},
		{[3]int{W, NW, SW}, W},
		{[3]int{E, NE, SE}, E},
	}

	for y, row := range rows {
		for x, r := range row {
			if r == '.' {
				continue
			}

			elves = append(elves, &elf{y, x})
			positions[[2]int{y, x}] = struct{}{}
		}
	}

	round := 0

	for round != maxRounds {
		round += 1
		orders := make([]*order, 0, len(elves))
		duplicates := make(map[[2]int]int, 0)

		for _, e := range elves {
			var coordsPair [2]int

			shouldMove := false
			hasNeighbour := false

			for _, d := range directions {
				if shouldMove && hasNeighbour {
					break
				}

				var possibleY, possibleX int

				canMove := true

				for _, consider := range d.consider {
					switch consider {
					case N:
						possibleY = e.y - 1
						possibleX = e.x
					case NE:
						possibleY = e.y - 1
						possibleX = e.x + 1
					case E:
						possibleY = e.y
						possibleX = e.x + 1
					case SE:
						possibleY = e.y + 1
						possibleX = e.x + 1
					case S:
						possibleY = e.y + 1
						possibleX = e.x
					case SW:
						possibleY = e.y + 1
						possibleX = e.x - 1
					case W:
						possibleY = e.y
						possibleX = e.x - 1
					case NW:
						possibleY = e.y - 1
						possibleX = e.x - 1
					}

					if _, ok := positions[[2]int{possibleY, possibleX}]; ok {
						hasNeighbour = true
						canMove = false
						break
					}
				}

				if !canMove || shouldMove {
					continue
				}

				shouldMove = true

				switch d.direction {
				case N:
					coordsPair = [2]int{e.y - 1, e.x}
				case E:
					coordsPair = [2]int{e.y, e.x + 1}
				case S:
					coordsPair = [2]int{e.y + 1, e.x}
				case W:
					coordsPair = [2]int{e.y, e.x - 1}
				}
			}

			if !shouldMove || !hasNeighbour {
				continue
			}

			orders = append(orders, &order{e, coordsPair})

			if v, ok := duplicates[coordsPair]; !ok {
				duplicates[coordsPair] = 1
			} else {
				duplicates[coordsPair] = v + 1
			}
		}

		moved := 0

		for _, o := range orders {
			if duplicates[o.to] != 1 {
				continue
			}

			moved += 1

			delete(positions, [2]int{o.elf.y, o.elf.x})
			positions[o.to] = struct{}{}

			o.elf.y, o.elf.x = o.to[0], o.to[1]
			lowY, highY = utils.MinInt(o.elf.y, lowY), utils.MaxInt(o.elf.y, highY)
			lowX, highX = utils.MinInt(o.elf.x, lowX), utils.MaxInt(o.elf.x, highX)
		}

		if moved == 0 {
			break
		}

		directions = append(directions[1:], directions[0])
	}

	return ((highY - lowY + 1) * (highX - lowX + 1)) - len(elves), round
}

func SolvePart1(input string) (string, error) {
	empty, _ := simulate(input, 10)

	return strconv.Itoa(empty), nil
}

func SolvePart2(input string) (string, error) {
	_, rounds := simulate(input, -1)

	return strconv.Itoa(rounds), nil
}
