package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []testCase{
		{
			name:     "example1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "example2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "example3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualIndices := TwoSum(tc.nums, tc.target)

			assert.EqualValues(t, tc.expected, actualIndices)
		})
	}
}
