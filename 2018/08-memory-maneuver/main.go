package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"strconv"
)

type node struct {
	children []*node
	metadata []int
}

func main() {
	p1, p2, err := solver2.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	integers, err := utils2.ToIntegerSlice(input, " ")

	if err != nil {
		return "", err
	}

	tree, _ := buildTree(integers)

	return strconv.Itoa(tree.Sum()), nil
}

func SolvePart2(input string) (string, error) {
	integers, err := utils2.ToIntegerSlice(input, " ")

	if err != nil {
		return "", err
	}

	tree, _ := buildTree(integers)

	return strconv.Itoa(tree.Value()), nil
}

func buildTree(integers []int) (*node, []int) {
	var childrenLen int
	var metadataLen int

	childrenLen, integers = integers[0], integers[1:]
	metadataLen, integers = integers[0], integers[1:]
	n := &node{
		children: make([]*node, childrenLen),
		metadata: make([]int, metadataLen),
	}

	for i := 0; i < childrenLen; i++ {
		n.children[i], integers = buildTree(integers)
	}

	for i := 0; i < metadataLen; i++ {
		n.metadata[i], integers = integers[0], integers[1:]
	}

	return n, integers
}

func (n *node) Sum() int {
	sum := 0

	for _, m := range n.metadata {
		sum += m
	}

	for _, c := range n.children {
		sum += c.Sum()
	}

	return sum
}

func (n *node) Value() int {
	value := 0
	childrenLen := len(n.children)

	if childrenLen == 0 {
		for _, v := range n.metadata {
			value += v
		}

		return value
	}

	for _, i := range n.metadata {
		i--

		if i >= childrenLen {
			continue
		}

		value += n.children[i].Value()
	}

	return value
}
