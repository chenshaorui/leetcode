package Solution_1

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)

	citationTotal := len(citations)
	left, right := 0, citationTotal-1

	for left <= right {
		middle := left + (right-left)>>1

		if citationTotal-middle <= citations[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return citationTotal - left
}
