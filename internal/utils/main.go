package utils

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

var one = big.NewInt(1)

func Combinations(characters []rune, length int) [][]rune {
	result := make([][]rune, 0)

	for _, c := range characters {
		if length > 1 {
			for _, r := range Combinations(characters, length-1) {
				result = append(result, append([]rune{c}, r...))
			}
		} else {
			result = append(result, []rune{c})
		}
	}

	return result
}

// courtesy of https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func Crt(a []*big.Int, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])

	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}

	var x, q, s, z big.Int

	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)

		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d is not a coprime", n1)
		}

		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}

	return x.Mod(&x, p), nil
}

func RotateRelativePoint(x2 int, y2 int, angle int) (int, int) {
	radians := float64(angle) * (math.Pi / 180.0)

	x3 := math.Cos(radians)*float64(x2) - math.Sin(radians)*float64(y2)
	y3 := math.Sin(radians)*float64(x2) + math.Cos(radians)*float64(y2)

	return int(math.Round(x3)), int(math.Round(y3))
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func MaxIntSlice(a []int) int {
	high := 0

	for _, b := range a {
		if b > high {
			high = b
		}
	}

	return high
}

func MinIntSlice(a []int) int {
	low := math.MaxInt64

	for _, b := range a {
		if b < low {
			low = b
		}
	}

	return low
}

func IntRange(min int, max int) []int {
	a := make([]int, max-min+1)

	for i := range a {
		a[i] = min + i
	}

	return a
}

func ManhattanDistance(x int, y int) int {
	return AbsInt(x) + AbsInt(y)
}

func ToIntegerSlice(input string, separator string) ([]int, error) {
	stringSlice := ToStringSlice(input, separator)
	integerSlice := make([]int, len(stringSlice))

	var err error

	for i, val := range stringSlice {
		integerSlice[i], err = strconv.Atoi(val)
		if err != nil {
			return integerSlice, err
		}
	}

	return integerSlice, nil
}

func ToRuneSlice(input string, separator string) [][]rune {
	stringSlice := ToStringSlice(input, separator)
	runeSlice := make([][]rune, len(stringSlice))

	for i, val := range stringSlice {
		runeSlice[i] = []rune(val)
	}

	return runeSlice
}

func ToStringSlice(input string, separator string) []string {
	return strings.Split(strings.TrimSpace(input), separator)
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func ToOptimisticIntSlice(input string, allowNegative bool) []int {
	fields := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsDigit(r) && (allowNegative == false || r != '-')
	})

	integers := make([]int, len(fields))

	for i, s := range fields {
		integers[i], _ = strconv.Atoi(s)
	}

	return integers
}

func In[T comparable](value T, list []T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}
