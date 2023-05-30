package Solution_1

import "sort"

const mod = 1000000007

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	numTotal := len(nums1)
	sortedNums1 := make([]int, numTotal)

	copy(sortedNums1, nums1)
	sort.Ints(sortedNums1)

	sum, maxDiff := 0, 0
	for i := 0; i < numTotal; i++ {
		diff := abs(nums2[i] - nums1[i])
		sum = (sum + diff) % mod

		middle := search(sortedNums1, numTotal, nums2[i])
		if middle < numTotal {
			maxDiff = max(maxDiff, diff-(sortedNums1[middle]-nums2[i]))
		}
		if middle > 0 {
			maxDiff = max(maxDiff, diff-(nums2[i]-sortedNums1[middle-1]))
		}
	}

	return (sum - maxDiff + mod) % mod
}

func search(nums []int, numTotal int, target int) int {
	left, right := 0, numTotal-1

	for left <= right {
		middle := left + (right-left)>>1

		if nums[middle] <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
