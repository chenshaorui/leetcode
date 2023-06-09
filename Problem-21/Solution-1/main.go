package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for true {
		if list1 == nil {
			tail.Next = list2
			break
		}

		if list2 == nil {
			tail.Next = list1
			break
		}

		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	return head.Next
}
