package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

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
	best, _ := simulate(input)

	return strconv.Itoa(best), nil
}

func SolvePart2(input string) (string, error) {
	_, result := simulate(input)

	return strconv.Itoa(result), nil
}

func simulate(input string) (int, int) {
	parts := utils.ToStringSlice(input, ", ")
	xParts := utils.ToStringSlice(parts[0][15:], "..")
	yParts := utils.ToStringSlice(parts[1][2:], "..")
	x1, _ := strconv.Atoi(xParts[0])
	x2, _ := strconv.Atoi(xParts[1])
	y1, _ := strconv.Atoi(yParts[0])
	y2, _ := strconv.Atoi(yParts[1])
	validInits := 0
	bestY := 0

	for xInit := int(math.Sqrt(float64(x1 * 2))); xInit <= x2; xInit++ {
		for yInit := y1; yInit < -y1+1; yInit++ {
			probe := [2]int{0, 0}
			xVel := xInit
			yVel := yInit
			highest := 0

			for probe[0] < x2 && probe[1] > y1 {
				probe[0] += xVel
				probe[1] += yVel

				if xVel > 0 {
					xVel--
				} else if xVel < 0 {
					xVel++
				}

				if probe[1] > highest {
					highest = probe[1]
				}

				yVel--

				if probe[0] >= x1 && probe[0] <= x2 && probe[1] >= y1 && probe[1] <= y2 {
					validInits++

					if highest > bestY {
						bestY = highest
					}
					break
				}
			}
		}
	}

	return bestY, validInits
}
