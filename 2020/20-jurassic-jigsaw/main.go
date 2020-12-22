package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"math"
	"os"
	"strconv"
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
	variants := getTileVariants(input)
	corners := findCorners(variants)
	sum := 1

	for _, c := range corners {
		sum *= c.id
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	variants := getTileVariants(input)
	corners := findCorners(variants)
	tileSet := buildTileSet(variants, corners)

	for _, t := range tileSet {
		removeBorders(t)
	}

	data := mergeTiles(tileSet)

	width := int(math.Sqrt(float64(len(data))))
	finds := false
	rotations := 0

	for !finds {
		for y := 1; y < width-1; y++ {
		MonsterLoop:
			for x := 19; x < width; x++ {
				monster := [][2]int{
					{x - 1, y - 1},
					{x, y},
					{x - 1, y},
					{x - 2, y},
					{x - 7, y},
					{x - 8, y},
					{x - 13, y},
					{x - 14, y},
					{x - 19, y},
					{x - 3, y + 1},
					{x - 6, y + 1},
					{x - 9, y + 1},
					{x - 12, y + 1},
					{x - 15, y + 1},
					{x - 18, y + 1},
				}

				for _, m := range monster {
					if !data[width*m[1]+m[0]] {
						continue MonsterLoop
					}
				}

				for _, m := range monster {
					data[width*m[1]+m[0]] = false
				}

				finds = true
			}
		}

		if !finds {
			rotations += 1

			if rotations > 3 {
				data = flip(data, width)
				rotations = 0
			}

			data = rotate(data, width, rotations)
		}
	}

	rough := 0

	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			if data[width*y+x] {
				rough++
			}
		}
	}

	return strconv.Itoa(rough), nil

}

type tileVariants struct {
	id       int
	height   int
	width    int
	variants []*tile
}

type tile struct {
	id      int
	variant int
	height  int
	width   int
	sides   [4]int
	data    []bool
}

func getSideIds(data []bool, width int, height int) [4]int {
	return [4]int{
		getSideId(data, 0, width, 1),                       // top
		getSideId(data, width-1, height*width, width),      // right
		getSideId(data, (height-1)*width, height*width, 1), // bottom
		getSideId(data, 0, height*width, width),            // left
	}
}

func getSideId(data []bool, i int, max int, increment int) int {
	b := ""

	for ; i < max; i += increment {
		if data[i] {
			b += "1"
		} else {
			b += "0"
		}
	}

	id, _ := strconv.ParseInt(b, 2, 64)

	return int(id)
}

func getTileVariants(input string) []*tileVariants {
	parts := utils.ToStringSlice(input, "\n\n")
	tiles := make([]*tileVariants, len(parts))

	for i, p := range parts {
		rows := utils.ToRuneSlice(p, "\n")
		height := len(rows[1:])
		width := len(rows[1])
		id, _ := strconv.Atoi(string(rows[0][5 : len(rows[0])-1]))
		tiles[i] = &tileVariants{
			id:       id,
			height:   height,
			width:    width,
			variants: make([]*tile, 0, 8),
		}

		data := make([]bool, height*width)

		for y, row := range rows[1:] {
			for x, state := range row {
				if state == '#' {
					data[width*y+x] = true
				}
			}
		}

		flipped := flip(data, width)

		for _, d := range [][]bool{
			data,
			rotate(data, width, 1),
			rotate(data, width, 2),
			rotate(data, width, 3),
			flipped,
			rotate(flipped, width, 1),
			rotate(flipped, width, 2),
			rotate(flipped, width, 3),
		} {
			sideIds := getSideIds(d, width, height)
			variantId := sideIds[0] + sideIds[1] + sideIds[2] + sideIds[3]
			tiles[i].variants = append(tiles[i].variants, &tile{
				id:      id,
				variant: variantId,
				height:  height,
				width:   width,
				data:    d,
				sides:   sideIds,
			})
		}
	}

	return tiles
}

