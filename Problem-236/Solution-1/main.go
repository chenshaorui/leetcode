package Solution_1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, commonAncestor := search(root, p, q)
	return commonAncestor
}

func search(root, p, q *TreeNode) (int, *TreeNode) {
	count := 0

	if root.Val == p.Val {
		count += 1
	} else if root.Val == q.Val {
		count += 1
	}

	if root.Left != nil {
		leftCount, commonAncestor := search(root.Left, p, q)
		if leftCount == 2 {
			return leftCount, commonAncestor
		}

		count += leftCount
		if count == 2 {
			return count, root
		}
	}

	if root.Right != nil {
		rightCount, commonAncestor := search(root.Right, p, q)
		if rightCount == 2 {
			return rightCount, commonAncestor
		}

		count += rightCount
		if count == 2 {
			return count, root
		}
	}

	return count, nil
}
