package Solution_1

import "math/rand"

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left int, right int) {
	pivot := partitionRandomly(nums, left, right)
	if left < pivot-1 {
		quickSort(nums, left, pivot-1)
	}
	if right > pivot+1 {
		quickSort(nums, pivot+1, right)
	}
}

func partitionRandomly(nums []int, left int, right int) int {
	pivot := rand.Intn(right-left+1) + left
	pivotNum := nums[pivot]

	nums[right], nums[pivot] = nums[pivot], nums[right]

	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivotNum {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[right], nums[i] = nums[i], nums[right]

	return i
}
