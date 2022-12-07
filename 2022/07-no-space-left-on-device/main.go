package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"sort"
	"strconv"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	directories []*directory
	files       []*file
	parent      *directory
	totalSize   int
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{
		name:        name,
		directories: make([]*directory, 0),
		files:       make([]*file, 0),
		parent:      parent,
	}
}

func buildTree(input string) (*directory, []*directory) {
	commands := utils.ToStringSlice(input[2:], "$ ")
	root := newDirectory("", nil)
	current := root
	directories := []*directory{root}

	for _, command := range commands {
		rows := utils.ToStringSlice(command, "\n")

		switch rows[0][0:2] {
		case "cd":
			name := rows[0][3:]
			switch name {
			case "/":
				current = root
			case "..":
				current = current.parent
			default:
				for _, d := range current.directories {
					if d.name == name {
						current = d
						break
					}
				}
			}
		case "ls":
			for _, row := range rows[1:] {
				parts := utils.ToStringSlice(row, " ")
				name := parts[1]

				if parts[0] == "dir" {
					dir := newDirectory(name, current)
					current.directories = append(current.directories, dir)
					directories = append(directories, dir)
				} else {
					size, _ := strconv.Atoi(parts[0])
					current.files = append(current.files, &file{name: name, size: size})

					for i := current; i != nil; i = i.parent {
						i.totalSize += size
					}
				}
			}
		}
	}

	sort.Slice(directories, func(i, j int) bool {
		return directories[i].totalSize < directories[j].totalSize
	})

	return root, directories
}

func SolvePart1(input string) (string, error) {
	_, directories := buildTree(input)
	highest := sort.Search(len(directories), func(i int) bool {
		return directories[i].totalSize > 100000
	})

	total := 0

	for i := highest - 1; i >= 0; i-- {
		total += directories[i].totalSize
	}

	return strconv.Itoa(total), nil
}

func SolvePart2(input string) (string, error) {
	root, directories := buildTree(input)
	left := 70000000 - root.totalSize
	needed := 30000000 - left
	highest := sort.Search(len(directories), func(i int) bool {
		return directories[i].totalSize > needed
	})

	return strconv.Itoa(directories[highest].totalSize), nil
}
