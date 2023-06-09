package Solution_1

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)

	sort.Ints(nums)

	results := make([][]int, 0)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		for left, right := i+1, n-1; left < right; left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}

			for right > left && nums[left]+nums[right] > target {
				right--
			}

			if left == right {
				break
			}

			if nums[left]+nums[right] == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
			}
		}
	}

	return results
}
