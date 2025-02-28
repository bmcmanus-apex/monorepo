package kokoeatingbananas

import "slices"

func minEatingSpeed(piles []int, h int) int {
	maxServingSize := slices.Max(piles)
	minServingSize := 1
	if len(piles) == h {
		return maxServingSize
	}
	for {
		trialSize := (maxServingSize + minServingSize) / 2
		if canFinish(piles, h, trialSize) {
			maxServingSize = trialSize
		} else {
			minServingSize = trialSize + 1
		}
		if maxServingSize == minServingSize {
			return minServingSize
		}

	}
}

// canFinish is a helper function use to determine whether the piles are all consumable given a serving
// size in the number of hours provided.
func canFinish(piles []int, hrs, serving int) bool {
	var totalHours int
	for _, pile := range piles {
		hoursToConsume := pile / serving
		if pile%serving != 0 {
			hoursToConsume++
		}
		totalHours += hoursToConsume
	}
	return totalHours <= hrs
}
