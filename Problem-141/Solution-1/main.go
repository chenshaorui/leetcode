package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		} else {
			fast = fast.Next
		}

		slow = slow.Next
		if slow == fast {
			return true
		}
	}

	return false
}
