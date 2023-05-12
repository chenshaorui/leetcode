package Solution_1

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)>>1

		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			if nums[middle] > nums[right] || target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if nums[middle] < nums[left] || target >= nums[left] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}

	return -1
}
