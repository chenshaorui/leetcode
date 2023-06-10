package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head != nil {
		slow, fast := head, head

		i := 0
		for fast != nil {
			if i > n {
				slow = slow.Next
			}

			fast = fast.Next
			i++
		}

		if i <= n {
			head = slow.Next
		} else {
			slow.Next = slow.Next.Next
		}
	}

	return head
}
