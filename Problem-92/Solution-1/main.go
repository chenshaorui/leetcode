package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var leftPrevNode, leftNode *ListNode
	curr := head

	i := 1
	for curr != nil {
		if i == left {
			leftNode = curr
			break
		}

		leftPrevNode = curr
		curr = curr.Next
		i++
	}

	if leftNode == nil {
		return head
	}

	var prev *ListNode
	for curr != nil && i <= right {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next

		i++
	}

	if leftPrevNode == nil {
		head = prev
	} else {
		leftPrevNode.Next = prev
	}
	leftNode.Next = curr

	return head
}
