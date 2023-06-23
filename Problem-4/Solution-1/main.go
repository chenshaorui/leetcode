package Solution_1

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	total := m + n

	if total%2 == 0 {
		leftNum, rightNum := findKthNumber(nums1, nums2, m, n, total/2, true)
		return float64(leftNum+rightNum) / 2
	} else {
		middleNum, _ := findKthNumber(nums1, nums2, m, n, total/2+1, false)
		return float64(middleNum)
	}
}

func findKthNumber(nums1 []int, nums2 []int, m int, n int, k int, isEven bool) (int, int) {
	i, j := 0, 0

	for true {
		if k < 2 {
			var left, right int

			if i >= m {
				left = nums2[j]
				j += 1
			} else if j >= n {
				left = nums1[i]
				i += 1
			} else if nums1[i] < nums2[j] {
				left = nums1[i]
				i += 1
			} else {
				left = nums2[j]
				j += 1
			}

			if isEven {
				if i >= m {
					right = nums2[j]
				} else if j >= n {
					right = nums1[i]
				} else {
					right = min(nums1[i], nums2[j])
				}
			}

			return left, right
		}

		if i == m {
			return nums2[j+k-1], nums2[j+k]
		}

		if j == n {
			return nums1[i+k-1], nums1[i+k]
		}

		halfOffset := k/2 - 1

		pivot1 := i + halfOffset
		if pivot1 >= m {
			pivot1 = m - 1
		}

		pivot2 := j + halfOffset
		if pivot2 >= n {
			pivot2 = n - 1
		}

		var offset int

		if nums1[pivot1] <= nums2[pivot2] {
			offset = (pivot1 - i) + 1
			i = pivot1 + 1
		} else {
			offset = (pivot2 - j) + 1
			j = pivot2 + 1
		}

		k -= offset
	}

	return 0, 0
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}
