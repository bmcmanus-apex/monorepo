package kokoeatingbananas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	name           string
	piles          []int
	hours          int
	expectedResult int
}

func TestMinEatingSpeed(t *testing.T) {
	testCases := []testCase{
		{
			name:           "pile count match hours",
			piles:          []int{30, 11, 23, 4, 20},
			hours:          5,
			expectedResult: 30,
		},
		{
			name:           "time to eat",
			piles:          []int{3, 6, 7, 11},
			hours:          8,
			expectedResult: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := minEatingSpeed(tc.piles, tc.hours)

			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
