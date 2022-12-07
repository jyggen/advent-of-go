package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
	"path/filepath"
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
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{
		name:        name,
		directories: make([]*directory, 0),
		files:       make([]*file, 0),
		parent:      parent,
	}
}

func buildTree(input string) *directory {
	commands := utils.ToStringSlice(input[2:], "$ ")
	root := newDirectory("", nil)
	current := root

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
					current.directories = append(current.directories, newDirectory(name, current))
				} else {
					size, _ := strconv.Atoi(parts[0])
					current.files = append(current.files, &file{name: name, size: size})
				}
			}
		}
	}

	return root
}

func calculateSizes(dir *directory, result map[string]int, path string) int {
	path = filepath.Join(path, dir.name)
	size := 0

	for _, f := range dir.files {
		size += f.size
	}

	for _, d := range dir.directories {
		size += calculateSizes(d, result, path)
	}

	result[path] = size

	return size
}

func SolvePart1(input string) (string, error) {
	root := buildTree(input)
	sizes := make(map[string]int)

	calculateSizes(root, sizes, "/")

	total := 0

	for _, size := range sizes {
		if size <= 100000 {
			total += size
		}
	}

	return strconv.Itoa(total), nil
}

func SolvePart2(input string) (string, error) {
	root := buildTree(input)
	sizes := make(map[string]int)

	calculateSizes(root, sizes, "/")

	left := 70000000 - sizes["/"]
	needed := 30000000 - left
	bestDiff := sizes["/"]
	bestSize := 0

	for _, size := range sizes {
		diff := size - needed
		if diff >= 0 && diff < bestDiff {
			bestDiff = diff
			bestSize = size
		}
	}

	return strconv.Itoa(bestSize), nil
}
