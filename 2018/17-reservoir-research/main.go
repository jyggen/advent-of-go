package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"math"
	"os"
	"regexp"
	"strconv"
)

type clayTile struct {
	y int
	x int
}

type waterTile struct {
	endless  bool
	children []*waterTile
	parent   *waterTile
	y        int
	x        int
}

func newWaterTile(y int, x int) *waterTile {
	return &waterTile{
		children: make([]*waterTile, 0),
		y:        y,
		x:        x,
	}
}

func (w *waterTile) addChild(child *waterTile) {
	w.children = append(w.children, child)
	child.parent = w
}

const sand = 0
const clay = 1
const water = 2
const spring = 3
const endless = 4

var inputRegex = regexp.MustCompile("^([xy])=(\\d+), [xy]=(\\d+)..(\\d+)$")

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	data := utils.ToStringSlice(input, "\n")
	clayTiles := make([]*clayTile, 0)
	minX := math.MaxInt16
	maxX := 0
	minY := math.MaxInt16
	maxY := 0

	for _, line := range data {
		match := inputRegex.FindStringSubmatch(line)

		if len(match) == 0 {
			panic("unable to find matches")
		}

		var x1, x2, y1, y2 int

		two, _ := strconv.Atoi(match[2])
		three, _ := strconv.Atoi(match[3])
		four, _ := strconv.Atoi(match[4])

		if match[1] == "x" {
			x1, x2 = two, two
			y1, y2 = three, four
		} else {
			y1, y2 = two, two
			x1, x2 = three, four
		}

		for x := x1; x <= x2; x++ {
			if x >= maxX {
				maxX = x + 1
			}

			if x <= minX {
				minX = x - 1
			}

			for y := y1; y <= y2; y++ {
				if y >= maxY {
					maxY = y
				}

				if y <= minY {
					minY = y - 1
				}

				clayTiles = append(clayTiles, &clayTile{x: x, y: y})
			}
		}
	}

	layout := make([][]int, maxY-minY+1)

	for y := 0; y < len(layout); y++ {
		layout[y] = make([]int, maxX-minX+1)
	}

	for _, c := range clayTiles {
		layout[c.y-minY][c.x-minX] = clay
	}

	current := newWaterTile(0, 500-minX)
	maxY, minY, maxX, minX = maxY-minY, 0, maxX-minX, 0
	layout[current.y][current.x] = spring
	waterTiles := make(map[string]*waterTile, 0)
	waterTiles[fmt.Sprintf("%d-%d", current.y, current.x)] = current
	numOfWater := 0
	atRest := 0

	for {
		if current.y+1 > maxY {
			layout[current.y][current.x] = endless
			current.endless = true
			numOfWater++
			current.parent.children = current.parent.children[1:]
			current.parent.endless = true
			current = current.parent
		} else if !current.endless && current.y+1 <= maxY && layout[current.y+1][current.x] == sand {
			child := newWaterTile(current.y+1, current.x)
			waterTiles[fmt.Sprintf("%d-%d", child.y, child.x)] = child
			current.addChild(child)
			layout[child.y][child.x] = spring
			current = child

			if current.y+1 <= maxY {
				if below, ok := waterTiles[fmt.Sprintf("%d-%d", current.y+1, current.x)]; ok && below.endless {
					current.endless = true
				}
			}
		} else if !current.endless && current.x+1 <= maxX && layout[current.y][current.x+1] == sand {
			child := newWaterTile(current.y, current.x+1)
			waterTiles[fmt.Sprintf("%d-%d", child.y, child.x)] = child
			current.addChild(child)
			layout[child.y][child.x] = spring
		} else if !current.endless && current.x-1 >= minX && layout[current.y][current.x-1] == sand {
			child := newWaterTile(current.y, current.x-1)
			waterTiles[fmt.Sprintf("%d-%d", child.y, child.x)] = child
			current.addChild(child)
			layout[child.y][child.x] = spring
		} else if len(current.children) > 0 {
			if current.parent != nil && current.parent.endless {
				current.endless = true
			}

			current = current.children[0]
		} else if current.parent != nil {
			numOfWater++
			current.parent.children = current.parent.children[1:]

			if current.endless {
				layout[current.y][current.x] = endless
				current.parent.endless = true
			} else if current.parent.endless {
				layout[current.y][current.x] = endless
				current.endless = true
			} else {
				layout[current.y][current.x] = water
				atRest++
			}

			current = current.parent
		} else {
			break
		}
	}

	return strconv.Itoa(numOfWater), nil
}

