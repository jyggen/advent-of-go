package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
	"os"
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

func createGraph(input string, multiplier int) *dijkstra.Graph {
	rows := utils.ToStringSlice(input, "\n")
	colLength := len(rows[0])
	rowLength := len(rows)
	data := make([]int64, rowLength*multiplier*colLength*multiplier)
	graph := dijkstra.NewGraph()
	vertexId, aboveId, leftId := 0, 0-(colLength*multiplier), -1

	for j := 0; j < multiplier; j++ {
		for i, r := range rows {
			integers, _ := utils.ToIntegerSlice(r, "")

			for k := 0; k < multiplier; k++ {
				for l, v := range integers {
					value := int64(v + j + k)

					if value > 9 {
						value -= 9
					}

					data[vertexId] = value

					graph.AddVertex(vertexId)

					if i != 0 || (i == 0 && j != 0) {
						graph.AddArc(vertexId, aboveId, data[aboveId])
						graph.AddArc(aboveId, vertexId, value)
					}

					if l != 0 || (l == 0 && k != 0) {
						graph.AddArc(vertexId, leftId, data[leftId])
						graph.AddArc(leftId, vertexId, value)
					}

					vertexId, aboveId, leftId = vertexId+1, aboveId+1, leftId+1
				}
			}
		}
	}

	return graph
}

func SolvePart1(input string) (string, error) {
	graph := createGraph(input, 1)
	best, err := graph.Shortest(0, len(graph.Verticies)-1)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(best.Distance)), nil
}

func SolvePart2(input string) (string, error) {
	graph := createGraph(input, 5)
	best, err := graph.Shortest(0, len(graph.Verticies)-1)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(best.Distance)), nil
}
