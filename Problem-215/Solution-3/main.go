package Solution_3

import "math/rand"

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

func quickSelect(nums []int, left int, right int, k int) int {
	pivot := quickSelectRandomly(nums, left, right)

	if pivot == k {
		return nums[pivot]
	} else if pivot < k {
		return quickSelect(nums, pivot+1, right, k)
	} else {
		return quickSelect(nums, left, pivot-1, k)
	}
}

func quickSelectRandomly(nums []int, left int, right int) int {
	pivot := rand.Intn(right-left+1) + left
	pivotNum := nums[pivot]

	nums[right], nums[pivot] = nums[pivot], nums[right]

	i := left
	for j := left; j < right; j++ {
		if nums[j] >= pivotNum {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]

	return i
}
