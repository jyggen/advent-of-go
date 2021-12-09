package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"math"
	"math/big"
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
	parts := utils2.ToStringSlice(input, "\n")
	earliest, _ := strconv.Atoi(parts[0])
	lines := utils2.ToStringSlice(parts[1], ",")
	bestWait := math.MaxInt16
	bestId := 0

	for _, l := range lines {
		if l == "x" {
			continue
		}

		num, _ := strconv.Atoi(l)
		next := num - (earliest % num)

		if next < bestWait {
			bestWait = next
			bestId = num
		}
	}

	return strconv.Itoa(bestWait * bestId), nil
}

func SolvePart2(input string) (string, error) {
	parts := utils2.ToStringSlice(input, "\n")
	busses := utils2.ToStringSlice(parts[1], ",")
	a := make([]*big.Int, 0)
	n := make([]*big.Int, 0)

	for k, v := range busses {
		if v == "x" {
			continue
		}

		line, _ := strconv.Atoi(v)
		n = append(n, big.NewInt(int64(line)))
		a = append(a, big.NewInt(int64(line-k)))
	}

	result, err := utils2.Crt(a, n)

	if err != nil {
		return "", err
	}

	return result.String(), nil
}
