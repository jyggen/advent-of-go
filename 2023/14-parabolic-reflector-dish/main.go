package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

const (
	cube  = '#'
	empty = '.'
	round = 'O'
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func north(platform [][]byte) {
	for x := 0; x < len(platform[0]); x++ {
		lastObstacle := -1

		for y := 0; y < len(platform); y++ {
			switch platform[y][x] {
			case cube:
				lastObstacle = y
			case empty:
				continue
			case round:
				lastObstacle++

				platform[y][x] = empty
				platform[lastObstacle][x] = round
			}
		}
	}
}

func west(platform [][]byte) {
	for y := 0; y < len(platform); y++ {
		lastObstacle := -1

		for x := 0; x < len(platform[0]); x++ {
			switch platform[y][x] {
			case cube:
				lastObstacle = x
			case empty:
				continue
			case round:
				lastObstacle++

				platform[y][x] = empty
				platform[y][lastObstacle] = round
			}
		}
	}
}

func south(platform [][]byte) {
	for x := 0; x < len(platform[0]); x++ {
		lastObstacle := len(platform)

		for y := len(platform) - 1; y >= 0; y-- {
			switch platform[y][x] {
			case cube:
				lastObstacle = y
			case empty:
				continue
			case round:
				lastObstacle--

				platform[y][x] = empty
				platform[lastObstacle][x] = round
			}
		}
	}
}

func east(platform [][]byte) {
	for y := 0; y < len(platform); y++ {
		lastObstacle := len(platform[0])

		for x := len(platform[0]) - 1; x >= 0; x-- {
			switch platform[y][x] {
			case cube:
				lastObstacle = x
			case empty:
				continue
			case round:
				lastObstacle--

				platform[y][x] = empty
				platform[y][lastObstacle] = round
			}
		}
	}
}

func load(platform [][]byte) int {
	sum := 0
	height := len(platform)

	for y, row := range platform {
		for _, column := range row {
			if column == round {
				sum += height - y
			}
		}
	}

	return sum
}

func SolvePart1(input string) (string, error) {
	platform := utils.ToByteSlice(input, '\n')

	north(platform)

	return strconv.Itoa(load(platform)), nil
}

func SolvePart2(input string) (string, error) {
	platform := utils.ToByteSlice(input, '\n')
	cycles := 1000000000
	looped := false
	cache := make(map[string]int, 100)
	cache[string(bytes.Join(platform, nil))] = 0

	for i := 1; i <= cycles; i++ {
		north(platform)
		west(platform)
		south(platform)
		east(platform)

		if !looped {
			key := string(bytes.Join(platform, nil))

			if v, ok := cache[key]; ok {
				i = cycles - (cycles-v)%(i-v)
				looped = true
			} else {
				cache[key] = i
			}
		}
	}

	return strconv.Itoa(load(platform)), nil
}
