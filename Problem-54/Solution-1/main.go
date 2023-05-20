package Solution_1

const (
	RightDirection = iota
	DownDirection
	LeftDirection
	UpDirection
)

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return make([]int, 0)
	}

	m := len(matrix[0])

	nums := make([]int, n*m)

	i, j, k := 0, 0, 0
	top, right, bottom, left := 0, m-1, n-1, 0
	direction := RightDirection

	for left <= right && top <= bottom {
		switch direction {
		case RightDirection:
			for ; j < right; j++ {
				nums[k] = matrix[i][j]
				k++
			}

			top++
			direction = DownDirection
		case DownDirection:
			for ; i < bottom; i++ {
				nums[k] = matrix[i][j]
				k++
			}

			right--
			direction = LeftDirection
		case LeftDirection:
			for ; j > left; j-- {
				nums[k] = matrix[i][j]
				k++
			}

			bottom--
			direction = UpDirection
		case UpDirection:
			for ; i > top; i-- {
				nums[k] = matrix[i][j]
				k++
			}

			left++
			direction = RightDirection
		}
	}

	nums[k] = matrix[i][j]

	return nums
}
