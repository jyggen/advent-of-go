package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
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
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	runeSlice := utils2.ToRuneSlice(input, "\n")
	x, y, degrees := 0, 0, east

	for _, v := range runeSlice {
		distance, _ := strconv.Atoi(string(v[1:]))

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
			degrees = (degrees - distance + 360) % 360
		case 'R':
			degrees = (degrees + distance) % 360
		case 'F':
			switch degrees {
			case north:
				y -= distance
			case east:
				x += distance
			case south:
				y += distance
			case west:
				x -= distance
			}
		}
	}

	return strconv.Itoa(utils2.ManhattanDistance(x, y)), nil
}

func SolvePart2(input string) (string, error) {
	runeSlice := utils2.ToRuneSlice(input, "\n")
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
			waypointX, waypointY = utils2.RotateRelativePoint(waypointX, waypointY, -distance)
		case 'R':
			waypointX, waypointY = utils2.RotateRelativePoint(waypointX, waypointY, distance)
		case 'F':
			shipX += waypointX * distance
			shipY += waypointY * distance
		}
	}

	return strconv.Itoa(utils2.ManhattanDistance(shipX, shipY)), nil
}
