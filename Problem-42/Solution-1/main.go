package Solution_1

func trap(height []int) int {
	n := len(height)

	rainWater := 0

	if n > 0 {
		leftMaxHeights, rightMaxHeights := make([]int, n), make([]int, n)

		leftMaxHeights[0] = height[0]
		for i := 1; i < n; i++ {
			leftMaxHeights[i] = max(leftMaxHeights[i-1], height[i])
		}

		rightMaxHeights[n-1] = height[n-1]
		for i := n - 2; i >= 0; i-- {
			rightMaxHeights[i] = max(rightMaxHeights[i+1], height[i])
		}

		for i := 1; i < n-1; i++ {
			diff := min(leftMaxHeights[i], rightMaxHeights[i]) - height[i]
			if diff > 0 {
				rainWater += diff
			}
		}
	}

	return rainWater
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}
