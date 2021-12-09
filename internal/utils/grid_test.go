package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	x      int
	y      int
	offset int
	length int
}{
	{0, 0, 0, 8},
	{1, 0, 1, 8},
	{0, 1, 8, 8},
	{1, 1, 9, 8},
}

func TestCoordinates(t *testing.T) {
	for _, testCase := range testCases {
		coordinates := fmt.Sprint(testCase.x, "x", testCase.y)

		t.Run(coordinates+"/from", func(t *testing.T) {
			assert.Equal(t, testCase.offset, FromCoordinates(testCase.x, testCase.y, testCase.length))
		})

		t.Run(coordinates+"/to", func(t *testing.T) {
			x, y := ToCoordinates(testCase.offset, testCase.length)

			assert.Equal(t, testCase.x, x)
			assert.Equal(t, testCase.y, y)
		})
	}
}
