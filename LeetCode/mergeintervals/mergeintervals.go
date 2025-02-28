package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	var answer [][]int

	// sort by first interval values to guarantee processing in order
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	currentInterval := 0
	for _, interval := range intervals {
		switch {
		case len(answer) == 0:
			// initial case when the slice is empty
			answer = append(answer, interval)
		case answer[currentInterval][1] >= interval[0] && answer[currentInterval][1] < interval[1]:
			// overlap only
			answer[currentInterval][1] = interval[1]
		case answer[currentInterval][1] >= interval[0] && answer[currentInterval][1] >= interval[1]:
			// interval fully contained by current interval
		default:
			// no overlap found
			answer = append(answer, interval)
			currentInterval++
		}
	}

	return answer
}
