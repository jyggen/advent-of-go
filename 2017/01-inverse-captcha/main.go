package main

import (
	"container/ring"
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
	r, err := makeRing(input)

	if err != nil {
		return "", err
	}

	rLen := r.Len()
	sum := 0

	for i := 0; i < rLen; i++ {
		if r.Value == r.Next().Value {
			sum += r.Value.(int)
		}

		r = r.Next()
	}

	return strconv.Itoa(sum), err
}

func SolvePart2(input string) (string, error) {
	r, err := makeRing(input)

	if err != nil {
		return "", err
	}

	rLen := r.Len()
	steps := rLen / 2
	sum := 0

	for i := 0; i < rLen; i++ {
		if r.Value == r.Move(steps).Value {
			sum += r.Value.(int)
		}

		r = r.Next()
	}

	return strconv.Itoa(sum), err
}

func makeRing(input string) (*ring.Ring, error) {
	intSlice, err := utils2.ToIntegerSlice(input, "")

	if err != nil {
		return nil, err
	}

	r := ring.New(len(intSlice))

	for _, i := range intSlice {
		r.Value = i
		r = r.Next()
	}

	return r, nil
}
