package Solution_1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	rootMaxPathSum := 0

	if root != nil {
		rootMaxPathSum = root.Val

		var search func(treeNode *TreeNode) int

		search = func(treeNode *TreeNode) int {
			var leftMaxSum, rightMaxSum int

			if treeNode.Left != nil {
				leftMaxSum = max(search(treeNode.Left), 0)
			}

			if treeNode.Right != nil {
				rightMaxSum = max(search(treeNode.Right), 0)
			}

			treeMaxPathSum := treeNode.Val + leftMaxSum + rightMaxSum

			if treeMaxPathSum > rootMaxPathSum {
				rootMaxPathSum = treeMaxPathSum
			}

			return treeNode.Val + max(leftMaxSum, rightMaxSum)
		}

		search(root)
	}

	return rootMaxPathSum
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
