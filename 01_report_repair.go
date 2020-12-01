package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	intLines := make([]int, len(lines))
	for i, val := range lines {
		numVal, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		intLines[i] = numVal
	}

	expectedSum := 2020
	partOne := -1
	partTwo := -1

Loop:
	for i, val := range intLines {
		for j, val2 := range intLines[i:] {
			if partOne == -1 && val + val2 == expectedSum {
				partOne = val * val2
			}

			for _, val3 := range intLines[j:] {
				if partTwo == -1 && val + val2 + val3 == expectedSum {
					partTwo = val * val2 * val3
				}

				if partTwo != -1 && partOne != -1 {
					break Loop
				}
			}
		}
	}

	fmt.Println(partOne)
	fmt.Println(partTwo)
}
