package Solution_1

func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	charIndexMap := map[byte]int{}

	maxLen := 0
	for left, right := -1, 0; right < strLen; right++ {
		rightChar := s[right]
		if index, ok := charIndexMap[rightChar]; ok {
			left = max(index, left)
		}
		charIndexMap[rightChar] = right
		maxLen = max(right-left, maxLen)
	}

	return maxLen
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}
