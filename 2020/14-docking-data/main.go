package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	lines := utils.ToStringSlice(input, "\n")
	memory := make(map[int]int)
	mask := [36]rune{}

	for _, l := range lines {
		if l[:7] == "mask = " {
			for k, v := range l[7:] {
				if v == 'X' {
					v = 0
				}

				mask[k] = v
			}
		} else {
			parts := utils.ToStringSlice(l, " = ")
			address, _ := strconv.Atoi(parts[0][4 : len(parts[0])-1])
			decVal, _ := strconv.Atoi(parts[1])
			binVal := []rune(fmt.Sprintf("%036s", strconv.FormatInt(int64(decVal), 2)))

			for k, v := range mask {
				if v != 0 {
					binVal[k] = v
				}
			}

			decVal64, _ := strconv.ParseInt(string(binVal), 2, 64)
			memory[address] = int(decVal64)
		}
	}

	sum := 0

	for _, v := range memory {
		sum += v
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	lines := utils.ToStringSlice(input, "\n")
	memory := make(map[int]int)
	masks := make([][]rune, 0)

	for _, l := range lines {
		if l[:7] == "mask = " {
			masks = make([][]rune, 0)

			for _, m := range utils.Combinations([]rune{'0', '1'}, strings.Count(l[7:], "X")) {
				mask := make([]rune, 36)
				offset := 0

				for k, v := range l[7:] {
					if v == '0' {
						continue
					}

					if v == 'X' {
						v = m[offset]
						offset++
					}

					mask[k] = v
				}

				masks = append(masks, mask)
			}
		} else {
			parts := utils.ToStringSlice(l, " = ")
			value, _ := strconv.Atoi(parts[1])
			decAddr, _ := strconv.Atoi(parts[0][4 : len(parts[0])-1])
			binAddr := []rune(fmt.Sprintf("%036s", strconv.FormatInt(int64(decAddr), 2)))

			for _, m := range masks {
				maskedBinAddr := make([]rune, len(binAddr))
				copy(maskedBinAddr, binAddr)

				for k, v := range m {
					if v != 0 {
						maskedBinAddr[k] = v
					}
				}

				decAddr64, _ := strconv.ParseInt(string(maskedBinAddr), 2, 64)
				memory[int(decAddr64)] = value
			}
		}
	}

	sum := 0

	for _, v := range memory {
		sum += v
	}

	return strconv.Itoa(sum), nil
}
