package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinations(t *testing.T) {
	assert.Equal(t, [][]rune{
		{'a', 'a', 'a'},
		{'a', 'a', 'b'},
		{'a', 'b', 'a'},
		{'a', 'b', 'b'},
		{'b', 'a', 'a'},
		{'b', 'a', 'b'},
		{'b', 'b', 'a'},
		{'b', 'b', 'b'},
	}, Combinations([]rune{'a', 'b'}, 3))
}
