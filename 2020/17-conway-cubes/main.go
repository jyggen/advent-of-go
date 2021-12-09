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
	cube, height, width, depth, hyper := createCube(input, false)

	for i := 0; i < 6; i++ {
		cube, height, width, depth, hyper = simulate(cube, height, width, depth, hyper, false)
	}

	return strconv.Itoa(countActive(cube, height, width, depth, hyper)), nil
}

func SolvePart2(input string) (string, error) {
	cube, height, width, depth, hyper := createCube(input, true)

	for i := 0; i < 6; i++ {
		cube, height, width, depth, hyper = simulate(cube, height, width, depth, hyper, true)
	}

	return strconv.Itoa(countActive(cube, height, width, depth, hyper)), nil
}

func countActive(cube []bool, height int, width int, depth int, hyper int) int {
	count := 0

	for w := 0; w < hyper; w++ {
		for z := 0; z < depth; z++ {
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if cube[getCubeOffset(x, y, z, w, height, width, depth)] {
						count++
					}
				}
			}
		}
	}

	return count
}

func createCube(input string, increaseHyper bool) ([]bool, int, int, int, int) {
	rows := utils2.ToRuneSlice(input, "\n")
	depth := 3
	height := len(rows) + 2
	width := len(rows[0]) + 2
	hyper := 1
	w := 0

	if increaseHyper {
		w++
		hyper += 2
	}

	cube := make([]bool, depth*height*width*hyper)

	for y, row := range rows {
		for x, state := range row {
			if state == '#' {
				cube[getCubeOffset(x+1, y+1, 1, w, height, width, depth)] = true
			}
		}
	}

	return cube, height, width, depth, hyper
}

func getCubeOffset(x int, y int, z int, w int, height int, width int, depth int) int {
	return x + height*(y+width*(z+depth*w))
}

func getNeighbours(x int, y int, z int, w int, height int, width int, depth int, hyper int) []int {
	neighbours := make([]int, 0, 3*3*3*3-1)

	for nw := utils2.MaxInt(w-1, 0); nw <= w+1 && nw < hyper; nw++ {
		for nz := utils2.MaxInt(z-1, 0); nz <= z+1 && nz < depth; nz++ {
			for ny := utils2.MaxInt(y-1, 0); ny <= y+1 && ny < height; ny++ {
				for nx := utils2.MaxInt(x-1, 0); nx <= x+1 && nx < width; nx++ {
					if nw == w && nz == z && ny == y && nx == x {
						continue
					}

					neighbours = append(neighbours, getCubeOffset(nx, ny, nz, nw, height, width, depth))
				}
			}
		}
	}

	return neighbours
}

func simulate(cube []bool, height int, width int, depth int, hyper int, increaseHyper bool) ([]bool, int, int, int, int) {
	newDepth := depth + 2
	newHeight := height + 2
	newWidth := width + 2
	newHyper := hyper

	if increaseHyper {
		newHyper += 2
	}

	newCube := make([]bool, newDepth*newHeight*newWidth*newHyper)

	for w := 0; w < hyper; w++ {
		for z := 0; z < depth; z++ {
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					oldOffset := getCubeOffset(x, y, z, w, height, width, depth)
					active := 0

					for _, n := range getNeighbours(x, y, z, w, height, width, depth, hyper) {
						if cube[n] {
							active++
						}
					}

					if cube[oldOffset] && active != 2 && active != 3 {
						continue

					}

					if !cube[oldOffset] && active != 3 {
						continue
					}

					newW := w

					if increaseHyper {
						newW += 1
					}

					newOffset := getCubeOffset(x+1, y+1, z+1, newW, newHeight, newWidth, newDepth)
					newCube[newOffset] = true
				}
			}
		}
	}

	return newCube, newHeight, newWidth, newDepth, newHyper
}