func flip(data []bool, width int) []bool {
	newData := make([]bool, len(data))

	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			newData[width*(width-1-y)+x] = data[width*y+x]
		}
	}

	return newData
}

func rotate(data []bool, width int, times int) []bool {
	for i := 0; i < times; i++ {
		newData := make([]bool, len(data))

		for x := 0; x < width; x++ {
			for y := 0; y < width; y++ {
				newData[width*x+(width-1-y)] = data[width*y+x]
			}
		}

		data = newData
	}

	return data
}

func findCorners(variants []*tileVariants) []*tileVariants {
	c := make([]*tileVariants, 0, 4)
	t := 0

	for _, v := range variants {
		matches := 0

	SideLoop:
		for _, v2 := range variants {
			if v.id == v2.id {
				continue
			}

			for _, t := range v2.variants {
				for _, s := range v.variants[0].sides {
					for _, s2 := range t.sides {
						if s == s2 {
							matches++
							continue SideLoop
						}
					}
				}
			}
		}

		if matches == 2 {
			c = append(c, v)
			t++

			if t == 4 {
				break
			}
		}
	}

	return c
}

func buildTileSet(variants []*tileVariants, corners []*tileVariants) []*tile {
	max := len(variants)
	width := int(math.Sqrt(float64(max)))
	firstCorner := 0
	firstCornerVariant := 0

	var tileSet []*tile

MainLoop:
	for {
		tileSet = make([]*tile, max)

		for y := 0; y < width; y++ {
			for x := 0; x < width; x++ {
				offset := width*y + x

				if x == 0 && y == 0 {
					tileSet[offset] = corners[firstCorner].variants[firstCornerVariant]
					continue
				}

				var selection []*tileVariants

				if (x == width-1 && y == 0) || (x == 0 && y == width-1) || (x == width-1 && y == width-1) {
					selection = corners
				} else {
					selection = variants
				}

			SelectionLoop:
				for _, v := range selection {
					if offset > 0 && tileSet[offset-1].id == v.id {
						continue
					}

					if offset > width && tileSet[offset-width].id == v.id {
						continue
					}

					for _, t := range v.variants {
						if y != 0 && t.sides[0] != tileSet[offset-width].sides[2] {
							continue
						}

						if x != 0 && t.sides[3] != tileSet[offset-1].sides[1] {
							continue
						}

						tileSet[offset] = t
						break SelectionLoop
					}
				}

				if tileSet[offset] == nil {
					firstCornerVariant++

					if firstCornerVariant == len(corners[firstCorner].variants) {
						firstCorner++
						firstCornerVariant = 0
					}

					continue MainLoop
				}
			}
		}

		break MainLoop
	}

	return tileSet
}

func mergeTiles(tiles []*tile) []bool {
	tileLen := len(tiles)
	perRow := int(math.Sqrt(float64(tileLen)))
	width := tiles[0].width * perRow
	height := tiles[0].height * perRow
	data := make([]bool, width*height)
	start := 0

	for i := start; i < perRow; i++ {
		for y := 0; y < tiles[0].height; y++ {
			for j := start; j < perRow; j++ {
				for x := 0; x < tiles[0].width; x++ {
					index := i*perRow + j
					oldOffset := tiles[0].width*y + x
					newOffset := width*y + (x + (perRow * tiles[0].width * tiles[0].width * i) + (j * tiles[0].width))
					data[newOffset] = tiles[index].data[oldOffset]
				}
			}
		}
	}

	return data
}

func removeBorders(t *tile) {
	newData := make([]bool, len(t.data)-2*(t.height-2)-2*t.width)

	for y := 1; y < t.height-1; y++ {
		for x := 1; x < t.width-1; x++ {
			newData[(t.width-2)*(y-1)+(x-1)] = t.data[t.width*y+x]
		}
	}

	t.height -= 2
	t.width -= 2
	t.data = newData
}
