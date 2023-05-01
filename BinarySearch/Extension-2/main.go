package Extension_2

func findTheFirstNumberGreaterThanOrEqualToTarget(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1

	for left <= right {
		middle := left + (right-left)>>1

		if nums[middle] >= target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	if left < size {
		return left
	} else {
		return -1
	}
}
