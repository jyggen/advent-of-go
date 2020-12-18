package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/tools/benchmark/parse"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type TestEvent struct {
	Time    time.Time
	Action  string
	Package string
	Test    string
	Elapsed float64
	Output  string
}

type Day struct {
	Pkg   string
	Year  int
	Day   int
	Name  string
	Parts [2]*Part
}

type Part struct {
	Ms float64
}

var pkgReplacer = strings.NewReplacer("github.com/jyggen/advent-of-go/", "")
var nameReplacer = strings.NewReplacer("-", " ")

func main() {
	results := make(map[string]*Day, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		test := &TestEvent{}
		err := json.Unmarshal(scanner.Bytes(), test)

		if err != nil {
			panic(err)
		}

		if test.Action != "output" || !strings.HasPrefix(test.Output, "Benchmark") {
			continue
		}

		benchmark, err := parse.ParseLine(test.Output)

		if err != nil {
			continue
		}

		y, d, n := parsePkg(test.Package)
		p := parseName(benchmark.Name)
		key := fmt.Sprint(y, d)

		if _, ok := results[key]; !ok {
			results[key] = &Day{
				Pkg:  test.Package,
				Year: y,
				Day:  d,
				Name: n,
			}
		}

		durr, err := time.ParseDuration(fmt.Sprintf("%.2f", benchmark.NsPerOp) + "ns")

		if err != nil {
			log.Println(err)
		}

		log.Println(y, d, n, p, durr.String())

		results[key].Parts[p] = &Part{
			Ms: benchmark.NsPerOp / 1000000,
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	output := make([]*Day, len(results))
	i := 0

	for _, p := range results {
		output[i] = p
		i++
	}

	txt, err := json.Marshal(output)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(txt))
}

func parsePkg(pkg string) (int, int, string) {
	pkg = pkgReplacer.Replace(pkg)
	year, err := strconv.Atoi(pkg[0:4])

	if err != nil {
		log.Println(err)
	}

	day, err := strconv.Atoi(pkg[5:7])

	if err != nil {
		log.Println(err)
	}

	return year, day, strings.Title(nameReplacer.Replace(pkg[8:]))
}

func parseName(name string) int {
	parts := strings.Split(name, "/")
	part, err := strconv.Atoi(strings.Split(parts[2], "-")[0])

	if err != nil {
		log.Println(err)
	}

	return part
}
