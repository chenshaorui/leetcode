package Solution_2

func findKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		swim(nums, i)
	}

	size := len(nums)
	for i := k; i < size; i++ {
		if nums[i] > nums[0] {
			nums[0], nums[i] = nums[i], nums[0]
			sink(nums, 0, k)
		}
	}

	return nums[0]
}

func swim(nums []int, i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if nums[i] >= nums[parent] {
			break
		}

		nums[parent], nums[i] = nums[i], nums[parent]
		i = parent
	}
}

func sink(nums []int, i int, size int) {
	for i*2 < size {
		smallest := i

		left := i*2 + 1
		if left < size && nums[left] < nums[smallest] {
			smallest = left
		}

		right := i*2 + 2
		if right < size && nums[right] < nums[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		nums[i], nums[smallest] = nums[smallest], nums[i]
		i = smallest
	}
}
