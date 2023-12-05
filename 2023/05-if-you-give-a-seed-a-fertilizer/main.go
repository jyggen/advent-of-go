package main

import (
	"fmt"
	"math"
	"os"
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

func parsePart1(input string) ([]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int) {
	sections := utils.ToStringSlice(input, "\n\n")
	seeds := utils.ToOptimisticIntSlice(sections[0], false)
	seedToSoil := parseMap(sections[1])
	soilToFertilizer := parseMap(sections[2])
	fertilizerToWater := parseMap(sections[3])
	waterToLight := parseMap(sections[4])
	lightToTemperature := parseMap(sections[5])
	temperatureToHumidity := parseMap(sections[6])
	humidityToLocation := parseMap(sections[7])

	return seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation
}

func parsePart2(input string) ([][2]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int) {
	sections := utils.ToStringSlice(input, "\n\n")
	seeds := parseSeedMap(sections[0])
	seedToSoil := parseMap(sections[1])
	soilToFertilizer := parseMap(sections[2])
	fertilizerToWater := parseMap(sections[3])
	waterToLight := parseMap(sections[4])
	lightToTemperature := parseMap(sections[5])
	temperatureToHumidity := parseMap(sections[6])
	humidityToLocation := parseMap(sections[7])

	return seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation
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
	seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation := parsePart1(input)
	lowest := math.MaxInt

	for _, seed := range seeds {
		soil := resolvePart1(seed, seedToSoil)
		fertilizer := resolvePart1(soil, soilToFertilizer)
		water := resolvePart1(fertilizer, fertilizerToWater)
		light := resolvePart1(water, waterToLight)
		temperature := resolvePart1(light, lightToTemperature)
		humidity := resolvePart1(temperature, temperatureToHumidity)
		location := resolvePart1(humidity, humidityToLocation)
		lowest = min(location, lowest)
	}

	return strconv.Itoa(lowest), nil
}

func SolvePart2(input string) (string, error) {
	seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation := parsePart2(input)

	for i := 0; ; i++ {
		humidity := resolvePart2(i, humidityToLocation)
		temperature := resolvePart2(humidity, temperatureToHumidity)
		light := resolvePart2(temperature, lightToTemperature)
		water := resolvePart2(light, waterToLight)
		fertilizer := resolvePart2(water, fertilizerToWater)
		soil := resolvePart2(fertilizer, soilToFertilizer)
		seed := resolvePart2(soil, seedToSoil)

		for _, s := range seeds {
			if seed >= s[0] && seed <= s[0]+s[1] {
				return strconv.Itoa(i), nil
			}
		}
	}
}
