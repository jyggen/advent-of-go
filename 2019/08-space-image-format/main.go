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
	"math"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

const width = 25
const height = 6

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func makeLayers(input string) ([][]int, int) {
	data, _ := utils.ToIntegerSlice(input, "")
	layerSize := width * height
	numOfLayers := len(data) / layerSize
	layers := make([][]int, numOfLayers)
	offset, sum, minZeros := 0, 0, math.MaxInt16

	for i := range layers {
		layers[i] = make([]int, layerSize)
		zeros, ones, twos := 0, 0, 0

		for j := range layers[i] {
			layers[i][j] = data[offset]

			switch data[offset] {
			case 0:
				zeros++
			case 1:
				ones++
			case 2:
				twos++
			}

			offset++
		}

		if zeros < minZeros {
			minZeros = zeros
			sum = ones * twos
		}
	}

	return layers, sum
}

func toImage(grid []bool) image.Image {
	maxX := width
	maxY := height
	imageMultiplier := 2
	black := color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: (maxX * imageMultiplier) + (5 * 2), Y: (maxY * imageMultiplier) + (5 * 2)},
	})

	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)

	for i, pixel := range grid {
		y := i % width
		x := i / width
		var c color.RGBA

		if pixel {
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

	blurred, _ := stackblur.Run(img, 1)

	return blurred
}

func SolvePart1(input string) (string, error) {
	_, sum := makeLayers(input)

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	layers, _ := makeLayers(input)
	pixels := make([]bool, width*height)

	for i := 0; i < width*height; i++ {
	LayerLoop:
		for l := range layers {
			switch layers[l][i] {
			case 0:
				break LayerLoop
			case 1:
				pixels[i] = true
				break LayerLoop
			}
		}
	}

	img := toImage(pixels)
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
