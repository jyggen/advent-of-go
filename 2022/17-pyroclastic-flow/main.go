package main

import (
	"container/ring"
	"fmt"
	"github.com/jyggen/advent-of-go/internal/utils"
	"math"
	"os"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type shapeRule struct {
	lowerBoundX int
	upperBoundX int
	lowerBoundY int
	upperBoundY int
}

type shapeRules []shapeRule

func (sr shapeRules) copy() shapeRules {
	newRules := make(shapeRules, len(sr))

	for i, r := range sr {
		newRules[i] = shapeRule{
			lowerBoundX: r.lowerBoundX,
			upperBoundX: r.upperBoundX,
			lowerBoundY: r.lowerBoundY,
			upperBoundY: r.upperBoundY,
		}
	}

	return newRules
}

func (sr shapeRules) moveDown() {
	for i := range sr {
		sr[i].lowerBoundY++
		sr[i].upperBoundY++
	}
}

func (sr shapeRules) moveLeft() {
	for i := range sr {
		sr[i].lowerBoundX--
		sr[i].upperBoundX--
	}
}

func (sr shapeRules) moveRight() {
	for i := range sr {
		sr[i].lowerBoundX++
		sr[i].upperBoundX++
	}
}

var shapes = [5]shapeRules{
	{ // _
		{
			2,
			5,
			3,
			3,
		},
	},
	{ // +
		{
			3,
			3,
			1,
			1,
		},
		{
			2,
			4,
			2,
			2,
		},
		{
			3,
			3,
			3,
			3,
		},
	},
	{ // ⅃
		{
			4,
			4,
			1,
			2,
		},
		{
			2,
			4,
			3,
			3,
		},
	},
	{ // |
		{
			2,
			2,
			0,
			3,
		},
	},
	{ // ■
		{
			2,
			3,
			2,
			3,
		},
	},
}

var allEmpty = [7]bool{false, false, false, false, false, false, false}

const (
	LEFT = iota
	RIGHT
)

func SolvePart1(input string) (string, error) {
	directions := utils.ToStringSlice(input, "")
	rDirections := ring.New(len(directions))

	for _, direction := range directions {
		if direction == "<" {
			rDirections.Value = LEFT
		} else {
			rDirections.Value = RIGHT
		}

		rDirections = rDirections.Next()
	}

	rShapes := ring.New(len(shapes))

	for _, shape := range shapes {
		rShapes.Value = shape
		rShapes = rShapes.Next()
	}

	shavedOff := -1
	chamber := [][7]bool{
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{true, true, true, true, true, true, true},
	}

	for i := 0; i < 2022; i++ {
		shape := rShapes.Value.(shapeRules).copy()

		for {
			movedShape := shape.copy()

			switch rDirections.Value.(int) {
			case LEFT:
				movedShape.moveLeft()
			case RIGHT:
				movedShape.moveRight()
			}

			rDirections = rDirections.Next()
			canMove := true

		CanMoveSidewaysLoop:
			for _, r := range movedShape {
				if r.upperBoundX == len(chamber[0]) || r.lowerBoundX == -1 {
					canMove = false
					break CanMoveSidewaysLoop
				}

				for y := r.lowerBoundY; y <= r.upperBoundY; y++ {
					for x := r.lowerBoundX; x <= r.upperBoundX; x++ {
						if chamber[y][x] {
							canMove = false
							break CanMoveSidewaysLoop
						}
					}
				}
			}

			if canMove {
				shape = movedShape
			}

			movedShape = shape.copy()
			movedShape.moveDown()
			canMove = true

		CanMoveDownLoop:
			for _, r := range movedShape {
				for y := r.lowerBoundY; y <= r.upperBoundY; y++ {
					for x := r.lowerBoundX; x <= r.upperBoundX; x++ {
						if chamber[y][x] {
							canMove = false
							break CanMoveDownLoop
						}
					}
				}
			}

			if !canMove {
				for _, r := range shape {
					for y := r.lowerBoundY; y <= r.upperBoundY; y++ {
						for x := r.lowerBoundX; x <= r.upperBoundX; x++ {
							chamber[y][x] = true
						}
					}
				}
				break
			} else {
				shape = movedShape
			}
		}

		rShapes = rShapes.Next()
		allFound := [7]bool{true, true, true, true, true, true, true}
		hasFound := [7]bool{false, false, false, false, false, false, false}
		offsets := [7]int{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}
		highestNeeded := 0

		for y := 0; y < len(chamber); y++ {
			for x := 0; x < 7; x++ {
				if chamber[y][x] {
					hasFound[x] = true
					offsets[x] = utils.MinInt(offsets[x], y)
				}
			}

			if hasFound == allFound {
				break
			}
		}

		for _, o := range offsets {
			highestNeeded = utils.MaxInt(highestNeeded, o)
		}

		shavedOff += len(chamber) - 1 - highestNeeded
		chamber = chamber[:highestNeeded+1]

		for y := 0; y < len(chamber); y++ {
			if chamber[y] != allEmpty {
				for j := y; j < 7; j++ {
					chamber = append([][7]bool{{false, false, false, false, false, false, false}}, chamber...)
				}
				break
			}
		}
	}

	shavedOff += len(chamber)

	for y := 0; y < len(chamber); y++ {
		if chamber[y] == allEmpty {
			shavedOff--
		}
	}

	return strconv.Itoa(shavedOff), nil
}

func SolvePart2(input string) (string, error) {
	directions := utils.ToStringSlice(input, "")
	rDirections := ring.New(len(directions))

	for _, direction := range directions {
		if direction == "<" {
			rDirections.Value = LEFT
		} else {
			rDirections.Value = RIGHT
		}

		rDirections = rDirections.Next()
	}

	rShapes := ring.New(len(shapes))

	for _, shape := range shapes {
		rShapes.Value = shape
		rShapes = rShapes.Next()
	}

	shavedOff := -1
	chamber := [][7]bool{
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{true, true, true, true, true, true, true},
	}

	cache := make(map[string][2]int)
	shouldCache := true
	sequentialCacheHits := 0

	for i := 0; i < 1000000000000; i++ {
		shape := rShapes.Value.(shapeRules).copy()

		for {
			movedShape := shape.copy()

			switch rDirections.Value.(int) {
			case LEFT:
				movedShape.moveLeft()
			case RIGHT:
				movedShape.moveRight()
			}

			rDirections = rDirections.Next()
			canMove := true

		CanMoveSidewaysLoop:
			for _, r := range movedShape {
				if r.upperBoundX == len(chamber[0]) || r.lowerBoundX == -1 {
					canMove = false
					break CanMoveSidewaysLoop
				}

				for y := r.lowerBoundY; y <= r.upperBoundY; y++ {
					for x := r.lowerBoundX; x <= r.upperBoundX; x++ {
						if chamber[y][x] {
							canMove = false
							break CanMoveSidewaysLoop
						}
					}
				}
			}

			if canMove {
				shape = movedShape
			}

			movedShape = shape.copy()
			movedShape.moveDown()
			canMove = true

		CanMoveDownLoop:
			for _, r := range movedShape {
				for y := r.lowerBoundY; y <= r.upperBoundY; y++ {
					for x := r.lowerBoundX; x <= r.upperBoundX; x++ {
						if chamber[y][x] {
							canMove = false
							break CanMoveDownLoop
						}
					}
				}
			}

			if !canMove {
				for _, r := range shape {
					for y := r.lowerBoundY; y <= r.upperBoundY; y++ {
						for x := r.lowerBoundX; x <= r.upperBoundX; x++ {
							chamber[y][x] = true
						}
					}
				}
				break
			} else {
				shape = movedShape
			}
		}

		rShapes = rShapes.Next()
		allFound := [7]bool{true, true, true, true, true, true, true}
		hasFound := [7]bool{false, false, false, false, false, false, false}
		offsets := [7]int{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}
		highestNeeded := 0

		for y := 0; y < len(chamber); y++ {
			for x := 0; x < 7; x++ {
				if chamber[y][x] {
					hasFound[x] = true
					offsets[x] = utils.MinInt(offsets[x], y)
				}
			}

			if hasFound == allFound {
				break
			}
		}

		for _, o := range offsets {
			highestNeeded = utils.MaxInt(highestNeeded, o)
		}

		shavedOff += len(chamber) - 1 - highestNeeded
		chamber = chamber[:highestNeeded+1]

		for y := 0; y < len(chamber); y++ {
			if chamber[y] != allEmpty {
				for j := y; j < 7; j++ {
					chamber = append([][7]bool{{false, false, false, false, false, false, false}}, chamber...)
				}
				break
			}
		}

		if shouldCache {
			if z, ok := cache[fmt.Sprint(chamber, rShapes.Value, rDirections.Value)]; ok {
				heightDiff := shavedOff - z[1]
				roundDiff := i - z[0]

				if sequentialCacheHits < len(shapes) {
					sequentialCacheHits++
					continue
				}

				roundsToAdd := ((1000000000000 - roundDiff) / roundDiff) - 1
				i += roundsToAdd * roundDiff
				shavedOff += roundsToAdd * heightDiff
				shouldCache = false
			} else {
				sequentialCacheHits = 0
			}

			cache[fmt.Sprint(chamber, rShapes.Value, rDirections.Value)] = [2]int{i, shavedOff}
		}
	}

	shavedOff += len(chamber)

	for y := 0; y < len(chamber); y++ {
		if chamber[y] == allEmpty {
			shavedOff--
		}
	}

	return strconv.Itoa(shavedOff), nil
}

func calcHeight(shavedOff int, chamber [][7]bool) int {
	shavedOff += len(chamber)

	for y := 0; y < len(chamber); y++ {
		if chamber[y] == allEmpty {
			shavedOff--
		}
	}

	return shavedOff
}
