package mergeintervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	name           string
	intervals      [][]int
	expectedResult [][]int
}

func TestMerge(t *testing.T) {
	testCases := []testCase{
		{
			name:           "single interval",
			intervals:      [][]int{{0, 1}},
			expectedResult: [][]int{{0, 1}},
		},
		{
			name:           "single overlap",
			intervals:      [][]int{{0, 2}, {1, 3}},
			expectedResult: [][]int{{0, 3}},
		},
		{
			name:           "second interval contains first",
			intervals:      [][]int{{1, 4}, {0, 4}},
			expectedResult: [][]int{{0, 4}},
		},
		{
			name:           "first interval contains second",
			intervals:      [][]int{{1, 4}, {2, 3}},
			expectedResult: [][]int{{1, 4}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := merge(tc.intervals)

			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