func SolvePart2(input string) (string, error) {
	data := utils.ToStringSlice(input, "\n")
	clayTiles := make([]*clayTile, 0)
	minX := math.MaxInt16
	maxX := 0
	minY := math.MaxInt16
	maxY := 0

	for _, line := range data {
		match := inputRegex.FindStringSubmatch(line)

		if len(match) == 0 {
			panic("unable to find matches")
		}

		var x1, x2, y1, y2 int

		two, _ := strconv.Atoi(match[2])
		three, _ := strconv.Atoi(match[3])
		four, _ := strconv.Atoi(match[4])

		if match[1] == "x" {
			x1, x2 = two, two
			y1, y2 = three, four
		} else {
			y1, y2 = two, two
			x1, x2 = three, four
		}

		for x := x1; x <= x2; x++ {
			if x >= maxX {
				maxX = x + 1
			}

			if x <= minX {
				minX = x - 1
			}

			for y := y1; y <= y2; y++ {
				if y >= maxY {
					maxY = y
				}

				if y <= minY {
					minY = y - 1
				}

				clayTiles = append(clayTiles, &clayTile{x: x, y: y})
			}
		}
	}

	layout := make([][]int, maxY-minY+1)

	for y := 0; y < len(layout); y++ {
		layout[y] = make([]int, maxX-minX+1)
	}

	for _, c := range clayTiles {
		layout[c.y-minY][c.x-minX] = clay
	}

	current := newWaterTile(0, 500-minX)
	maxY, minY, maxX, minX = maxY-minY, 0, maxX-minX, 0
	layout[current.y][current.x] = spring
	waterTiles := make(map[string]*waterTile, 0)
	waterTiles[fmt.Sprintf("%d-%d", current.y, current.x)] = current
	numOfWater := 0
	atRest := 0

	for {
		if current.y+1 > maxY {
			layout[current.y][current.x] = endless
			current.endless = true
			numOfWater++
			current.parent.children = current.parent.children[1:]
			current.parent.endless = true
			current = current.parent
		} else if !current.endless && current.y+1 <= maxY && layout[current.y+1][current.x] == sand {
			child := newWaterTile(current.y+1, current.x)
			waterTiles[fmt.Sprintf("%d-%d", child.y, child.x)] = child
			current.addChild(child)
			layout[child.y][child.x] = spring
			current = child

			if current.y+1 <= maxY {
				if below, ok := waterTiles[fmt.Sprintf("%d-%d", current.y+1, current.x)]; ok && below.endless {
					current.endless = true
				}
			}
		} else if !current.endless && current.x+1 <= maxX && layout[current.y][current.x+1] == sand {
			child := newWaterTile(current.y, current.x+1)
			waterTiles[fmt.Sprintf("%d-%d", child.y, child.x)] = child
			current.addChild(child)
			layout[child.y][child.x] = spring
		} else if !current.endless && current.x-1 >= minX && layout[current.y][current.x-1] == sand {
			child := newWaterTile(current.y, current.x-1)
			waterTiles[fmt.Sprintf("%d-%d", child.y, child.x)] = child
			current.addChild(child)
			layout[child.y][child.x] = spring
		} else if len(current.children) > 0 {
			if current.parent != nil && current.parent.endless {
				current.endless = true
			}

			current = current.children[0]
		} else if current.parent != nil {
			numOfWater++
			current.parent.children = current.parent.children[1:]

			if current.endless {
				layout[current.y][current.x] = endless
				current.parent.endless = true
			} else if current.parent.endless {
				layout[current.y][current.x] = endless
				current.endless = true
			} else {
				layout[current.y][current.x] = water
				atRest++
			}

			current = current.parent
		} else {
			break
		}
	}

	return strconv.Itoa(atRest), nil
}

func printLayout(layout [][]int) {
	for _, row := range layout {
		for _, tile := range row {
			switch tile {
			case sand:
				fmt.Print(".")
			case clay:
				fmt.Print("#")
			case water:
				fmt.Print("~")
			case spring:
				fmt.Print("+")
			case endless:
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
