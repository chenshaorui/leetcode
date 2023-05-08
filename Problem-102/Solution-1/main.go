package Solution_1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levelValues := make([][]int, 0)

	if root != nil {
		levelNodes := []*TreeNode{root}

		for len(levelNodes) != 0 {
			var values []int
			var nextLevelNodes []*TreeNode

			for _, node := range levelNodes {
				values = append(values, node.Val)

				if node.Left != nil {
					nextLevelNodes = append(nextLevelNodes, node.Left)
				}
				if node.Right != nil {
					nextLevelNodes = append(nextLevelNodes, node.Right)
				}
			}

			levelNodes = nextLevelNodes
			levelValues = append(levelValues, values)
		}
	}

	return levelValues
}
