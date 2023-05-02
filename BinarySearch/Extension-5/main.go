package Extension_5

func findTheLastNumberLessThanOrEqualToTarget(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1

	for left <= right {
		middle := left + (right-left)>>1

		if nums[middle] <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	if right >= 0 {
		return right
	} else {
		return -1
	}
}
