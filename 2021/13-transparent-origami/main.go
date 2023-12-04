package main

import (
	"bytes"
	"fmt"
	"github.com/esimov/stackblur-go"
	"github.com/otiai10/gosseract/v2"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
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

type fold struct {
	direction uint8
	index     int
}

func parseInput(input string) ([][]bool, []fold) {
	sections := utils.ToStringSlice(input, "\n\n")
	intSlice := utils.ToOptimisticIntSlice(sections[0], false)
	dots := make([][2]int, 0, len(intSlice)/2)
	maxX, maxY := 0, 0

	for i, j := 0, 1; i < len(intSlice); i, j = i+2, j+2 {
		y, x := intSlice[i], intSlice[j]

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}

		dots = append(dots, [2]int{x, y})
	}

	grid := make([][]bool, maxX+1)

	for i := range grid {
		grid[i] = make([]bool, maxY+1)
	}

	for _, dot := range dots {
		grid[dot[0]][dot[1]] = true
	}

	stringSlice := utils.ToStringSlice(sections[1], "\n")
	folds := make([]fold, 0, len(stringSlice))

	for _, s := range stringSlice {
		s = s[11:]
		direction := s[0]
		index, _ := strconv.Atoi(s[2:])
		folds = append(folds, fold{direction, index})
	}

	return grid, folds
}

func doFold(grid [][]bool, f fold) [][]bool {
	if f.direction == 'y' {
		for x := f.index; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				if !grid[x][y] {
					continue
				}

				grid[f.index-(x-f.index)][y] = grid[x][y]
			}
		}

		grid = grid[:f.index]
	} else {
		for x := 0; x < len(grid); x++ {
			for y := f.index; y < len(grid[x]); y++ {
				if !grid[x][y] {
					continue
				}

				grid[x][f.index-(y-f.index)] = grid[x][y]
			}

			grid[x] = grid[x][:f.index]
		}
	}

	return grid
}

func toImage(grid [][]bool) image.Image {
	maxX := len(grid)
	maxY := len(grid[0])
	imageMultiplier := 8
	black := color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: (maxY * imageMultiplier) + (5 * 2), Y: (maxX * imageMultiplier) + (5 * 2)},
	})

	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)

	for x, rows := range grid {
		for y, l := range rows {
			var c color.RGBA

			if l {
				c = black
			} else {
				c = white
			}

			for y2 := 0; y2 < imageMultiplier; y2++ {
				for x2 := 0; x2 < imageMultiplier; x2++ {
					img.Set(5+(y*imageMultiplier+y2), 5+(x*imageMultiplier+x2), c)
				}
			}
		}
	}

	blurred, _ := stackblur.Process(img, 1)

	return blurred
}

func SolvePart1(input string) (string, error) {
	grid, folds := parseInput(input)
	grid = doFold(grid, folds[0])
	sum := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] {
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	grid, folds := parseInput(input)

	for _, f := range folds {
		grid = doFold(grid, f)
	}

	img := toImage(grid)
	buf := new(bytes.Buffer)
	client := gosseract.NewClient()
	defer client.Close()

	jpeg.Encode(buf, img, &jpeg.Options{Quality: 100})
	client.SetWhitelist("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	client.SetVariable("load_system_dawg", "F")
	client.SetVariable("load_freq_dawg", "F")
	client.SetImageFromBytes(buf.Bytes())
	client.SetPageSegMode(gosseract.PSM_SINGLE_WORD)

	return client.Text()
}
