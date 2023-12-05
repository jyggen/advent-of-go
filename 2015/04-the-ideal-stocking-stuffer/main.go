package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

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

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(solve([]byte(strings.TrimSpace(input)), 5)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(solve([]byte(strings.TrimSpace(input)), 6)), nil
}

func solve(key []byte, zeroes int) int {
	iteration := int64(0)
	numBytes := int(math.Floor(float64(zeroes) / 2))
HashLoop:
	for {
		iteration++
		hash := md5.Sum(strconv.AppendInt(key, iteration, 10))

		if zeroes%2 != 0 && hash[numBytes] > 15 {
			continue HashLoop
		}

		for i := 0; i < numBytes; i++ {
			if hash[i] != 0 {
				continue HashLoop
			}
		}

		return int(iteration)
	}
}
