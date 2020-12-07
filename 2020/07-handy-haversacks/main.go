package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
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

var parentBagRegex *regexp.Regexp
var childBagRegex *regexp.Regexp

func init() {
	parentBagRegex = regexp.MustCompile(`(.+?) bags`)
	childBagRegex = regexp.MustCompile(`(\d+) (.+?) bags?`)
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
	bags, err := parseBags(input)

	if err != nil {
		return "", err
	}

	shinyBag := bagByColor(bags, "shiny gold")

	return strconv.Itoa(parentCount(shinyBag, make(map[string]bool, 0))), nil
}

func SolvePart2(input string) (string, error) {
	bags, err := parseBags(input)

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

func parseBags(input string) (map[string]*Bag, error) {
	rules := utils.ToStringSlice(input, "\n")
	bags := make(map[string]*Bag, len(rules))

	for _, r := range rules {
		parts := strings.SplitN(r, " contain ", 2)
		parentMatch := parentBagRegex.FindStringSubmatch(parts[0])
		childMatches := childBagRegex.FindAllStringSubmatch(parts[1], -1)
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
