package Solution_1

func lengthOfLIS(nums []int) int {
	numTotal := len(nums)

	if len(nums) == 0 {
		return 0
	}

	lengthMinNums := []int{0, nums[0]}
	maxLen := 1

	for i := 1; i < numTotal; i++ {
		num := nums[i]

		if num > lengthMinNums[maxLen] {
			lengthMinNums = append(lengthMinNums, num)
			maxLen++
		} else {
			index := findTheLastNumberLessThanTarget(lengthMinNums, maxLen, num)
			lengthMinNums[index+1] = num
		}
	}

	return maxLen
}

func findTheLastNumberLessThanTarget(nums []int, right int, target int) int {
	left := 1

	for left <= right {
		middle := left + (right-left)>>1

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return right
}
