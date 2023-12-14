package main

import (
	"bytes"
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

const (
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
	for y, row := range platform {
		for x, column := range row {
			if column != round {
				continue
			}

			y2 := y - 1

			for ; y2 >= 0; y2-- {
				if platform[y2][x] != empty {
					break
				}

				platform[y2][x] = round
				platform[y2+1][x] = empty
			}

			y2++
		}
	}
}

func west(platform [][]byte) {
	height := len(platform)

	for x := 0; x < len(platform[0]); x++ {
		for y := 0; y < height; y++ {
			if platform[y][x] != round {
				continue
			}

			x2 := x - 1

			for ; x2 >= 0; x2-- {
				if platform[y][x2] != empty {
					break
				}

				platform[y][x2] = round
				platform[y][x2+1] = empty
			}

			x2++
		}
	}
}

func south(platform [][]byte) {
	height := len(platform)

	for y := height - 1; y >= 0; y-- {
		for x, column := range platform[y] {
			if column != round {
				continue
			}

			y2 := y + 1

			for ; y2 < height; y2++ {
				if platform[y2][x] != empty {
					break
				}

				platform[y2][x] = round
				platform[y2-1][x] = empty
			}

			y2--
		}
	}
}

func east(platform [][]byte) {
	height := len(platform)

	for x := len(platform[0]) - 1; x >= 0; x-- {
		for y := height - 1; y >= 0; y-- {
			if platform[y][x] != round {
				continue
			}

			x2 := x + 1

			for ; x2 < len(platform[0]); x2++ {
				if platform[y][x2] != empty {
					break
				}

				platform[y][x2] = round
				platform[y][x2-1] = empty
			}

			x2--
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
	cache := map[string]int{
		string(bytes.Join(platform, nil)): 0,
	}

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
