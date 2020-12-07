package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"regexp"
	"strconv"
)

type Bag struct {
	color    string
	parents  []*Bag
	children []*Bag
}

var parentBagRegex *regexp.Regexp
var childBagRegex *regexp.Regexp

func init() {
	parentBagRegex = regexp.MustCompile(`^([a-z ]+) bags contain`)
	childBagRegex = regexp.MustCompile(`([\d]+) ([a-z ]+) bags?`)
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
	rules := utils.ToStringSlice(input, "\n")
	bags, err := parseBags(rules)

	if err != nil {
		return "", err
	}

	shinyBag := bagByColor(bags, "shiny gold")

	return strconv.Itoa(parentCount(shinyBag, make(map[string]bool, 0))), nil
}

func SolvePart2(input string) (string, error) {
	rules := utils.ToStringSlice(input, "\n")
	bags, err := parseBags(rules)

	if err != nil {
		return "", err
	}

	shinyBag := bagByColor(bags, "shiny gold")

	return strconv.Itoa(childrenCount(shinyBag, make(map[string]int, 0))), nil
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

func parseBags(rules []string) (map[string]*Bag, error) {
	bags := make(map[string]*Bag, 0)

	for _, r := range rules {
		parentMatch := parentBagRegex.FindStringSubmatch(r)
		childMatches := childBagRegex.FindAllStringSubmatch(r, -1)
		parentBag := bagByColor(bags, parentMatch[1])

		for _, child := range childMatches {
			childBag := bagByColor(bags, child[2])
			childBag.parents = append(childBag.parents, parentBag)
			childCount, err := strconv.Atoi(child[1])

			if err != nil {
				return bags, err
			}

			for i := 0; i < childCount; i++ {
				parentBag.children = append(parentBag.children, childBag)
			}
		}
	}

	return bags, nil
}
