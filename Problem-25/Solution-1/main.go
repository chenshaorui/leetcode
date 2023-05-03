package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var reversedHead, reversedTail *ListNode

	curr := head
	for curr != nil {
		start := curr

		prev := (*ListNode)(nil)

		i := 0
		for ; i < k && curr != nil; i++ {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		if i != k {
			prev, curr = (*ListNode)(nil), prev

			for curr != nil {
				next := curr.Next
				curr.Next = prev
				prev = curr
				curr = next
			}

			reversedTail.Next = prev

			break
		}

		if reversedTail == nil {
			reversedHead = prev
		} else {
			reversedTail.Next = prev
		}

		reversedTail = start
	}

	if reversedHead == nil {
		return head
	} else {
		return reversedHead
	}
}
