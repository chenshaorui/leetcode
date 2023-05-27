package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode

	prev1, curr1 := (*ListNode)(nil), l1
	for curr1 != nil {
		next := curr1.Next
		curr1.Next = prev1
		prev1 = curr1
		curr1 = next
	}

	prev2, curr2 := (*ListNode)(nil), l2
	for curr2 != nil {
		next := curr2.Next
		curr2.Next = prev2
		prev2 = curr2
		curr2 = next
	}

	curr1, curr2 = nil, nil
	carry := 0

	for prev1 != nil || prev2 != nil {
		var digit1, digit2 int

		if prev1 == nil {
			digit1 = 0
		} else {
			digit1 = prev1.Val

			next := prev1.Next
			prev1.Next = curr1
			curr1 = prev1
			prev1 = next
		}

		if prev2 == nil {
			digit2 = 0
		} else {
			digit2 = prev2.Val

			next := prev2.Next
			prev2.Next = curr2
			curr2 = prev2
			prev2 = next
		}

		digitSum := digit1 + digit2 + carry
		carry = digitSum / 10

		head = &ListNode{
			Val:  digitSum % 10,
			Next: head,
		}
	}

	if carry == 1 {
		head = &ListNode{
			Val:  1,
			Next: head,
		}
	}

	return head
}
