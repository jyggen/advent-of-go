package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"sort"
	"strconv"
	"strings"
)

type branch struct {
	branches []*branch
	cache    int
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
	adapters, _ := utils.ToIntegerSlice(input, "\n")
	gaps := map[int]int{
		1: 0,
		2: 0,
		3: 1,
	}

	sort.Ints(adapters)

	for k, v := range adapters {

		if k == 0 {
			gaps[v]++
		} else {
			gaps[v-adapters[k-1]]++
		}
	}

	return strconv.Itoa(gaps[1] * gaps[3]), nil
}

func SolvePart2(input string) (string, error) {
	adapters, _ := utils.ToIntegerSlice(strings.TrimSpace(input)+"\n0", "\n")

	sort.Ints(adapters)

	adaptersLen := len(adapters)
	adapters = append(adapters, adapters[adaptersLen-1]+3)
	numbers := make(map[int]*branch, adaptersLen+1)

	for _, x := range adapters {
		numbers[x] = &branch{
			branches: make([]*branch, 0),
			cache:    -1,
		}
	}

	for i := 0; i <= adaptersLen; i++ {
		lookahead := i + 1

		for {
			if lookahead > adaptersLen || adapters[lookahead]-adapters[i] > 3 {
				break
			}

			numbers[adapters[i]].branches = append(numbers[adapters[i]].branches, numbers[adapters[lookahead]])

			lookahead++
		}
	}

	return strconv.Itoa(getBranchCount(numbers[0])), nil
}

func getBranchCount(n *branch) int {
	if n.cache != -1 {
		return n.cache
	}

	branches := 0

	if len(n.branches) == 0 {
		branches += 1
	} else {
		for _, f := range n.branches {
			branches += getBranchCount(f)
		}
	}

	n.cache = branches

	return branches
}
