package plusone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	name     string
	input    []int
	expected []int
}

func TestPlusOne(t *testing.T) {
	testCases := []testCase{
		{
			name:     "single digit no carry",
			input:    []int{8},
			expected: []int{9},
		},
		{
			name:     "single carry",
			input:    []int{9},
			expected: []int{1, 0},
		},
		{
			name:     "LC example 1",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "multiple carryover",
			input:    []int{9, 9},
			expected: []int{1, 0, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := plusOne(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
