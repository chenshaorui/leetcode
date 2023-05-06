package Solution_1

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := currentSum

	n := len(nums)
	for left, right := 0, 1; left < n && right < n; right++ {
		if currentSum < 0 {
			currentSum = nums[right]
			left = right
		} else {
			currentSum += nums[right]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
