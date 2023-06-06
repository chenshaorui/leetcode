package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head

	for fast != nil {
		fast = fast.Next

		if fast == nil {
			break
		} else {
			fast = fast.Next
		}

		slow = slow.Next
	}

	prev, curr := (*ListNode)(nil), slow

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	curr = head
	left, right := head.Next, prev

	for left != right {
		curr.Next = right
		curr = right
		right = right.Next

		if right == left {
			break
		}

		curr.Next = left
		curr = left
		left = left.Next
	}
}
