package Solution_1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)

	if root != nil {
		if root.Left != nil {
			nums = append(nums, inorderTraversal(root.Left)...)
		}

		nums = append(nums, root.Val)

		if root.Right != nil {
			nums = append(nums, inorderTraversal(root.Right)...)
		}
	}

	return nums
}
