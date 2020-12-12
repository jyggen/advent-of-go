package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
)

const (
	north = 0
	east  = 90
	south = 180
	west  = 270
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
	runeSlice := utils.ToRuneSlice(input, "\n")
	x, y, degrees := 0, 0, east

	for _, v := range runeSlice {
		distance, _ := strconv.Atoi(string(v[1:]))

		for {
			switch v[0] {
			case 'N':
				y -= distance
			case 'S':
				y += distance
			case 'E':
				x += distance
			case 'W':
				x -= distance
			case 'L':
				degrees -= distance
			case 'R':
				degrees += distance
			case 'F':
				switch degrees {
				case north:
					v[0] = 'N'
				case east:
					v[0] = 'E'
				case south:
					v[0] = 'S'
				case west:
					v[0] = 'W'
				}
				continue
			}

			break
		}

		if degrees < 0 {
			degrees = 360 - utils.AbsInt(degrees)
		} else if degrees >= 360 {
			degrees -= 360
		}
	}

	return strconv.Itoa(utils.ManhattanDistance(x, y)), nil
}

func SolvePart2(input string) (string, error) {
	runeSlice := utils.ToRuneSlice(input, "\n")
	shipX, shipY, waypointX, waypointY := 0, 0, 10, -1

	for _, v := range runeSlice {
		distance, _ := strconv.Atoi(string(v[1:]))

		switch v[0] {
		case 'N':
			waypointY -= distance
		case 'S':
			waypointY += distance
		case 'E':
			waypointX += distance
		case 'W':
			waypointX -= distance
		case 'L':
			waypointX, waypointY = utils.RotatePoint(waypointX, waypointY, -distance)
		case 'R':
			waypointX, waypointY = utils.RotatePoint(waypointX, waypointY, distance)
		case 'F':
			shipX += waypointX * distance
			shipY += waypointY * distance
		}
	}

	return strconv.Itoa(utils.ManhattanDistance(shipX, shipY)), nil
}
