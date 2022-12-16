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

type valve struct {
	index int
	id    string
	rate  int
	open  bool
}

type item struct {
	id        int
	at        int
	minute    int
	total     int
	rate      int
	remaining []int
}

func SolvePart1(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	valves := make([]*valve, len(rows))
	graph := dijkstra.NewGraph()
	at := 0
	lookup := make(map[string]int, len(valves))

	for i := range valves {
		graph.AddVertex(i)
		lookup[rows[i][6:8]] = i
		valves[i] = &valve{}
	}

	for i, r := range rows {
		valveId := r[6:8]
		j := 23
		rateStr := make([]uint8, 0, 3)
		for ; j < len(r); j++ {
			if r[j] == ';' {
				break
			}

			rateStr = append(rateStr, r[j])
		}

		rate, _ := strconv.Atoi(string(rateStr))
		valves[i].index = i
		valves[i].id = valveId
		valves[i].rate = rate

		if valveId == "AA" {
			at = i
		}

		for _, v := range utils.ToStringSlice(r[j+24:], ", ") {
			graph.AddArc(i, lookup[v], 1)
		}
	}

	total := 0
	remaining := make([]int, 0, len(valves)-1)

	for _, v := range valves {
		if v.index != at && v.rate != 0 {
			remaining = append(remaining, v.index)
		}
	}

	queue := []item{
		{
			at:        at,
			minute:    1,
			total:     0,
			rate:      0,
			remaining: remaining,
		},
	}

	var q item

	maxRounds := 31

	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]

		if len(q.remaining) == 0 {
			q.total += q.rate * (31 - q.minute)
			q.minute = maxRounds
		}

		if q.minute == maxRounds {
			total = utils.MaxInt(total, q.total)
			continue
		}

		for i, r := range q.remaining {
			path, _ := graph.Shortest(q.at, r)
			newRemaining := make([]int, len(q.remaining))
			copy(newRemaining, q.remaining)
			newRemaining = append(newRemaining[:i], newRemaining[i+1:]...)
			queue = append(queue, item{
				at:        r,
				minute:    q.minute + utils.MinInt(int(path.Distance)+1, maxRounds-q.minute),
				total:     q.total + (q.rate * utils.MinInt(int(path.Distance)+1, maxRounds-q.minute)),
				rate:      q.rate + valves[r].rate,
				remaining: newRemaining,
			})
		}
	}

	return strconv.Itoa(total), nil
}

func SolvePart2(input string) (string, error) {
	rows := utils.ToStringSlice(input, "\n")
	valves := make([]*valve, len(rows))
	graph := dijkstra.NewGraph()
	at := 0
	lookup := make(map[string]int, len(valves))

	for i := range valves {
		graph.AddVertex(i)
		lookup[rows[i][6:8]] = i
		valves[i] = &valve{}
	}

	for i, r := range rows {
		valveId := r[6:8]
		j := 23
		rateStr := make([]uint8, 0, 3)
		for ; j < len(r); j++ {
			if r[j] == ';' {
				break
			}

			rateStr = append(rateStr, r[j])
		}

		rate, _ := strconv.Atoi(string(rateStr))
		valves[i].index = i
		valves[i].id = valveId
		valves[i].rate = rate

		if valveId == "AA" {
			at = i
		}

		for _, v := range utils.ToStringSlice(r[j+24:], ", ") {
			graph.AddArc(i, lookup[v], 1)
		}
	}

	total := 0
	remaining := make([]int, 0, len(valves)-1)

	for _, v := range valves {
		if v.index != at && v.rate != 0 {
			remaining = append(remaining, v.index)
		}
	}

	queue := []item{
		{
			at:        at,
			minute:    1,
			total:     0,
			rate:      0,
			remaining: remaining,
		},
	}

	var q item

	maxRounds := 27

	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]

		if len(q.remaining) == 0 {
			q.total += q.rate * (31 - q.minute)
			q.minute = maxRounds
		}

		if q.minute == maxRounds {
			total = utils.MaxInt(total, q.total)
			continue
		}

		for i, r := range q.remaining {
			path, _ := graph.Shortest(q.at, r)
			newRemaining := make([]int, len(q.remaining))
			copy(newRemaining, q.remaining)
			newRemaining = append(newRemaining[:i], newRemaining[i+1:]...)
			queue = append(queue, item{
				at:        r,
				minute:    q.minute + utils.MinInt(int(path.Distance)+1, maxRounds-q.minute),
				total:     q.total + (q.rate * utils.MinInt(int(path.Distance)+1, maxRounds-q.minute)),
				rate:      q.rate + valves[r].rate,
				remaining: newRemaining,
			})
		}
	}

	return strconv.Itoa(total), nil
}
