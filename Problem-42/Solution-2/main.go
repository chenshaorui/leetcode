package Solution_2

func trap(height []int) int {
	n := len(height)

	rainWater := 0

	if n > 0 {
		i, j := 0, n-1
		leftMaxHeight, rightMaxHeight := height[i], height[j]

		for i <= j {
			if leftMaxHeight < rightMaxHeight {
				leftHeight := height[i]

				if leftHeight > leftMaxHeight {
					leftMaxHeight = leftHeight
				} else if leftHeight < leftMaxHeight {
					rainWater += leftMaxHeight - leftHeight
				}
			} else {
				rightHeight := height[j]

				if rightHeight > rightMaxHeight {
					rightMaxHeight = rightHeight
				} else if rightHeight < rightMaxHeight {
					rainWater += rightMaxHeight - rightHeight
				}
			}

			if leftMaxHeight < rightMaxHeight {
				i++
			} else {
				j--
			}
		}
	}

	return rainWater
}
