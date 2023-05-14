package Solution_1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	levelValues := make([][]int, 0)

	if root != nil {
		level := 0
		levelNodes := []*TreeNode{root}

		for len(levelNodes) != 0 {
			levelNodeCount := len(levelNodes)
			values := make([]int, levelNodeCount)
			var nextLevelNodes []*TreeNode

			for i, node := range levelNodes {
				if level%2 == 0 {
					values[i] = node.Val
				} else {
					values[levelNodeCount-1-i] = node.Val
				}

				if node.Left != nil {
					nextLevelNodes = append(nextLevelNodes, node.Left)
				}
				if node.Right != nil {
					nextLevelNodes = append(nextLevelNodes, node.Right)
				}
			}

			level++
			levelNodes = nextLevelNodes
			levelValues = append(levelValues, values)
		}
	}

	return levelValues
}
