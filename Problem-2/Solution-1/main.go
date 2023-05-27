package Solution_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode

	curr1, curr2 := l1, l2
	carry := 0

	for curr1 != nil || curr2 != nil {
		var digit1, digit2 int

		if curr1 == nil {
			digit1 = 0
		} else {
			digit1 = curr1.Val
			curr1 = curr1.Next
		}

		if curr2 == nil {
			digit2 = 0
		} else {
			digit2 = curr2.Val
			curr2 = curr2.Next
		}

		digitSum := digit1 + digit2 + carry
		carry = digitSum / 10

		digit := &ListNode{
			Val:  digitSum % 10,
			Next: nil,
		}

		if tail == nil {
			head = digit
		} else {
			tail.Next = digit
		}
		tail = digit
	}

	if carry == 1 {
		digit := &ListNode{
			Val:  1,
			Next: nil,
		}

		if tail == nil {
			head = digit
		} else {
			tail.Next = digit
		}
	}

	return head
}
