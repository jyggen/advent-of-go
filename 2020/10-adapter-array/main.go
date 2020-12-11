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
	value    int
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
	gaps := []int{0, 0, 1}

	sort.Ints(adapters)

	for k, v := range adapters {
		if k != 0 {
			v -= adapters[k-1]
		}

		gaps[v-1]++
	}

	return strconv.Itoa(gaps[0] * gaps[2]), nil
}

func SolvePart2(input string) (string, error) {
	adapters, _ := utils.ToIntegerSlice(strings.TrimSpace(input)+"\n0", "\n")

	sort.Ints(adapters)

	adaptersLen := len(adapters)
	adapters = append(adapters, adapters[adaptersLen-1]+3)
	branches := make([]*branch, adaptersLen+1)

	for k, v := range adapters {
		b := &branch{
			branches: make([]*branch, 0),
			cache:    -1,
			value:    v,
		}

		for i := k - 1; i >= k-3 && i >= 0; i-- {
			if branches[i].value >= v-3 {
				branches[i].branches = append(branches[i].branches, b)
			}
		}

		branches[k] = b
	}

	return strconv.Itoa(getBranchCount(branches[0])), nil
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
