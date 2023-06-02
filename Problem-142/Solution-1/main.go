package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for true {
		if fast == nil {
			return nil
		}

		fast = fast.Next
		if fast == nil {
			return nil
		}

		fast = fast.Next
		if fast == nil {
			return nil
		}

		slow = slow.Next

		if fast == slow {
			break
		}
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
