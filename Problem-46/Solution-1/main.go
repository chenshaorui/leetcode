package Solution_1

func permute(nums []int) [][]int {
	return search(nums, 0, [][]int{})
}

func search(nums []int, m int, results [][]int) [][]int {
	n := len(nums)

	if m == n {
		result := make([]int, n)
		copy(result, nums)

		results = append(results, result)
	} else {
		for i := m; i < n; i++ {
			nums[m], nums[i] = nums[i], nums[m]
			results = search(nums, m+1, results)
			nums[m], nums[i] = nums[i], nums[m]
		}
	}

	return results
}
