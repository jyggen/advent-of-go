package main

import (
	"container/ring"
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
	"strings"
)

const p1rounds = 100
const p2rounds = 10000000
const p2cups = 1000000

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	cups, _ := utils.ToIntegerSlice(input, "")
	max := utils.MaxIntSlice(cups)
	min := utils.MinIntSlice(cups)
	r := play(cups, min, max, p1rounds)

	var b strings.Builder

	for j := 1; j < r.Len(); j++ {
		r = r.Next()

		b.WriteString(fmt.Sprint(r.Value.(int)))
	}

	return b.String(), nil
}

func SolvePart2(input string) (string, error) {
	cups, _ := utils.ToIntegerSlice(input, "")
	max := utils.MaxIntSlice(cups)
	min := utils.MinIntSlice(cups)
	additional := make([]int, p2cups-max)

	for i := 0; i < len(additional); i++ {
		additional[i] = max + i + 1
	}

	max = p2cups
	cups = append(cups, additional...)
	r := play(cups, min, max, p2rounds).Next()
	val1 := r.Value.(int)
	val2 := r.Next().Value.(int)

	return strconv.Itoa(val1 * val2), nil
}

func play(cups []int, min int, max int, rounds int) *ring.Ring {
	rLen := len(cups)
	r := ring.New(rLen)
	lookup := map[int]*ring.Ring{}

	for _, c := range cups {
		lookup[c] = r
		r.Value = c
		r = r.Next()
	}

	for i := 0; i < rounds; i++ {
		removed := r.Unlink(3)
		destination := r.Value.(int) - 1

		for {
			if destination < min {
				destination = max
			}

			isRemoved := false

			removed.Do(func(p interface{}) {
				if p.(int) == destination {
					isRemoved = true
				}
			})

			if !isRemoved {
				lookup[destination].Link(removed)
				break
			}

			destination--
		}

		r = r.Next()
	}

	return lookup[1]
}
