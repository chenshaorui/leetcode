package Solution_1

func twoSum(nums []int, target int) []int {
	numIndexMap := map[int]int{}

	for i, num := range nums {
		if j, ok := numIndexMap[target-num]; ok {
			return []int{j, i}
		}

		numIndexMap[num] = i
	}

	return nil
}
