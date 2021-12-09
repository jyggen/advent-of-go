package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
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
	tiles := getTiles(input)

	return strconv.Itoa(countBlacks(tiles)), nil
}

func SolvePart2(input string) (string, error) {
	tiles := getTiles(input)

	for i := 0; i < 100; i++ {
		tiles = simulate(tiles)
	}

	return strconv.Itoa(countBlacks(tiles)), nil
}

func countBlacks(tiles map[string]bool) int {
	blacks := 0

	for _, t := range tiles {
		if t {
			blacks++
		}
	}

	return blacks
}

func simulate(tiles map[string]bool) map[string]bool {
	newTiles := make(map[string]bool, len(tiles))
	queue := make([]string, 0, len(tiles))

	for tile := range tiles {
		queue = append(queue, tile)
	}

	for i := 0; i < len(queue); i++ {
		if _, ok := newTiles[queue[i]]; ok {
			continue
		}

		isBlack := false

		if val, ok := tiles[queue[i]]; ok {
			isBlack = val
		}

		blackNeighbours := 0

		for _, n := range getNeighbours(queue[i]) {
			if black, ok := tiles[n]; !ok && isBlack {
				queue = append(queue, n)
			} else if black {
				blackNeighbours++
			}
		}

		if isBlack {
			if blackNeighbours == 0 || blackNeighbours > 2 {
				newTiles[queue[i]] = false
			} else {
				newTiles[queue[i]] = true
			}
		} else {
			if blackNeighbours == 2 {
				newTiles[queue[i]] = true
			} else {
				newTiles[queue[i]] = false
			}
		}
	}

	return newTiles
}

func flip(tiles map[string]bool, coordinates string) {
	if tile, ok := tiles[coordinates]; !ok || !tile {
		tiles[coordinates] = true
	} else {
		tiles[coordinates] = false
	}
}

func getNeighbours(coordinates string) [6]string {
	var x, y int

	fmt.Sscanf(coordinates, "%dx%d", &x, &y)

	return [6]string{
		coordinatesKey(x-1, y),
		coordinatesKey(x+1, y),
		coordinatesKey(x+1, y+1),
		coordinatesKey(x, y+1),
		coordinatesKey(x-1, y-1),
		coordinatesKey(x, y-1),
	}
}

func coordinatesKey(x int, y int) string {
	return strconv.Itoa(x) + "x" + strconv.Itoa(y)
}

func getTiles(input string) map[string]bool {
	rows := utils2.ToRuneSlice(input, "\n")
	tiles := make(map[string]bool, len(rows))

	for _, r := range rows {
		var buffer rune

		x := 0
		y := 0

		for _, c := range r {
			if c == 's' || c == 'n' {
				buffer = c
				continue
			}

			if buffer == 's' {
				switch c {
				case 'e':
					y++
					x++
				case 'w':
					y++
				}
			} else if buffer == 'n' {
				switch c {
				case 'e':
					y--
				case 'w':
					y--
					x--
				}
			} else {
				switch c {
				case 'e':
					x++
				case 'w':
					x--
				}
			}

			buffer = 0
		}

		flip(tiles, coordinatesKey(x, y))
	}

	return tiles
}
