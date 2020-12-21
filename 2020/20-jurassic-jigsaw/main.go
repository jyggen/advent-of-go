package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
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
	tiles := utils.ToStringSlice(input, "\n\n")
	tileIds := make(map[int][]int, len(tiles))

	for _, t := range tiles {
		rows := utils.ToRuneSlice(t, "\n")
		height := len(rows[1:])
		width := len(rows[1])
		id, _ := strconv.Atoi(string(rows[0][5:len(rows[0])-1]))
		tile := make([]bool, width * height)

		for y, row := range rows[1:] {
			for x, state := range row {
				if state == '#' {
					tile[width*y+x] = true
				}
			}
		}

		tileIds[id] = make([]int, 8)
		tileIds[id][0], tileIds[id][1] = getSideId(tile, 0, width, 1)
		tileIds[id][2], tileIds[id][3] = getSideId(tile, 0, height*width, width)
		tileIds[id][4], tileIds[id][5] = getSideId(tile, width-1, height*width, width)
		tileIds[id][6], tileIds[id][7] = getSideId(tile, (height-1)*width, height*width, 1)
	}

	sum := 1

	for id, sides := range tileIds {
		matches := 0

	SideLoop:
		for _, s := range sides {
			for id2, sides2 := range tileIds {
				if id == id2 {
					continue
				}

				for _, s2 := range sides2 {
					if s2 == s {
						matches++
						continue SideLoop
					}
				}
			}
		}

		if matches == 4 {
			fmt.Println(id)
			sum *= id
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(0), nil
}

func getSideId(tile []bool, i int, max int, increment int) (int, int) {
	b := ""
	b2 := ""

	for ; i < max; i += increment {
		if tile[i] {
			b += "1"
			b2 = "1" + b2
		} else {
			b += "0"
			b2 = "0" + b2
		}
	}

	id, _ := strconv.ParseInt(b, 2, 64)
	id2, _ := strconv.ParseInt(b2, 2, 64)

	return int(id), int(id2)
}
