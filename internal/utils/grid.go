package utils

import (
	"strings"
)

const (
	North     = iota
	NorthEast = iota
	East      = iota
	SouthEast = iota
	South     = iota
	SouthWest = iota
	West      = iota
	NorthWest = iota
)

func FromCoordinates(x int, y int, columnLength int) int {
	return x + (columnLength * y)
}

func ToCoordinates(offset int, columnLength int) (x int, y int) {
	x = offset % columnLength
	y = offset / columnLength

	return x, y
}

func Neighbour(x int, y int, direction int, rowLength int, columnLength int) int {
	switch direction {
	case North:
		y--
	case NorthEast:
		y--
		x++
	case East:
		x++
	case SouthEast:
		y++
		x++
	case South:
		y++
	case SouthWest:
		y++
		x--
	case West:
		x--
	case NorthWest:
		y--
		x--
	}

	if x < 0 || y < 0 || x >= columnLength || y >= rowLength {
		return -1
	}

	return FromCoordinates(x, y, columnLength)
}

func GridToString(grid []rune, columnLength int) string {
	var sb strings.Builder

	for offset, cell := range grid {
		x, y := ToCoordinates(offset, columnLength)

		if x == 0 && y != 0 {
			sb.WriteString("\n")
		}

		sb.WriteString(string(cell))
	}

	return sb.String()
}
