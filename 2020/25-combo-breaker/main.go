package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"strconv"
)

const subjectNumber = 7
const divisionNumber = 20201227

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	publicKeys, _ := utils.ToIntegerSlice(input, "\n")
	cardKey, doorKey := publicKeys[0], publicKeys[1]
	cardLoopSize := 1

	for i := 1; i != cardKey; cardLoopSize++ {
		i = i * subjectNumber % divisionNumber
	}

	encryptionKey := 1

	for i := 1; i < cardLoopSize; i++ {
		encryptionKey = encryptionKey * doorKey % divisionNumber
	}

	return strconv.Itoa(encryptionKey), nil
}

func SolvePart2(input string) (string, error) {
	return strconv.Itoa(0), nil
}
