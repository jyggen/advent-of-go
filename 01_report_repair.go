package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func solve(input []int, expected int, iterations int, attempt []int) (int, error) {
	for i, val := range input {
		newAttempt := append(attempt, val)

		if iterations > 1 {
			returnVal, err := solve(input[i + 1:], expected, iterations - 1, newAttempt)

			if err == nil {
				return returnVal, nil
			}
		} else {
			result := 0

			for _, val := range newAttempt {
				result += val
			}

			if result == expected {
				result = newAttempt[0]

				for _, val := range newAttempt[1:] {
					result = result * val
				}

				return result, nil
			}
		}
	}

	return 0, errors.New("the attempt does not result in the expected number")
}

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

	result, err := solve(intLines, 2020, 2, make([]int, 0))

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	result, err = solve(intLines, 2020, 3, make([]int, 0))

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
