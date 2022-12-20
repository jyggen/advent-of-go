package main

import (
	"container/ring"
	"fmt"
	"github.com/jyggen/advent-of-go/internal/utils"
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

func solve(input string, mixCount int, decryptionKey int) int {
	numbers := utils.ToOptimisticIntSlice(input, true)

	for i, n := range numbers {
		numbers[i] = n * decryptionKey
	}

	r := ring.New(len(numbers))
	zeroIndex := 0

	for i, n := range numbers {
		if n == 0 {
			zeroIndex = i
		}

		r.Value = i
		r = r.Next()
	}

	for i := 0; i < mixCount; i++ {
		for j, n := range numbers {
			for {
				if r.Value == j {
					break
				}

				r = r.Next()
			}

			r = r.Prev()
			value := r.Unlink(1)
			r = r.Move(n % r.Len())
			r.Link(value)
		}
	}

	for {
		if r.Value == zeroIndex {
			break
		}

		r = r.Next()
	}

	return numbers[r.Move(1000).Value.(int)] + numbers[r.Move(2000).Value.(int)] + numbers[r.Move(3000).Value.(int)]
}

func SolvePart1(input string) (string, error) {
	return strconv.Itoa(solve(input, 1, 1)), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(solve(input, 10, 811589153)), nil
}
