package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/jyggen/advent-of-go/internal/solver"
	"github.com/jyggen/advent-of-go/internal/utils"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

type almanac struct {
	seeds                 [][2]int
	seedToSoil            [][3]int
	soilToFertilizer      [][3]int
	fertilizerToWater     [][3]int
	waterToLight          [][3]int
	lightToTemperature    [][3]int
	temperatureToHumidity [][3]int
	humidityToLocation    [][3]int
}

func parseMap(section string) [][3]int {
	numbers := utils.ToOptimisticIntSlice(section, false)
	numbersMap := make([][3]int, 0, len(numbers)/3)

	for i := 0; i < len(numbers); i += 3 {
		numbersMap = append(numbersMap, [3]int{numbers[i], numbers[i+1], numbers[i+2]})
	}

	return numbersMap
}

func parseSeedMap(section string) [][2]int {
	numbers := utils.ToOptimisticIntSlice(section, false)
	numbersMap := make([][2]int, 0, len(numbers)/2)

	for i := 0; i < len(numbers); i += 2 {
		numbersMap = append(numbersMap, [2]int{numbers[i], numbers[i+1]})
	}

	return numbersMap
}

func parse(sections []string) almanac {
	a := almanac{
		seedToSoil:            parseMap(sections[1]),
		soilToFertilizer:      parseMap(sections[2]),
		fertilizerToWater:     parseMap(sections[3]),
		waterToLight:          parseMap(sections[4]),
		lightToTemperature:    parseMap(sections[5]),
		temperatureToHumidity: parseMap(sections[6]),
		humidityToLocation:    parseMap(sections[7]),
	}

	return a
}

func parsePart1(input string) almanac {
	sections := utils.ToStringSlice(input, "\n\n")
	a := parse(sections)
	seeds := utils.ToOptimisticIntSlice(sections[0], false)
	a.seeds = make([][2]int, 0, len(seeds))

	for _, s := range seeds {
		a.seeds = append(a.seeds, [2]int{s, 0})
	}

	return a
}

func parsePart2(input string) almanac {
	sections := utils.ToStringSlice(input, "\n\n")
	a := parse(sections)
	a.seeds = parseSeedMap(sections[0])

	return a
}

func resolvePart1(source int, sourceDestinationMap [][3]int) int {
	for _, v := range sourceDestinationMap {
		if source >= v[1] && source < v[1]+v[2] {
			return (v[0] + v[2]) - ((v[1] + v[2]) - source)
		}
	}

	return source
}

func resolvePart2(destination int, sourceDestinationMap [][3]int) int {
	for _, v := range sourceDestinationMap {
		if destination >= v[0] && destination < v[0]+v[2] {
			return (v[1] + v[2]) - ((v[0] + v[2]) - destination)
		}
	}

	return destination
}

func SolvePart1(input string) (string, error) {
	a := parsePart1(input)
	lowest := math.MaxInt

	for _, seed := range a.seeds {
		soil := resolvePart1(seed[0], a.seedToSoil)
		fertilizer := resolvePart1(soil, a.soilToFertilizer)
		water := resolvePart1(fertilizer, a.fertilizerToWater)
		light := resolvePart1(water, a.waterToLight)
		temperature := resolvePart1(light, a.lightToTemperature)
		humidity := resolvePart1(temperature, a.temperatureToHumidity)
		location := resolvePart1(humidity, a.humidityToLocation)
		lowest = min(location, lowest)
	}

	return strconv.Itoa(lowest), nil
}

func reverse(i int, a almanac) int {
	humidity := resolvePart2(i, a.humidityToLocation)
	temperature := resolvePart2(humidity, a.temperatureToHumidity)
	light := resolvePart2(temperature, a.lightToTemperature)
	water := resolvePart2(light, a.waterToLight)
	fertilizer := resolvePart2(water, a.fertilizerToWater)
	soil := resolvePart2(fertilizer, a.soilToFertilizer)
	seed := resolvePart2(soil, a.seedToSoil)

	for _, s := range a.seeds {
		if seed >= s[0] && seed <= s[0]+s[1] {
			return i
		}
	}

	return -1
}

func SolvePart2(input string) (string, error) {
	a := parsePart2(input)
	fmt.Println(a)
	minSeed, maxSeed := math.MaxInt, 0

	for _, s := range a.seeds {
		minSeed = min(s[0], minSeed)
		maxSeed = max(s[0]+s[1], maxSeed)
	}

	increase := int(math.Floor(math.Sqrt(float64(maxSeed - minSeed))))

	for i := 0; ; i += increase {
		result := reverse(i, a)

		if result == -1 {
			continue
		}

		result = sort.Search(increase, func(j int) bool {
			return reverse(i-increase+j, a) != -1
		})

		return strconv.Itoa(i - increase + result), nil
	}
}
