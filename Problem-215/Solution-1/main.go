package Solution_1

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)

	for i := len(nums) - 1; i > len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeap(nums, 0, heapSize)
	}

	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeap(nums, i, heapSize)
	}
}

func maxHeap(nums []int, i int, heapSize int) {
	largest := i

	left := i*2 + 1
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}

	right := i*2 + 2
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeap(nums, largest, heapSize)
	}
}
