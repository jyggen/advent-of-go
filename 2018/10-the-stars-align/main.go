package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"regexp"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"github.com/otiai10/gosseract/v2"
)

type light struct {
	x         int
	y         int
	velocityX int
	velocityY int
}

var inputRegex = regexp.MustCompile("^position=< *([\\-\\d]+), *([\\-\\d]+)> velocity=< *([\\-\\d]+), *([\\-\\d]+)>$")

func (l *light) Move() {
	l.x += l.velocityX
	l.y += l.velocityY
}

func (l *light) Reverse() {
	l.x -= l.velocityX
	l.y -= l.velocityY
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	img, _ := solve(utils.ToStringSlice(input, "\n"))
	buf := new(bytes.Buffer)

	client := gosseract.NewClient()
	defer client.Close()

	jpeg.Encode(buf, img, &jpeg.Options{Quality: 100})
	client.SetWhitelist("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	client.SetVariable("load_system_dawg", "F")
	client.SetVariable("load_freq_dawg", "F")
	client.SetImageFromBytes(buf.Bytes())

	return client.Text()
}

func SolvePart2(input string) (string, error) {
	_, seconds := solve(utils.ToStringSlice(input, "\n"))

	return strconv.Itoa(seconds), nil
}

func solve(input []string) (image.Image, int) {
	lights := make([]*light, len(input))

	for k, v := range input {
		var x, y, velocityX, velocityY int

		match := inputRegex.FindStringSubmatch(v)

		if len(match) == 0 {
			panic(fmt.Sprintf("unable to parse input \"%s\"", v))
		}

		x, _ = strconv.Atoi(match[1])
		y, _ = strconv.Atoi(match[2])
		velocityX, _ = strconv.Atoi(match[3])
		velocityY, _ = strconv.Atoi(match[4])
		lights[k] = &light{
			x:         x,
			y:         y,
			velocityX: velocityX,
			velocityY: velocityY,
		}
	}

	var sky [][]bool

	area := 0
	i := 0

	var minX, minY, maxX, maxY int

	for {

		minX, minY, maxX, maxY = 0, 0, 0, 0

		for _, l := range lights {
			l.Move()

			if l.x > maxX {
				maxX = l.x
			}

			if l.y > maxY {
				maxY = l.y
			}

			if l.x < minX {
				minX = l.x
			}

			if l.y < minY {
				minY = l.y
			}
		}

		maxY = maxY - minY + 1
		maxX = maxX - minX + 1
		newArea := maxY * maxX

		if area != 0 && newArea > area {
			sky = make([][]bool, maxY)

			for x := range sky {
				sky[x] = make([]bool, maxX)
			}

			for _, l := range lights {
				l.Reverse()
				sky[l.y-minY][l.x-minX] = true
			}

			break
		}

		area = newArea
		i++
	}

	imageMultiplier := 2
	black := color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: maxX * imageMultiplier, Y: maxY * imageMultiplier},
	})

	for y, lights := range sky {
		for x, l := range lights {
			var c color.RGBA

			if l {
				c = black
			} else {
				c = white
			}

			for y2 := 0; y2 < imageMultiplier; y2++ {
				for x2 := 0; x2 < imageMultiplier; x2++ {
					img.Set(x*imageMultiplier+x2, y*imageMultiplier+y2, c)
				}
			}
		}
	}

	return img, i
}
