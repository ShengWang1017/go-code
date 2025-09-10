package array

import (
	"math"
	"sort"
)

func Merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int

	for _, interval := range intervals {
		// boundary condition
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}

		lastResInterval := res[len(res)-1]

		if interval[0] <= lastResInterval[1] {
			lastResInterval[1] = int(math.Max(float64(interval[1]), float64(lastResInterval[1])))
		} else {
			res = append(res, interval)
		}
	}

	return res
}
