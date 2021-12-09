package main

import (
	"fmt"
	solver2 "github.com/jyggen/advent-of-go/internal/solver"
	utils2 "github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	parents  []*Bag
	children []*Bag
}

type Match struct {
	parent   *Bag
	children [][]string
}

const ourBagColor = "shiny gold"

var parentBagRegex *regexp.Regexp
var childBagRegex *regexp.Regexp

func init() {
	parentBagRegex = regexp.MustCompile(`(.+?) bags`)
	childBagRegex = regexp.MustCompile(`(\d+) (.+?) bags?`)
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
	ourBag, err := parseBags(input)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(parentCount(ourBag, make(map[string]bool, 0))), nil
}

func SolvePart2(input string) (string, error) {
	ourBag, err := parseBags(input)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(childrenCount(ourBag, make(map[string]int, 0))), nil
}

func bagByColor(bags map[string]*Bag, color string) *Bag {
	if _, ok := bags[color]; !ok {
		bags[color] = &Bag{
			color:    color,
			children: make([]*Bag, 0),
		}
	}

	return bags[color]
}

func childrenCount(bag *Bag, colors map[string]int) (count int) {
	for _, c := range bag.children {
		if _, ok := colors[c.color]; !ok {
			colors[c.color] = childrenCount(c, colors) + 1
		}

		count += colors[c.color]
	}

	return count
}

func parentCount(bag *Bag, colors map[string]bool) (count int) {
	for _, p := range bag.parents {
		if _, ok := colors[p.color]; !ok {
			colors[p.color] = true
			count++
		}

		count += parentCount(p, colors)
	}

	return count
}

func parseBags(input string) (*Bag, error) {
	rules := utils2.ToStringSlice(input, "\n")
	rulesLen := len(rules)
	bags := make(map[string]*Bag, rulesLen)
	matches := make([]*Match, rulesLen)

	for i, r := range rules {
		parts := strings.SplitN(r, " contain ", 2)
		parentMatch := parentBagRegex.FindStringSubmatch(parts[0])
		bags[parentMatch[1]] = &Bag{
			color:    parentMatch[1],
			children: make([]*Bag, 0),
		}
		matches[i] = &Match{
			parent:   bags[parentMatch[1]],
			children: childBagRegex.FindAllStringSubmatch(parts[1], -1),
		}
	}

	for _, m := range matches {
		for _, c := range m.children {
			childBag := bags[c[2]]
			childBag.parents = append(childBag.parents, m.parent)
			childCount, err := strconv.Atoi(c[1])

			if err != nil {
				return nil, err
			}

			for i := 0; i < childCount; i++ {
				m.parent.children = append(m.parent.children, childBag)
			}
		}
	}

	return bags[ourBagColor], nil
}
